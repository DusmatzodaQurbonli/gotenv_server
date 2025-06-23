package service

import (
	"Gotenv/internal/app/models"
	"Gotenv/internal/app/service/validators"
	"Gotenv/internal/repository"
	"Gotenv/pkg/logger"
	"Gotenv/pkg/utils"
)

func GetProjectsUser(userID uint) (projects []models.Project, err error) {
	projects, err = repository.GetProjectsUser(userID)
	if err != nil {
		return nil, err
	}

	for projectI, project := range projects {
		vars, err := utils.GetVarsDecrypted(project.Vars)
		if err != nil {
			logger.Error.Printf("[service.GetProjectsUser] Error while getting project vars: %v\n", err)

			return nil, err
		}

		project.Vars = vars

		projects[projectI] = project
	}

	return projects, nil
}

func GetProjectByIDAndUserID(userID uint, projectID uint) (project models.Project, err error) {
	project, err = repository.GetProjectByIDAndUserID(userID, projectID)
	if err != nil {
		return models.Project{}, err
	}

	vars, err := utils.GetVarsDecrypted(project.Vars)
	if err != nil {
		logger.Error.Printf("[repository.GetProjectByIDAndUserID] Error while getting project vars: %v\n", err)

		return project, err
	}

	project.Vars = vars

	return project, nil
}

func CreateProject(project *models.Project) (err error) {
	if err = validators.ValidateProject(*project); err != nil {
		return err
	}

	project.Code = utils.GenerateHash(project.Code)
	project.IP = utils.GenerateHash(project.IP)

	err = repository.CreateProject(project)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProject(project models.Project) (err error) {
	project.Code = utils.GenerateHash(project.Code)
	project.IP = utils.GenerateHash(project.IP)

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
