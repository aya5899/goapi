package repositories_test

import (
	"testing"

	"github.com/aya5899/goapi/models"
	"github.com/aya5899/goapi/repositories"
	"github.com/aya5899/goapi/repositories/testdata"
)

func TestSelectCommentList(t *testing.T) {
	expectedCommentNum := len(testdata.CommentTestData)
	got, err := repositories.SelectCommentList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}
	if commentNum := len(got); commentNum != expectedCommentNum {
		t.Errorf("want %d but got %d articles \n", expectedCommentNum, commentNum)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "comment insertion test",
	}
	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)

	}
	// clean up
	t.Cleanup(func() {
		const sqlStr = `
		delete from comments
		where message = ?;
		`
		testDB.Exec(sqlStr, comment.Message)
	})
}
