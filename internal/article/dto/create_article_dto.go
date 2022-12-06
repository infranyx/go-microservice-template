package articleDto

import (
	validator "github.com/go-ozzo/ozzo-validation"
)

type CreateArticle struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

func (ca *CreateArticle) ValidateCreateArticleDto() error {
	return validator.ValidateStruct(ca,
		validator.Field(&ca.Name, validator.Required, validator.Length(3, 50)),
		validator.Field(&ca.Description, validator.Required, validator.Length(5, 100)),
	)
}
