package middlewares

import (
	"Gotenv/internal/app/service"
	"Gotenv/pkg/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	ProjectIDCtx = "project_id"
	VarsIDCtx    = "vars_id"
)

func CheckUsersProject(c *gin.Context) {
	userID := c.GetUint(UserIDCtx)

	projectIDStr := c.Param("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errs.ErrInvalidID.Error()})
		return
	}

	_, err = service.GetProjectByIDAndUserID(userID, uint(projectID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errs.ErrPermissionDenied.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
		return
	}

	c.Set(ProjectIDCtx, uint(projectID))
	c.Next()
}

func CheckUsersProjectByVarsID(c *gin.Context) {
	userID := c.GetUint(UserIDCtx)

	varsIDStr := c.Param("id")
	varsID, err := strconv.Atoi(varsIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidID.Error()})
		return
	}

	vars, err := service.GetProjectVarByID(uint(varsID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errs.ErrPermissionDenied.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
		return
	}

	_, err = service.GetProjectByIDAndUserID(userID, vars.ProjectID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errs.ErrPermissionDenied.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
		return
	}

	c.Set(VarsIDCtx, uint(varsID))
	c.Set(ProjectIDCtx, vars.ProjectID)

	c.Next()
}
