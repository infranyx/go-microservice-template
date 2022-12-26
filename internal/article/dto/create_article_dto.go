package articleDto

import (
	validator "github.com/go-ozzo/ozzo-validation"
)

type CreateArticleDto struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

func (caDto *CreateArticleDto) ValidateCreateArticleDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.Name,
			validator.Required,
			validator.Length(3, 50),
		),
		validator.Field(
			&caDto.Description,
			validator.Required,
			validator.Length(5, 100),
		),
	)
}
