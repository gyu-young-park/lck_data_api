package validator

import (
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func checkIsTeamResultValid(fl validator.FieldLevel) bool {
	field := strings.ToLower(fl.Field().String())
	return field == "w" || field == "l"
}

func checkIsSortOptValid(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	switch field {
	case "asc":
		return true
		// case "desc":
		// 	return true
	}
	return false
}

func checkIsPublishedAtValid(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	_, err := strconv.ParseInt(field, 10, 64)
	if err != nil {
		return false
	}
	return true
}

func checkIsLimitValid(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	limit, err := strconv.ParseInt(field, 10, 64)
	if err != nil {
		return false
	}
	if limit <= 0 {
		return false
	}
	return true
}
