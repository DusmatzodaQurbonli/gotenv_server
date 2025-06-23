package controllers

import (
	"Gotenv/internal/app/models"
	"Gotenv/internal/app/service"
	"Gotenv/internal/controllers/middlewares"
	"Gotenv/pkg/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllProjectVars godoc
// @Summary Получить переменные проекта
// @Description Возвращает список переменных, привязанных к проекту
// @Tags ProjectVars
// @Accept json
// @Produce json
// @Param id path int true "ID проекта"
// @Param LoginProject body models.LoginProject true "Данные проекта"
// @Success 200 {array} models.VarsReq
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/vars/val/{id} [post]
// @Security ApiKeyAuth
func GetAllProjectVars(c *gin.Context) {
	var LoginProject models.LoginProject
	if err := c.Bind(&LoginProject); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	LoginProject.ProjectIP = c.ClientIP()

	projectID := c.GetUint(middlewares.ProjectIDCtx)

	vars, err := service.GetAllProjectVars(projectID, LoginProject)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, vars)
}

// CreateProjectVars godoc
// @Summary Создать переменные проекта
// @Description Создаёт новую переменную для проекта
// @Tags ProjectVars
// @Accept json
// @Produce json
// @Param id path int true "ID проекта"
// @Param Vars body models.VarsReq true "Новая переменная"
// @Success 200 {object} models.DefaultResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/vars/{id} [post]
// @Security ApiKeyAuth
func CreateProjectVars(c *gin.Context) {
	projectID := c.GetUint(middlewares.ProjectIDCtx)

	var Vars []models.Vars
	if err := c.Bind(&Vars); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	for VarI, Var := range Vars {
		Var.ProjectID = projectID

		Vars[VarI] = Var
	}

	err := service.CreateProjectVar(Vars)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project variables created",
	})
}

// UpdateProjectVars godoc
// @Summary Обновить переменные проекта
// @Description Обновляет переменные по переданному ID
// @Tags ProjectVars
// @Accept json
// @Produce json
// @Param id path int true "ID переменной"
// @Param Vars body models.VarsReq true "Обновлённая переменная"
// @Success 200 {object} models.DefaultResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/vars/{id} [put]
// @Security ApiKeyAuth
func UpdateProjectVars(c *gin.Context) {
	varsID := c.GetUint(middlewares.VarsIDCtx)
	userID := c.GetUint(middlewares.UserIDCtx)

	var Vars models.Vars
	if err := c.Bind(&Vars); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	Vars.ID = varsID

	err := service.UpdateProjectVar(Vars, userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project variables updated",
	})
}

// DeleteProjectVars godoc
// @Summary Удалить переменную проекта
// @Description Удаляет переменную проекта по ID, переданному через middleware
// @Tags ProjectVars
// @Produce json
// @Param id path int true "ID переменной"
// @Success 200 {object} models.DefaultResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/vars/{id} [delete]
// @Security ApiKeyAuth
func DeleteProjectVars(c *gin.Context) {
	varsID := c.GetUint(middlewares.VarsIDCtx)

	vars, err := service.GetProjectVarByID(varsID)
	if err != nil {
		HandleError(c, err)
		return
	}

	err = service.DeleteProjectVar(vars)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project variables deleted",
	})
}
