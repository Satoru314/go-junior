package main

import (
	"database/sql"
	"fmt"
	"log"
	"myapi/api"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	defer db.Close()
	r := api.NewRouter(db)
	log.Println("server start at pory 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
