package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"pretest-indihomesmart/internal/validator/custom"
	"reflect"
	"strings"
)

type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func New(validate *validator.Validate, db *gorm.DB) *Validator {
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

	// register custom validator
	custom.Register(validate, trans, db)

	return &Validator{
		validate: validate,
		trans:    trans,
	}
}

func (v *Validator) Validate(s interface{}) interface{} {
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

	return nil
}
