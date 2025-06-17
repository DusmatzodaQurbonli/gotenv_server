package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping godoc
// @Summary Пинг
// @Description Чтобы проверить, работает ли сервис
// @Tags general
// @Accept json
// @Produce json
// @Success 200 {object} models.DefaultResponse "Успешный ответ"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
