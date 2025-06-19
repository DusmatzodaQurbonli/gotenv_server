package utils

import (
	"Gotenv/pkg/logger"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// EncryptAES256 шифрует входные данные с помощью AES-256 и ключа
func EncryptAES256(plainText string) (string, error) {
	key := os.Getenv("KEY")

	// Проверка длины ключа
	if len(key) != 32 {
		logger.Error.Printf("[utils.EncryptAES256] Error while encrypting data: %s", fmt.Errorf("ключ должен быть длиной 32 символа").Error())

		return "", fmt.Errorf("ключ должен быть длиной 32 символа")
	}

	// Преобразуем ключ в байты
	keyBytes := []byte(key)

	// Создаем блок шифра AES
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		logger.Error.Printf("[utils.EncryptAES256] Error while creating AES cipher: %v", err)

		return "", err
	}

	// Используем режим шифрования GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logger.Error.Printf("[utils.EncryptAES256] Error while creating GCM: %v", err)

		return "", err
	}

	// Генерируем случайный nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		logger.Error.Printf("[utils.EncryptAES256] Error while creating nonce: %v", err)

		return "", err
	}

	// Шифруем данные
	cipherText := aesGCM.Seal(nil, nonce, []byte(plainText), nil)

	// Объединяем nonce и шифротекст
	result := append(nonce, cipherText...)

	// Кодируем в base64 для удобства передачи
	encoded := base64.StdEncoding.EncodeToString(result)

	return encoded, nil
}

// DecryptAES256 расшифровывает данные, зашифрованные функцией encryptAES256
func DecryptAES256(encrypted string) (string, error) {
	key := os.Getenv("KEY")

	// Проверка длины ключа
	if len(key) != 32 {
		logger.Error.Printf("[utils.DecryptAES256] Error while encrypting data: %s", fmt.Errorf("ключ должен быть длиной 32 символа").Error())

		return "", fmt.Errorf("ключ должен быть длиной 32 символа")
	}

	// Преобразуем ключ в байты
	keyBytes := []byte(key)

	// Декодируем из base64
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		logger.Error.Printf("[utils.DecryptAES256] Error while decoding data: %v", err)

		return "", err
	}

	// Создаем блок шифра AES
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		logger.Error.Printf("[utils.DecryptAES256] Error while creating AES cipher: %v", err)

		return "", err
	}

	// Используем режим GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logger.Error.Printf("[utils.DecryptAES256] Error while creating GCM: %v", err)

		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	// Проверяем что длина данных больше чем nonce
	if len(data) < nonceSize {
		logger.Error.Printf("[utils.DecryptAES256] Error while creating nonce: %v", fmt.Errorf("данные слишком короткие"))

		return "", fmt.Errorf("данные слишком короткие")
	}

	// Извлекаем nonce и зашифрованные данные
	nonce, cipherText := data[:nonceSize], data[nonceSize:]

	// Расшифровка
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		logger.Error.Printf("[utils.DecryptAES256] Error while opening data: %v", err)

		return "", err
	}

	return string(plainText), nil
}
