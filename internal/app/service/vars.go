package service

import (
	"Gotenv/internal/app/models"
	"Gotenv/internal/app/service/validators"
	"Gotenv/internal/repository"
	"Gotenv/pkg/errs"
	"Gotenv/pkg/utils"
	"errors"
)

func GetAllProjectVars(projectID uint, projectLogin models.LoginProject) (vars []models.Vars, err error) {
	project, err := repository.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	if !project.IsActive {
		return nil, errs.ErrProjectIsUnactive
	}

	projectLogin.Code = utils.GenerateHash(projectLogin.Code)
	projectLogin.ProjectIP = utils.GenerateHash(projectLogin.ProjectIP)

	vars, err = repository.GetAllProjectVars(project, projectID, projectLogin)
	if err != nil {
		return nil, err
	}

	var resVars []models.Vars
	for _, v := range vars {
		title, err := utils.DecryptAES256(v.Title)
		if err != nil {
			return nil, err
		}

		value, err := utils.DecryptAES256(v.Value)
		if err != nil {
			return nil, err
		}

		decryptedVars := models.Vars{
			Model:     v.Model,
			Title:     title,
			Value:     value,
			ProjectID: v.ProjectID,
		}

		resVars = append(resVars, decryptedVars)
	}

	return resVars, nil
}

func GetProjectVarByID(varID uint) (vars models.Vars, err error) {
	vars, err = repository.GetProjectVarByID(varID)
	if err != nil {
		return vars, err
	}

	return vars, nil
}

func GetProjectVarByTitle(projectID uint, title string) (vars models.Vars, err error) {
	vars, err = repository.GetProjectVarByTitle(projectID, title)
	if err != nil {
		return vars, err
	}

	return vars, nil
}

func CreateProjectVar(vars []models.Vars) (err error) {
	for variableI, variable := range vars {
		if err = validators.ValidateVars(&variable); err != nil {
			return err
		}

		vars[variableI] = variable
	}

	err = repository.CreateProjectVar(vars)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProjectVar(vars models.Vars, userID uint) (err error) {
	if vars.ProjectID != 0 {
		_, err = GetProjectByIDAndUserID(userID, vars.ProjectID)
		if err != nil {
			if errors.Is(err, errs.ErrRecordNotFound) {
				return errs.ErrPermissionDenied
			}

			return errs.ErrSomethingWentWrong
		}
	}

	if err = validators.ValidateVars(&vars); err != nil {
		return err
	}

	err = repository.UpdateProjectVar(vars)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProjectVar(vars models.Vars) (err error) {
	err = repository.DeleteProjectVar(vars)
	if err != nil {
		return err
	}

	return nil
}
