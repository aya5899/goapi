package repositories_test

import (
	"testing"

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
