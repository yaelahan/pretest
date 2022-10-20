package custom

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CustomValidator struct {
	db *gorm.DB
}

type CustomValidatorList struct {
	Key      string
	Message  string
	Callback func(level validator.FieldLevel) bool
}

func Register(validate *validator.Validate, trans ut.Translator, db *gorm.DB) {
	customValidator := &CustomValidator{db}

	customValidatorLists := []CustomValidatorList{
		{
			Key:      "uniquedb",
			Message:  "{0} already taken",
			Callback: customValidator.Uniquedb,
		},
	}

	for _, v := range customValidatorLists {
		validate.RegisterValidation(v.Key, v.Callback)
		validate.RegisterTranslation(
			v.Key,
			trans,
			customValidator.registerTranslation(v.Key, v.Message),
			customValidator.translation(v.Key),
		)
	}
}

func (v *CustomValidator) registerTranslation(key string, message string) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		return ut.Add(key, message, true)
	}
}

func (v *CustomValidator) translation(key string) validator.TranslationFunc {
	return func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(key, fe.Field())
		return t
	}
}
