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
				id integer AUTO_INCREMENT primary key
				, title varchar(40)
				, username varchar(40)
				, content MEDIUMTEXT
				, FULLTEXT INDEX(title, content)
			) Engine=Mroonga DEFAULT CHARSET=utf8;`); err != nil {
		panic(err)
	}
}

func Init(title, username, content string) *Question {
	return &Question{Title: title, Username: username, Content: content}
}

func (q *Question) Insert(db entity.DB) error {
	return meddler.Insert(db, "question", q)
}

func Delete(db entity.DB, id int) error {
	result, err := db.Exec("DELETE FROM question WHERE id = ?", id)
	if err != nil {
		return err
	}
	af, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if af != 1 {
		return errors.New("failed to delete")
	}
	return nil
}

func SelectAll(db entity.DB) []*Question {
	var ques []*Question
	err := meddler.QueryAll(db, &ques, "SELECT * FROM question")
	if err != nil {
		panic(err)
	}
	return ques
}

func SelectById(db entity.DB, id int) *Question {
	que := new(Question)
	err := meddler.QueryRow(db, que, "SELECT * FROM question WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	return que
}

// SEEALSO: http://redmine.groonga.org/projects/mroonga/wiki/Mroonga_full_text_Search_in_Boolean_mode_%28jp%29_
func SelectByWord(db entity.DB, word string) []*Question {
	var ques []*Question
	mroongaPredicate := "*DOR " + word
	err := meddler.QueryAll(db, &ques, "SELECT * FROM question WHERE MATCH(title, content) AGAINST(?)", mroongaPredicate)
	if err != nil {
		panic(err)
	}
	return ques
}
