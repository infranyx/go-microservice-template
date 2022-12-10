package articleHttp

import (
	"github.com/gorilla/mux"
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
)

// Service service
type Router struct {
	articleCtrl articleDomain.ArticleUseCase
}

// NewAuthAPI NewAuthAPI
func NewAuthAPI(ahc articleDomain.ArticleUseCase) *Router {
	return &Router{
		articleCtrl: ahc,
	}
}

// Register Register
func (r *Router) Register(router *mux.Router) {
	// router.Use(middleware.Logger)

	router.HandleFunc("/admin/signin", r.articleCtrl.Create).Methods("POST")
}
