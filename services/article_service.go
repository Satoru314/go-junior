package services

import (
	"database/sql"
	"errors"
	"myapi/apperrors"
	"myapi/models"
	"myapi/repositories"
)

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	// TODO : 実装
	resArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to insert article")
		return models.Article{}, err
	}
	return resArticle, nil
}

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// 記事IDを送ると記事にコメントを紐づけた後、Article型で
	// TODO : sql.DB 型を手に入れて、変数 db に代入する
	// 1. repositories 層の関数 SelectArticleDetail で記事の詳細を取得
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	type articleResult struct {
		article models.Article
		err     error
	}
	articleChan := make(chan articleResult)
	defer close(articleChan)
	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, articleGetErr = repositories.SelectArticleDetail(db, articleID)
		ch <- articleResult{article: article, err: articleGetErr}
	}(articleChan, s.db, articleID)

	type commentResult struct {
		commentList []models.Comment
		err         error
	}
	commentChan := make(chan commentResult)
	defer close(commentChan)
	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		commentList, commentGetErr = repositories.SelectCommentList(db, articleID)
		ch <- commentResult{commentList: commentList, err: commentGetErr}
	}(commentChan, s.db, articleID)

	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = cr.commentList, cr.err
		}
	}

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			articleGetErr = apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, articleGetErr
		}
		articleGetErr = apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, articleGetErr
	}
	// 2. repositories 層の関数 SelectCommentList でコメント一覧を取得
	if commentGetErr != nil {
		commentGetErr = apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, commentGetErr
	}
	// 3. 2 で得たコメント一覧を、1 で得た Article 構造体に紐付ける
	article.CommentList = commentList
	return article, nil
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	// TODO : 実装
	resArticleArray, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(resArticleArray) == 0 {
		err = apperrors.NAData.Wrap(ErrNoData, "no data")
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
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice num")
		return models.Article{}, err
	}
	resArticle, err := repositories.SelectArticleDetail(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}
	return resArticle, nil
}
