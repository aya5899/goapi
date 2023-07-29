package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/aya5899/goapi/controllers/services"
	"github.com/aya5899/goapi/models"
	"github.com/gorilla/mux"
)

type MyAppController struct {
	service services.MyAppServicer
}

func NewMyAppController(s services.MyAppServicer) *MyAppController {
	return &MyAppController{service: s}
}

// GET /hello
func (c *MyAppController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

// POST /article
func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// json -> Goのデコード処理
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}

	// Go　-> json のエンコード処理
	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "internal exec failure\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

// GET /article/list
func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		// pageに対応する値が複数個ある場合には、最初の値を使用する
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "internal exec failure\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articleList)
}

// GET /article/{id}
func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "internal exec failure\n", http.StatusInternalServerError)
		return
	}
	// Go構造体 -> json
	json.NewEncoder(w).Encode(article)
}

// POST /article/nice
func (c *MyAppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}
	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "internal exec failure\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

// POST /comment
func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
		return
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "internal exec failure\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
