package testdata

import "github.com/aya5899/goapi/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "mayah",
		NiceNum:  2,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "mayah",
		NiceNum:  4,
	},
}

var CommentTestData = []models.Comment{
	models.Comment{
		ArticleID: 1,
		Message:   "1st comment yeah",
	},
	models.Comment{
		ArticleID: 1,
		Message:   "welcome",
	},
}
