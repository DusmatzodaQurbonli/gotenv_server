package controllers

import (
	"Gotenv/internal/app/models"
	"Gotenv/internal/app/service"
	"Gotenv/internal/controllers/middlewares"
	"Gotenv/internal/repository"
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

	var projectReq models.ProjectReq
	if err := c.ShouldBindJSON(&projectReq); err != nil {
		HandleError(c, err)
		return
	}

	var project models.Project

	project.UserID = userID
	project.Code = projectReq.Code
	project.Title = projectReq.Title
	project.IP = projectReq.IP

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
	var projectReq models.ProjectReq
	if err := c.ShouldBindJSON(&projectReq); err != nil {
		HandleError(c, err)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	projectID := c.GetUint(middlewares.ProjectIDCtx)

	var project models.Project

	project.ID = projectID
	project.UserID = userID
	project.Code = projectReq.Code
	project.Title = projectReq.Title
	project.IP = projectReq.IP

	if err := service.UpdateProject(project); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project updated successfully",
	})
}

// UpdateProjectActive godoc
// @Summary Обновить статус проекта
// @Description Обновляет статус проекта, если он принадлежит пользователю
// @Tags Projects
// @Accept json
// @Produce json
// @Param id path int true "ID проекта"
// @Success 200 {object} models.DefaultResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/{id}/active [patch]
// @Security ApiKeyAuth
func UpdateProjectActive(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	projectID := c.GetUint(middlewares.ProjectIDCtx)

	project, err := service.GetProjectByIDAndUserID(userID, projectID)
	if err != nil {
		HandleError(c, err)
		return
	}

	if project.IsActive {
		project.IsActive = false
	} else {
		project.IsActive = true
	}

	if err = repository.UpdateProject(project); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project active updated successfully",
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
