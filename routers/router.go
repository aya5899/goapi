package routers

import (
	"net/http"

	"github.com/aya5899/goapi/controllers"
	"github.com/gorilla/mux"
)

func NewRouter(aCon *controllers.ArticleController, cCon *controllers.CommentController) *mux.Router {
	r := mux.NewRouter()
	// register Handlers
	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}
