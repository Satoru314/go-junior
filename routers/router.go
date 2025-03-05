package routers

import (
	"myapi/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(con *controllers.MyAppController) *mux.Router {
	r := mux.NewRouter()
	// 1. ArticleHandler を登録
	// 3. ArticleListHandler を登録
	// 4. NiceHandler を登録
	// 5. CommentHandler を登録
	// 6. ArticleDetailHandler を登録
	// 7. ルーティングを返却
	r.HandleFunc("/article", con.ArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.ArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.CommentHandler).Methods(http.MethodPost)

	return r
}
