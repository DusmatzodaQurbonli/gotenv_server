package utils

import "Gotenv/internal/app/models"

func GetVarsDecrypted(vars []models.Vars) (resVars []models.Vars, err error) {
	for _, v := range vars {
		title, err := DecryptAES256(v.Title)
		if err != nil {
			return nil, err
		}

		value, err := DecryptAES256(v.Value)
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
