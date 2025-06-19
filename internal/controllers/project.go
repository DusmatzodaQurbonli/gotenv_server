package controllers

import (
	"Gotenv/internal/app/models"
	"Gotenv/internal/app/service"
	"Gotenv/internal/controllers/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetProjectsUser godoc
// @Summary Получить проекты пользователя
// @Description Возвращает список всех проектов, связанных с пользователем
// @Tags Projects
// @Produce json
// @Success 200 {array} models.ProjectReq
// @Failure 500 {object} models.ErrorResponse
// @Router /projects [get]
// @Security ApiKeyAuth
func GetProjectsUser(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	projects, err := service.GetProjectsUser(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, projects)
}

// GetProjectByIDAndUserID godoc
// @Summary Получить проект по ID
// @Description Возвращает один проект пользователя по ID, если он принадлежит ему
// @Tags Projects
// @Produce json
// @Param id path uint true "ID проекта"
// @Success 200 {object} models.ProjectReq
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/{id} [get]
// @Security ApiKeyAuth
func GetProjectByIDAndUserID(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	projectIDStr := c.Param("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		HandleError(c, err)
		return
	}

	project, err := service.GetProjectByIDAndUserID(userID, uint(projectID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, project)
}

// CreateProject godoc
// @Summary Создать проект
// @Description Создаёт новый проект и привязывает его к текущему пользователю
// @Tags Projects
// @Accept json
// @Produce json
// @Param project body models.ProjectReq true "Данные проекта"
// @Success 200 {object} models.ProjectReq
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects [post]
// @Security ApiKeyAuth
func CreateProject(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		HandleError(c, err)
		return
	}

	project.UserID = userID

	if err := service.CreateProject(&project); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, project)
}

// UpdateProject godoc
// @Summary Обновить проект
// @Description Обновляет данные проекта, если он принадлежит пользователю
// @Tags Projects
// @Accept json
// @Produce json
// @Param id path int true "ID проекта"
// @Param project body models.ProjectReq true "Обновлённые данные проекта"
// @Success 200 {object} models.DefaultResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/{id} [put]
// @Security ApiKeyAuth
func UpdateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		HandleError(c, err)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	projectID := c.GetUint(middlewares.ProjectIDCtx)

	project.ID = projectID
	project.UserID = userID

	if err := service.UpdateProject(project); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project updated successfully",
	})
}

// DeleteProject godoc
// @Summary Удалить проект
// @Description Удаляет проект пользователя по ID
// @Tags Projects
// @Produce json
// @Param id path uint true "ID проекта"
// @Success 200 {object} models.DefaultResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/{id} [delete]
// @Security ApiKeyAuth
func DeleteProject(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	projectID := c.GetUint(middlewares.ProjectIDCtx)

	if err := service.DeleteProject(userID, projectID); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted project",
	})
}
