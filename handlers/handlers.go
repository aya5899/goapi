package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/aya5899/goapi/models"
	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// json -> Goのデコード処理
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}

	// Go　-> json のエンコード処理
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
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
	log.Println(page)
	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	log.Println(articleID)
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
		return
	}
	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}
