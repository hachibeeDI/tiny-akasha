package main

import (
	"database/sql"

	// "html/template"

	_ "github.com/mattn/go-sqlite3"

	"github.com/zenazn/goji"
	// "golang.org/x/net/context"

	// "github.com/hachibeeDI/tiny-akasha/view"
	"github.com/hachibeeDI/tiny-akasha/controller"
	"github.com/hachibeeDI/tiny-akasha/model/entity"
	"github.com/hachibeeDI/tiny-akasha/model/entity/question"
)

func PrePareDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec("drop table if exists questions"); err != nil {
		panic(err)
	}
	if _, err := db.Exec(
		`create table
			questions(
				id integer primary key
				, title varchar(40)
				, username varchar(40)
				, content varchar(254)
			)`); err != nil {
		panic(err)
	}

	questions := []*question.Question{
		question.Init("Haskellをやる", "Daiki", "ところでHaskellってなんですか?"),
		question.Init("Clojureをやる", "Daiki", "ところでClojureってなんですか?"),
	}
	for _, q := range questions {
		_, err := db.Exec("insert into questions(title, username, content) values(?, ?, ?)", q.Title, q.Username, q.Content)
		if err != nil {
			panic(err)
		}
	}
	return db
}

func main() {
	entity.Db = PrePareDB()
	// tmpls := template.Must(template.ParseGlob("./template/*.html"))

	// ctx := context.Background()
	// ctx = db.WithSQL(ctx, "main", dbcon)
	// ctx = view.NewContext(ctx, tmpls)

	// dbを閉じる
	// defer db.CloseSQLAll(ctx)

	// 神コンテキスト！
	// kami.Context = ctx
	controller.InitRoute()
	goji.Serve()
}
