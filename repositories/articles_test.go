package repositories_test

import (
	"database/sql"
	"fmt"
	"myapi/models"
	"myapi/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleDetail 関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:4000)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	tests := []struct {
		testTitle string         // テストのタイトル
		expected  models.Article // テストで期待する値
	}{
		{
			// 記事 ID1 番のテストデータ
			testTitle: "subtest1",
			expected: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  2,
			},
		}, {
			// 記事 ID2 番のテストデータ
			testTitle: "subtest2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum:  4,
			},
		},
	}
	for _, test := range tests {
		// (test を使ってサブテストを書く)
		got, err := repositories.SelectArticleDetail(db, test.expected.ID)
		if err != nil {
			t.Fatal(err)
		}
		if got.ID != test.expected.ID {
			t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
		}
		if got.Title != test.expected.Title {
			t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
		}
		if got.Contents != test.expected.Contents {
			t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
		}
		if got.UserName != test.expected.UserName {
			t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
		}
		// 記事のいいね数が同じかどうか比較
		if got.NiceNum != test.expected.NiceNum {
			t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
		}
	}
}
