package services_test

import (
	"database/sql"
	"fmt"
	"myapi/services"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var aSer *services.MyAppService

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbHost     = "localhost"
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:4000)/%s?parseTime=true", dbUser,
		dbPassword, dbHost, dbDatabase)
)

func TestMain(m *testing.M) {
	// sql.DB 型を作る
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// sql.DB
	aSer = services.NewMyAppService(db)
	// 個別のベンチマークテストの実行
	m.Run()
}
func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
