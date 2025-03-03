package main

import (
	"database/sql"
	"fmt"
	"log"
	"myapi/controllers"
	"myapi/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}
	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/article", con.ArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.ArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.CommentHandler).Methods(http.MethodPost)
	log.Println("server start at pory 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
