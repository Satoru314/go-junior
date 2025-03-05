package controllers

import (
	"encoding/json"
	"fmt"
	"io"
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
	//
	var reqArticle models.Article
	err := json.NewDecoder(req.Body).Decode(&reqArticle)
	if err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	resArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
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
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	//ページから対応する記事郡をリクエスト
	resArticle, err := c.service.GetArticleListService(page)
	resString := fmt.Sprintf("Article List (page %d)\n", page)
	if err != nil {
		return
	}
	io.WriteString(w, resString)
	json.NewEncoder(w).Encode(resArticle)

}

func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	//求める記事のIDを取得
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Incalid query parameter", http.StatusBadRequest)
		return
	}
	//IDに対応した記事を取得
	resArticle, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "Incalid query parameter", http.StatusBadRequest)
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
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	resArticle, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(resArticle)
}
