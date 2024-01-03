package translate

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"lancer/conf"
	"lancer/global"
	"strings"
)

//Translation information

func Translate(err error) string {

	var result string

	switch global.ModeData.Language {
	case "zh":

		if err.Error() == "EOF" {

			return "参数不能为空"
		}

		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errors {
				result = err.Translate(conf.Trans)
			}
			return result
		}

		//将 UnmarshalTypeError 转为中文
		if errors, ok := err.(*json.UnmarshalTypeError); ok {
			field := strings.Split(errors.Field, ".")
			fieldName := field[len(field)-1]
			message := fmt.Sprintf("字段 '%s' 的值类型应为 %s", fieldName, errors.Type.String())
			return message
		}
		return err.Error()
	case "en":
		if err.Error() == "EOF" {
			return "the parameter cannot be empty"
		}
		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errors {
				result = err.Translate(conf.Trans)
			}
			return result
		}
		return err.Error()
	default:
		return err.Error()
	}
}
