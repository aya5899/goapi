package services

import (
	"github.com/aya5899/goapi/models"
	"github.com/aya5899/goapi/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// 記事の取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// コメント一覧の取得
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 取得したコメントを記事に紐付け
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	// databaseへの記事の挿入
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	// 指定pageの記事一覧の返却
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return []models.Article{}, err
	}
	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	// 指定した記事のいいね数の更新（+1）
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		return models.Article{}, err
	}
	// ここよくわからん　return articleではダメ？
	// 多分引数部分のarticleをそのまま返すといいね数が更新されていないので
	// 擬似的にいいね数を＋1したarticleを定義して返している？
	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
