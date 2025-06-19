package service

import (
	"Gotenv/internal/app/models"
	"Gotenv/internal/app/service/validators"
	"Gotenv/internal/repository"
)

func GetProjectsUser(userID uint) (projects []models.Project, err error) {
	projects, err = repository.GetProjectsUser(userID)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func GetProjectByIDAndUserID(userID uint, projectID uint) (project models.Project, err error) {
	project, err = repository.GetProjectByIDAndUserID(userID, projectID)
	if err != nil {
		return models.Project{}, err
	}

	return project, nil
}

func CreateProject(project *models.Project) (err error) {
	if err = validators.ValidateProject(*project); err != nil {
		return err
	}

	err = repository.CreateProject(project)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProject(project models.Project) (err error) {
	if err = validators.ValidateProject(project); err != nil {
		return err
	}

	err = repository.UpdateProject(project)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProject(userID, projectID uint) (err error) {
	project, err := GetProjectByIDAndUserID(userID, projectID)
	if err != nil {
		return err
	}

	err = repository.DeleteProject(project)
	if err != nil {
		return err
	}

	return nil
}
