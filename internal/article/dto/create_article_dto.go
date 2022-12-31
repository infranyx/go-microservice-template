package articleDto

import (
	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type CreateArticleRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

func (caDto *CreateArticleRequestDto) ValidateCreateArticleDto() error {
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

type CreateArticleResponseDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
}
