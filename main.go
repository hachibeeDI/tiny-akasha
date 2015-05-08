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
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec("drop table if exists question"); err != nil {
		panic(err)
	}
	question.CreateTable(db)
	err = question.Init("Haskellをやる", "Daiki", "ところでHaskellってなんですか?").Insert(db)
	if err != nil {
		panic(err)
	}
	err = question.Init("Clojureをやる", "Daiki", "ところでClojureってなんですか?").Insert(db)
	if err != nil {
		panic(err)
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
