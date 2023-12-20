package validation

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateEqualsIds(paramId int, reqId int32) (bool, error) {
	if int32(paramId) != reqId {
		return false, errors.New("id from request != id from route param")
	}

	return true, nil
}

func ValidateTitle(title string) (bool, error) {
	err := validation.Validate(title, validation.Required)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ValidateContent(content *string) (bool, error) {
	err := validation.Validate(&content, validation.Required)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ValidateCategories(categories []int32) (bool, error) {
	err := validation.Validate(categories, validation.Required)
	if err != nil {
		return false, err
	}

	return true, nil
}
