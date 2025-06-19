package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIPProject godoc
// @Summary получить IP проекта
// @Description дает ваш IP
// @Tags auth
// @Accept json
// @Produce json
// @Router /auth/project [get]
func GetIPProject(c *gin.Context) {
	// Получаем IP с помощью c.ClientIP()
	ip := c.ClientIP()

	// Дополнительно получаем заголовки для диагностики
	xForwardedFor := c.GetHeader("X-Forwarded-For")
	xRealIP := c.GetHeader("X-Real-IP")
	remoteAddr := c.Request.RemoteAddr

	// Формируем ответ с подробной информацией
	c.JSON(http.StatusOK, gin.H{
		"client_ip":       ip,
		"x_forwarded_for": xForwardedFor,
		"x_real_ip":       xRealIP,
		"remote_addr":     remoteAddr,
	})
}
