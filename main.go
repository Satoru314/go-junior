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
	dbHost     = "localhost"
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:4000)/%s?parseTime=true", dbUser,
		dbPassword, dbHost, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}
	defer db.Close()
	r := api.NewRouter(db)
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
