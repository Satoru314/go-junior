package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"myapi/models"
	"myapi/services"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type MyAppController struct {
	// 2. フィールドに MyAppService 構造体を含める
	service *services.MyAppService
}

// コンストラクタの定義
func NewMyAppController(s *services.MyAppService) *MyAppController {
	return &MyAppController{service: s}
}

func (c *MyAppController) ArticleHandler(w http.ResponseWriter, req *http.Request) {
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
func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
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

func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
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
func (c *MyAppController) ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {

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
func (c *MyAppController) CommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
	var reqComment models.Comment
	err := json.NewDecoder(req.Body).Decode(&reqComment)
	if err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	resComment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(resComment)
}
