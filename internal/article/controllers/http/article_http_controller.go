package articleHttp

import (
	"encoding/json"
	"net/http"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
)

type articleHttpController struct {
	articleUC articleDomain.ArticleUseCase
}

func NewArticleHttpController(uc articleDomain.ArticleUseCase) articleDomain.ArticleHttpController {
	return &articleHttpController{
		articleUC: uc,
	}
}

func (ac articleHttpController) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var aDto articleDto.CreateArticle
	err := decoder.Decode(&aDto)
	if err != nil {
		// respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err = aDto.ValidateCreateArticleDto(); err != nil {
		//
		return
	}

}
