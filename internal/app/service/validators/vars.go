package validators

import (
	"Gotenv/internal/app/models"
	"Gotenv/pkg/errs"
	"Gotenv/pkg/utils"
)

func ValidateVars(vars *models.Vars) (err error) {
	if vars.Title == emptyString {
		return errs.ErrInvalidTitle
	}

	if vars.ProjectID == emptyInt {
		return errs.ErrProjectIDIsEmpty
	}

	if vars.Value == emptyString {
		return errs.ErrValueIsEmpty
	}

	vars.Title, err = utils.EncryptAES256(vars.Title)
	if err != nil {
		return err
	}

	vars.Value, err = utils.EncryptAES256(vars.Value)
	if err != nil {
		return err
	}

	return nil
}
