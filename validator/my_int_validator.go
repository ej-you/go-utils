package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	validate "github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)


func myIntValidator(fieldInfo reflect.StructField, fieldValue int64, validateTagValues string, errors *validate.Errors) {
	// имя поля для составления ошибки (выбирает значение из тега json; если такого нет - берёт собственно имя поля)
	fieldNameForError, isFound := fieldInfo.Tag.Lookup("json")
	if !isFound {
		fieldNameForError = fieldInfo.Name
	}

	// перебираем значения тега validateTagValues
	for _, tagValue := range strings.Split(validateTagValues, "|") {
		switch {
			// обязательное поле
			case tagValue == "required":
				// валидация средствами библиотеки
				errors.Append(validate.Validate(
					&validators.IntIsPresent{
						Name: "validate_"+fieldNameForError, // название поля
						Field: int(fieldValue), // значение поля
						Message: fmt.Sprintf("%s field must not be blank", fieldNameForError),
					},
				))

			// число (int) больше чем ... (пример, "min:8")
			case strings.HasPrefix(tagValue, "min"):
				// парсинг минимального значения из тега (в int64)
				minInt, err := strconv.ParseInt(strings.TrimPrefix(tagValue, "min:"), 10, 64)
				if err != nil {
					continue
				}
				// проверка значения поля на соответствие минимальному значению
				if fieldValue < minInt {
					errors.Add("validate_"+fieldNameForError, fmt.Sprintf("%s field must be greater than or equal to %d", fieldNameForError, minInt))
				}

			// число (int) меньше чем ... (пример, "max:100")
			case strings.HasPrefix(tagValue, "max"):
				// парсинг максимального значения из тега (в int64)
				maxInt, err := strconv.ParseInt(strings.TrimPrefix(tagValue, "max:"), 10, 64)
				if err != nil {
					continue
				}
				// проверка значения поля на соответствие максимальному значению
				if fieldValue > maxInt {
					errors.Add("validate_"+fieldNameForError, fmt.Sprintf("%s field must be less than or equal to %d", fieldNameForError, maxInt))
				}
		}
	}
}

