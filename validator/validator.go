package validator

import (
	"strings"

	validatorModule "github.com/go-playground/validator/v10"

	enLocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)


type Validator struct {
	ValidatorInstance	*validatorModule.Validate
	Translator			ut.Translator
}


// получение валидатора с настроенной обработкой английского языка
func New() *Validator {
	en := enLocale.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := validatorModule.New(validatorModule.WithRequiredStructEnabled())
	enTrans.RegisterDefaultTranslations(validate, trans)

	return &Validator{validate, trans}
}


// валидация переданной через указатель структуры s, возвращает validatorModule.ValidationErrors
func (this *Validator) Validate(s any) error {
	return this.ValidatorInstance.Struct(s)
}


// получение обработанного словаря ошибок (второй параметр - false для неудачного приведения к validatorModule.ValidationErrors)
func (this Validator) GetMapFromValidationError(err error) (map[string]string, bool) {
	// приводим ошибку к validatorModule.ValidationErrors
	validateErrors, ok := err.(validatorModule.ValidationErrors)
	if !ok {
		return nil, ok
	}

	rawTranstaledMap := validateErrors.Translate(this.Translator)
	// для обработанного словаря
	transtaledMap := make(map[string]string, len(rawTranstaledMap))

	var tempSlice []string
	var key string
	for k, v := range rawTranstaledMap {
		tempSlice = strings.Split(k, ".")
		key = "field" + tempSlice[len(tempSlice) - 1]
		transtaledMap[key] = v
	}
	return transtaledMap, ok
}
