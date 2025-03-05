package routers

import (
	"myapi/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(aCon *controllers.ArticleController, cCon *controllers.CommentController) *mux.Router {
	r := mux.NewRouter()
	// 1. ArticleHandler を登録
	// 3. ArticleListHandler を登録
	// 4. NiceHandler を登録
	// 5. CommentHandler を登録
	// 6. ArticleDetailHandler を登録
	// 7. ルーティングを返却
	r.HandleFunc("/article", aCon.ArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.ArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.CommentHandler).Methods(http.MethodPost)

	return r
}
