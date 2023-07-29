package services

import (
	"github.com/aya5899/goapi/models"
	"github.com/aya5899/goapi/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	// databaseへの接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// 記事の取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// コメント一覧の取得
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 取得したコメントを記事に紐付け
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func PostArticleService(article models.Article) (models.Article, error) {
	// databaseへの接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()
	// databaseへの記事の挿入
	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	// databaseへの接続
	db, err := connectDB()
	if err != nil {
		return []models.Article{}, err
	}
	defer db.Close()
	// 指定pageの記事一覧の返却
	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return []models.Article{}, err
	}
	return articleList, nil
}

func PostNiceService(article models.Article) (models.Article, error) {
	// databaseへの接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// 指定した記事のいいね数の更新（+1）
	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}
	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
