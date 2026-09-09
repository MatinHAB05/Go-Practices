package validation

import (
	utils "movieapp/util"

	"github.com/go-playground/validator/v10"
)

func MovieTitleValid(fl validator.FieldLevel) bool {
	return utils.IsValidMovieTitle(fl.Field().String())
}

func ComparatorOperatorValid(fl validator.FieldLevel) bool {
	return utils.IsValudComparatorOperator(fl.Field().String())

}

func MovieQualityValid(fl validator.FieldLevel) bool {
	return utils.IsValidVideoQuality(fl.Field().String())

}

func ActorNameValid(fl validator.FieldLevel) bool {
	return utils.IsValidActorName(fl.Field().String())

}

func ReleaseYearValid(fl validator.FieldLevel) bool {
	return utils.IsValidReleaseYear(int(fl.Field().Int()))

}
