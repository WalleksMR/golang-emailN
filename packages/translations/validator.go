package translations

import (
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptBrTranslations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var uni *ut.UniversalTranslator

func Register(validate *validator.Validate) ut.Translator {
	ptBR := pt_BR.New()
	uni = ut.New(ptBR, ptBR)
	translate, _ := uni.GetTranslator("pt_BR")
	ptBrTranslations.RegisterDefaultTranslations(validate, translate)
	return translate
}
