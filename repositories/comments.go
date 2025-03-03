package repositories

import (
	"database/sql"
	"fmt"
	"myapi/models"
	"time"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
	insert into comments (article_id, message, created_at) values
	(?, ?, ?);
	`
	// (問 5) 構造体 models.Comment を受け取って、それをデータベースに挿入する処理
	comment.CreatedAt = time.Now()
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message, comment.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return comment, err
	}
	comment.CommentID = int(id)
	newComment := comment
	return newComment, nil
}

// 指定 ID の記事についたコメント一覧を取得する関数
// -> 取得したコメントデータと、発生したエラーを返り値にする
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
	select *
	from comments
	where article_id = ?;
	`
	// (問 6) 指定 ID の記事についたコメント一覧をデータベースから取得して、
	// それを `models.Comment`構造体のスライス `[]models.Comment`に詰めて返す処理
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var comment models.Comment
	var commentArray []models.Comment
	for rows.Next() {
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &comment.CreatedAt)
		commentArray = append(commentArray, comment)
	}
	return commentArray, nil
}
