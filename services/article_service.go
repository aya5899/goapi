package services

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/aya5899/goapi/apperrors"
	"github.com/aya5899/goapi/models"
	"github.com/aya5899/goapi/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	var amu sync.Mutex
	var cmu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	// 記事の取得
	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		amu.Lock()
		article, articleGetErr = repositories.SelectArticleDetail(db, articleID)
		amu.Unlock()
	}(s.db, articleID)

	// コメント一覧の取得
	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		cmu.Lock()
		commentList, commentGetErr = repositories.SelectCommentList(db, articleID)
		cmu.Unlock()
	}(s.db, articleID)

	wg.Wait()

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "failed to get data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "failed to get data")
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
		err = apperrors.InsertDataFailed.Wrap(err, "failed to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	// 指定pageの記事一覧の返却
	articleList, err := repositories.SelectArticleList(s.db, page)
	// select文クエリの実行に失敗した場合
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to get data")
		return []models.Article{}, err
	}
	// 記事の取得件数が0件だった場合
	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	// 指定した記事のいいね数の更新（+1）
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "failed to update nice count")
		return models.Article{}, err
	}
	// ここよくわからん　return articleではダメ？
	// 多分引数部分のarticleをそのまま返すといいね数が更新されていないので
	// 擬似的にいいね数を＋1
	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
