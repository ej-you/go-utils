package validator

import (
	"strings"

	validatorModule "github.com/go-playground/validator/v10"

	enLocale "github.com/go-playground/locales/en"
	uniTrans "github.com/go-playground/universal-translator"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)


// получение валидатора с настроенной обработкой английского языка
func GetValidator() (*validatorModule.Validate, uniTrans.Translator) {
	validator := validatorModule.New()

	en := enLocale.New()
	uni := uniTrans.New(en, en)

	trans, _ := uni.GetTranslator("en")
	enTrans.RegisterDefaultTranslations(validator, trans)

	return validator, trans
}


// получение обработанного словаря ошибок
func GetTranslatedMap(err error, trans uniTrans.Translator) map[string]string {
	// приводим ошибку к validatorModule.ValidationErrors
	validateErrors, ok := err.(validatorModule.ValidationErrors)
	if !ok {
		return map[string]string{}
	}

	rawTranstaledMap := validateErrors.Translate(trans)
	// для обработанного словаря
	transtaledMap := make(map[string]string, len(rawTranstaledMap))

	var tempSlice []string
	var key string
	for k, v := range rawTranstaledMap {
		tempSlice = strings.Split(k, ".")
		key = "validate" + tempSlice[len(tempSlice) - 1]
		transtaledMap[key] = v
	}
	return transtaledMap
}
