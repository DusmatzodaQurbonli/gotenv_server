package validators

import (
	"Gotenv/internal/app/models"
	"Gotenv/pkg/errs"
)

func ValidateProject(project models.Project) (err error) {
	if project.UserID == emptyInt {
		return errs.ErrUserIDIsEmpty
	}

	if project.Code == emptyString {
		return errs.ErrCodeIsEmpty
	}

	if project.IP == emptyString {
		return errs.ErrIPIsEmpty
	}

	return nil
}
