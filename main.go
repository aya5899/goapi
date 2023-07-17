package main

import (
	"log"
	"net/http"

	"github.com/aya5899/goapi/handlers"
)

func main() {
	// register Handlers
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	// launch server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
