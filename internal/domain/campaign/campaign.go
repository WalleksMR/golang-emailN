package campaign

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/walleksmr/golang-emailn/packages/translations"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}
	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}

	validate := validator.New()

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	TranslateValidator := translations.Register(validate)
	err := validate.Struct(campaign)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return nil, errors.New(validationErrors[0].Translate(TranslateValidator))
	}

	return campaign, nil
}
