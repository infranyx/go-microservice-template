package article_dto

import (
	"encoding/json"
	"io"

	validator "github.com/go-ozzo/ozzo-validation"
)

// CreateArticleRequest is an representation request body to create a new Article
type CreateArticle struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

// FromJSONCreateArticleRequest converts json body request to a CreateArticleRequest struct
func FromJSONCreateArticleRequest(body io.Reader) (*CreateArticle, error) {
	createArticleRequest := CreateArticle{}
	if err := json.NewDecoder(body).Decode(&createArticleRequest); err != nil {
		return nil, err
	}

	return &createArticleRequest, nil
}

func (ca *CreateArticle) ValidateCreateArticleDto() error {
	return validator.ValidateStruct(ca,
		validator.Field(&ca.Name, validator.Required, validator.Length(3, 50)),
		validator.Field(&ca.Description, validator.Required, validator.Length(5, 100)),
	)
}
