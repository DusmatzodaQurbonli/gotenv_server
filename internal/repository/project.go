package repository

import (
	"Gotenv/internal/app/models"
	"Gotenv/pkg/db"
	"Gotenv/pkg/logger"
)

func GetProjectByID(projectID uint) (project models.Project, err error) {
	if err = db.GetDBConn().Model(&models.Project{}).Where("id = ?", projectID).Find(&project).Error; err != nil {
		logger.Error.Printf("[repository.GetProjectByID] Error while getting project by id: %v\n", err)

		return project, TranslateGormError(err)
	}

	return project, nil
}

func GetProjectsUser(userID uint) (projects []models.Project, err error) {
	if err = db.GetDBConn().Model(&models.Project{}).Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		logger.Error.Printf("[repository.GetProjectsUser] Error while getting projects: %v\n", err)

		return nil, TranslateGormError(err)
	}

	return projects, nil
}

func GetProjectByIDAndUserID(userID uint, projectID uint) (project models.Project, err error) {
	if err = db.GetDBConn().Model(&models.Project{}).Where("user_id = ? AND id = ?", userID, projectID).First(&project).Error; err != nil {
		logger.Error.Printf("[repository.GetProjectByIDAndUserID] Error while getting project: %v\n", err)

		return project, TranslateGormError(err)
	}

	return project, nil
}

func CreateProject(project *models.Project) (err error) {
	if err = db.GetDBConn().Model(&models.Project{}).Create(project).Error; err != nil {
		logger.Error.Printf("[repository.CreateProject] Error while creating project: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateProject(project models.Project) (err error) {
	if err = db.GetDBConn().Model(&models.Project{}).Where("id = ?", project.ID).Updates(&project).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProject] Error while updating project: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteProject(project models.Project) (err error) {
	if err = db.GetDBConn().Model(&models.Project{}).Delete(&project).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProject] Error while deleting project: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}
