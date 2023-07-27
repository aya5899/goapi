package repositories

import (
	"database/sql"
	"fmt"

	"github.com/aya5899/goapi/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) 
		values (?, ?, ?, 0, now());
		`
	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName
	result, err := db.Exec(sqlStr, newArticle.Title, newArticle.Contents, newArticle.UserName)
	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}
	id, _ := result.LastInsertId()
	newArticle.ID = int(id)
	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset?;
		`
	pageLimit := 5
	rows, err := db.Query(sqlStr, pageLimit, ((page - 1) * pageLimit))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		articleArray = append(articleArray, article)
	}
	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
		`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}
	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println()
		return models.Article{}, err
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	return article, nil
}

func updateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
		`
	const sqlUpdateNice = `
		update articles
		set nice = ?
		where article_id = ?;
		`
	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 指定したIDの記事のいいね数の取得
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}
	// いいね数をnicenumに格納
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}
	// いいね数の更新
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
