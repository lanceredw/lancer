package conf

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"lancer/global"
	"lancer/translations"
	"reflect"
	"strings"
)

var Trans ut.Translator

// register validation

func InitValidation(local string) (err error) {

	switch local {
	case "en":
		global.TranslateMessage = translations.MessagesEn
	case "zh":
		global.TranslateMessage = translations.MessagesZh
	default:
		global.TranslateMessage = translations.MessagesEn

	}

	err = registerValidation(local)
	if err != nil {
		return err
	}

	registerCustomValidator()
	return nil

}

func registerValidation(locale string) (err error) {

	//modify gin framework validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//register json tag method
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			label := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return label + " " + name
		})
		zhT := zh.New() //Chinese translator
		enT := en.New() //English translator
		//The first parameter is the backup language environment, and the following parameters are the language environments that should be supported
		uni := ut.New(enT, zhT, enT)

		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			err = entranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhtranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = entranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}

// customer validator
func registerCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && 0 == strings.Index(value, "a") {
					return true
				}
			}

			return false
		})
	}

	//注册自定义验证器
	//if v, ok = binding.Validator.Engine().(*validator.Validate); ok {

	//手机验证
	//v.RegisterValidation("MobileValida", request.MobileValida)
	//v.RegisterTranslation("MobileValida", Trans, func(ut ut.Translator) error {
	//	return ut.Add("MobileValida", "{0}手机号有误", true)
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("MobileValida", fe.Field())
	//	return t
	//})
}
