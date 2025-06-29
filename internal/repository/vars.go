package repository

import (
	"Gotenv/internal/app/models"
	"Gotenv/pkg/db"
	"Gotenv/pkg/errs"
	"Gotenv/pkg/logger"
)

func GetAllProjectVars(project models.Project, projectID uint, projectLogin models.LoginProject) (vars []models.Vars, err error) {
	if !(project.Code == projectLogin.Code && project.IP == projectLogin.ProjectIP) {
		return nil, errs.ErrPermissionDenied
	}

	if err = db.GetDBConn().Model(&models.Vars{}).Where("project_id = ?", projectID).Find(&vars).Error; err != nil {
		logger.Error.Printf("[repository.GetAllProjectVars] Error while fetching project vars: %v\n", err)

		return nil, TranslateGormError(err)
	}

	return vars, nil
}

func GetProjectVarByID(varID uint) (vars models.Vars, err error) {
	if err = db.GetDBConn().Model(&models.Vars{}).Where("id = ?", varID).First(&vars).Error; err != nil {
		logger.Error.Printf("[repository.GetProjectVarByID] Error while fetching project vars by id: %v\n", err)

		return vars, TranslateGormError(err)
	}

	return vars, nil
}

func GetProjectVarByTitle(projectID uint, title string) (vars models.Vars, err error) {
	if err = db.GetDBConn().Model(&models.Vars{}).Where("title = ? AND project_id = ?", title, projectID).First(&vars).Error; err != nil {
		logger.Error.Printf("[repository.GetProjectVarByTitle] Error while fetching project vars by title: %v\n", err)

		return vars, TranslateGormError(err)
	}

	return vars, nil
}

func CreateProjectVar(vars []models.Vars) (err error) {
	if err = db.GetDBConn().Model(&models.Vars{}).Create(&vars).Error; err != nil {
		logger.Error.Printf("[repository.CreateProjectVar] Error while creating project vars: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateProjectVar(vars []models.Vars) (err error) {
	tx := db.GetDBConn().Begin()

	for _, variable := range vars {
		if variable.ID != 0 {
			if err = tx.Model(&models.Vars{}).Where("id = ?", variable.ID).Updates(variable).Error; err != nil {
				logger.Error.Printf("[repository.UpdateProjectVar] Error while updating project vars: %v\n", err)

				return TranslateGormError(err)
			}
		} else {
			if err = tx.Model(&models.Vars{}).Create(&variable).Error; err != nil {
				logger.Error.Printf("[repository.UpdateProjectVar] Error while creating project vars: %v\n", err)

				return TranslateGormError(err)
			}
		}
	}

	if err = tx.Commit().Error; err != nil {
		logger.Error.Printf("[repository.UpdateProjectVar] Error while commiting transaction: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteProjectVar(vars models.Vars) (err error) {
	if err = db.GetDBConn().Model(&models.Vars{}).Delete(&vars).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProjectVar] Error while deleting project vars: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}
