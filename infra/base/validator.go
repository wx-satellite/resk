package base

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"resk/infra"
)

var validate *validator.Validate
var trans *ut.Translator

func Validate() *validator.Validate {
	return validate
}

func Trans() *ut.Translator {
	return trans
}

type ValidatorStarter struct {
	infra.BaseStarter
}

func (s *ValidatorStarter) Init(ctx infra.StarterContext) {
	validate := validator.New()
	cn := zh.New()
	uni := ut.New(cn, cn)
	tran, found := uni.GetTranslator("zh")
	if found {
		_ = translations.RegisterDefaultTranslations(validate, tran)
	} else {
		logrus.Error("Translator not found zh")
	}
}
