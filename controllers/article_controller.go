package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"myapi/apperrors"
	"myapi/controllers/services"
	"myapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ArticleControllerの定義
type ArticleController struct {
	service services.ArticleServicer
}

// 外部でのArticleControllerの作成
func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

// 外部からのArticleControllerの利用
func (c *ArticleController) ArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	err := json.NewDecoder(req.Body).Decode(&reqArticle)
	if err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "fail to decode json")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	resArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		err = apperrors.ServiceFuncFailed.Wrap(err, "fail to post article service")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(resArticle)

}

func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	//クエリのマップを取得
	queryMap := req.URL.Query()
	//マップからページのみ取り出し
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}
	//ページから対応する記事郡をリクエスト
	resArticle, err := c.service.GetArticleListService(page)
	if err != nil {
		err = apperrors.BadPathParam.Wrap(err, "fail to get article list")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
	json.NewEncoder(w).Encode(resArticle)

}

func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	//求める記事のIDを取得
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadPathParam.Wrap(err, "pathparam must be number")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	//IDに対応した記事を取得
	resArticle, err := c.service.GetArticleService(articleID)
	if err != nil {
		err = apperrors.ServiceFuncFailed.Wrap(err, "fail to get article detail")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	//記事をエンコード
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
	json.NewEncoder(w).Encode(resArticle)
}

func (c *ArticleController) ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
	var reqArticle models.Article
	err := json.NewDecoder(req.Body).Decode(&reqArticle)
	if err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	resArticle, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		err = apperrors.ServiceFuncFailed.Wrap(err, "fail to post nice service")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(resArticle)
}
