package question

import (
	"github.com/russross/meddler"

	"github.com/hachibeeDI/tiny-akasha/model/entity"
)

type Question struct {
	ID       int    `meddler:"id,pk" json:"id"`
	Title    string `meddler:"title" json:"title"`
	Username string `meddler:"username" json:"username"`
	Content  string `meddler:"content" json:"content"`
}

func Init(title, username, content string) *Question {
	return &Question{Title: title, Username: username, Content: content}
}

func (q *Question) Insert(db entity.DB) error {
	return meddler.Insert(db, "question", &q)
}

func SelectAll(db entity.DB) []*Question {
	var ques []*Question
	meddler.QueryAll(db, &ques, "select * from question")
	return ques
}
