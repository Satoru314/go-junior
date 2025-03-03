package repositories

import (
	"database/sql"
	"fmt"
	"myapi/models"
	"time"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, ?);
	`
	article.CreatedAt = time.Now()
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName, article.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return article, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return article, err

	}
	article.ID = int(id)

	newArticle := article
	// (問 1) 構造体 `models.Article`を受け取って、それをデータベースに挿入する処理
	return newArticle, nil
}

// 変数 page で指定されたページに表示する投稿一覧をデータベースから取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
	select article_id, title, contents, username, nice from articles
limit ? offset ?;
`
	articleLimit := 6
	rows, err := db.Query(sqlStr, articleLimit, articleLimit*(page-1))
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	defer rows.Close()
	var article models.Article
	var articleArray []models.Article
	for rows.Next() {
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		articleArray = append(articleArray, article)
	}

	// (問 2) 指定された記事データをデータベースから取得して、
	// それを models.Article 構造体のスライス []models.Article に詰めて返す処理
	return articleArray, nil
}

// 投稿 ID を指定して、記事データを取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
select *
from articles
where article_id = ?;
`
	var article models.Article
	rows := db.QueryRow(sqlStr, articleID)
	rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreatedAt)
	// (問 3) 指定 ID の記事データをデータベースから取得して、それを models.Article 構造体の形で返す処理
	return article, nil
}

// いいねの数を update する関数
// -> 発生したエラーを返り値にする
func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
select nice
from articles
where article_id = ?;
`
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	// (問 4) 指定された ID の記事のいいね数を+1 するようにデータベースの中身を更新する処理
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer tx.Rollback()
	// 現在のいいね数を取得するクエリを実行する

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return err
	}
	// 変数 nicenum に現在のいいね数を読み込む
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// いいね数を+1 する更新処理を行う

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// コミットして処理内容を確定させる
	tx.Commit()
	return nil
}
