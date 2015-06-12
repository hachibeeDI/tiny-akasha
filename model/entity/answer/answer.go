package answer

import (
	"errors"

	"github.com/russross/meddler"

	"github.com/hachibeeDI/tiny-akasha/model/entity"
)

type Answer struct {
	ID         int    `meddler:"id,pk" json:"id"`
	QuestionID int    `meddler:"question_id" json:"question_id"`
	Username   string `meddler:"username" json:"username"`
	Content    string `meddler:"content" json:"content"`
}

func DisposeTable(db entity.DB) {
	if _, err := db.Exec("drop table if exists answer"); err != nil {
		panic(err)
	}
}

func CreateTableIfNotExists(db entity.DB) {
	// FOREIGN KEY(question_id) REFERENCES question(id)を入れたいけど、
	// 外部キーに対応したバージョンのMroongaはCentでしか配布されていないので保留
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS
			answer(
				id integer AUTO_INCREMENT primary key
				, question_id integer
				, username varchar(40)
				, content TEXT
			)CHARSET=utf8;`); err != nil {
		panic(err)
	}
}

func Init(questionID int, username, content string) *Answer {
	return &Answer{QuestionID: questionID, Username: username, Content: content}
}

func (q *Answer) Insert(db entity.DB) error {
	if q.Username == "" || q.Content == "" {
		return errors.New("answer's name or content is empty. so create answer is failed.")
	}
	return meddler.Insert(db, "answer", q)
}

func Update(db entity.DB, id int, username, content string) error {
	ans := SelectById(db, id)
	// TODO: 認証っぽい機能
	if ans.Username != username {
		return errors.New("only same user can update")
	}
	ans.Content = content
	return meddler.Update(db, "answer", ans)
}

func Delete(db entity.DB, id int) error {
	result, err := db.Exec("DELETE FROM answer WHERE id = ?", id)
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

func DeleteByQuestionID(db entity.DB, questionID int) error {
	result, err := db.Exec("DELETE FROM answer WHERE question_id = ?", questionID)
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

func SelectAll(db entity.DB) []*Answer {
	var anss []*Answer
	err := meddler.QueryAll(db, &anss, "SELECT * FROM answer")
	if err != nil {
		panic(err)
	}
	return anss
}

func SelectById(db entity.DB, id int) *Answer {
	ans := new(Answer)
	err := meddler.QueryRow(db, ans, "SELECT * FROM answer WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	return ans
}

func SelectByQuestionId(db entity.DB, question_id int) []*Answer {
	var anss []*Answer
	err := meddler.QueryAll(db, &anss, "SELECT * FROM answer WHERE question_id = ?", question_id)
	if err != nil {
		panic(err)
	}
	if anss == nil {
		return []*Answer{}
	}
	return anss
}
