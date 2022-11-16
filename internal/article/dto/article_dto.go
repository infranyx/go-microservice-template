package article_dto

import (
	"encoding/json"
	"io"
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
