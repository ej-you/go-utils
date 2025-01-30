package validator

import (
	"strings"

	validatorModule "github.com/go-playground/validator/v10"

	enLocale "github.com/go-playground/locales/en"
	uniTrans "github.com/go-playground/universal-translator"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)


// получение "переводчика" для обработки сообщений ошибок валидации
func GetTranslator() *uniTrans.Translator {
	en := enLocale.New()
	uni := uniTrans.New(en, en)

	trans, _ := uni.GetTranslator("en")
	return &trans
}

// получение валидатора с настроенной обработкой английского языка
func GetValidator(*uniTrans.Translator) *validatorModule.Validate {
	validator := validatorModule.New()
	enTrans.RegisterDefaultTranslations(validator, *trans)

	return validator
}


// получение обработанного словаря ошибок
func GetTranslatedMap(err error, trans *uniTrans.Translator) map[string]string {
	// приводим ошибку к validatorModule.ValidationErrors
	validateErrors, ok := err.(validatorModule.ValidationErrors)
	if !ok {
		return map[string]string{}
	}

	rawTranstaledMap := validateErrors.Translate(*trans)
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
