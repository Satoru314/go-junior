package api

import (
	"database/sql"
	"myapi/api/middlewares"
	"myapi/controllers"
	"myapi/services"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()
	r.HandleFunc("/article", aCon.ArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.ArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.CommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.AuthMiddleware)
	return r
}
