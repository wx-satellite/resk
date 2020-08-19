package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	translations "gopkg.in/go-playground/validator.v9/translations/zh"
)


type Student struct {
	Name string `validate:"required"`
	Age int64 	`validate:"gte=0,lte=100"`
}

// 参考文章：https://my.oschina.net/u/4339939/blog/3320338
func main() {
	user := &Student{
		Name: "",
		Age:200,
	}
	validate := validator.New()
	cn := zh.New()
	uni := ut.New(cn,cn)
	tran, found := uni.GetTranslator("zh")
	if found {
		_ = translations.RegisterDefaultTranslations(validate, tran)
	}
	err := validate.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(tran))
		}
	}
}
