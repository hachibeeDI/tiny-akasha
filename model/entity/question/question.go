package question

import (
	"errors"

	"github.com/russross/meddler"

	"github.com/hachibeeDI/tiny-akasha/model/entity"
)

type Question struct {
	ID       int    `meddler:"id,pk" json:"id"`
	Title    string `meddler:"title" json:"title"`
	Username string `meddler:"username" json:"username"`
	Content  string `meddler:"content" json:"content"`
}

func DisposeTable(db entity.DB) {
	if _, err := db.Exec("drop table if exists question"); err != nil {
		panic(err)
	}
}

func CreateTableIfNotExists(db entity.DB) {
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS
			question(
				id integer primary key
				, title varchar(40)
				, username varchar(40)
				, content varchar(254)
			)`); err != nil {
		panic(err)
	}
}

func Init(title, username, content string) *Question {
	return &Question{Title: title, Username: username, Content: content}
}

func (q *Question) Insert(db entity.DB) error {
	return meddler.Insert(db, "question", q)
}

func SelectAll(db entity.DB) []*Question {
	var ques []*Question
	err := meddler.QueryAll(db, &ques, "select * from question")
	if err != nil {
		panic(err)
	}
	return ques
}
