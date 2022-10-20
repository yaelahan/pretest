package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strings"
)

type CustomValidator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func NewCustomValidator(validate *validator.Validate) *CustomValidator {
	// override Field() to get json tag
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	// setup default translation
	english := en.New()
	uni := ut.New(english, english)

	trans, _ := uni.GetTranslator("en")
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		panic(err.Error())
	}

	return &CustomValidator{
		validate: validate,
		trans:    trans,
	}
}

func (v *CustomValidator) Validate(s interface{}) interface{} {
	var errorValidations []interface{}

	err := v.validate.Struct(s)
	if err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			errorValidations = append(errorValidations, fiber.Map{
				fieldError.Field(): fieldError.Translate(v.trans),
			})
		}

		return errorValidations
	}

	//if len(errorValidations) > 0 {
	//	panic(exceptions.ValidationError{
	//		Message: "Validation has error.",
	//		Errors:  errorValidations,
	//	})
	//}

	return nil
}
