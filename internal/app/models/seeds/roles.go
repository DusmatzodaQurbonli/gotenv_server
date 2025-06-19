package seeds

import (
	"Gotenv/internal/app/models"
	"Gotenv/pkg/logger"
	"errors"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	// Определяем стандартные ралли
	roles := []models.Role{
		{ID: 1, Name: "Admin"},
		{ID: 2, Name: "User"},
	}

	for _, role := range roles {
		// Проверяем, существует ли роль
		var existingRole models.Role
		if err := db.First(&existingRole, "name = ?", role.Name).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Если роль не найдена, создаем её
				db.Create(&role)
			} else {
				// Обработка других ошибок
				logger.Error.Printf("[seeds.SeedRoles] Error seeding roles: %v", err)

				return err
			}
		}
	}

	return nil
}
