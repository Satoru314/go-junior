package services

import (
	"myapi/models"
	"myapi/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// 記事IDを送ると記事にコメントを紐づけた後、Article型で
	// TODO : sql.DB 型を手に入れて、変数 db に代入する
	// 1. repositories 層の関数 SelectArticleDetail で記事の詳細を取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 2. repositories 層の関数 SelectCommentList でコメント一覧を取得
	article.CommentList, err = repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 3. 2 で得たコメント一覧を、1 で得た Article 構造体に紐付ける
	return article, nil
}
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	// TODO : 実装
	resArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}
	return resArticle, nil
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	// TODO : 実装
	resArticleArray, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}
	return resArticleArray, nil
}

// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	// TODO : 実装
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		return models.Article{}, err
	}
	resArticle, err := repositories.SelectArticleDetail(s.db, article.ID)
	if err != nil {
		return models.Article{}, err
	}
	return resArticle, nil
}
