package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	// "html/template"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/zenazn/goji"
	// "golang.org/x/net/context"

	// "github.com/hachibeeDI/tiny-akasha/view"
	"github.com/hachibeeDI/tiny-akasha/controller"
	"github.com/hachibeeDI/tiny-akasha/model/entity"
	"github.com/hachibeeDI/tiny-akasha/model/entity/question"
)

func makeConnectionString() string {
	flag.Parse()
	host := "localhost"
	port := "3306"
	user := os.Args[1]
	pass := os.Args[1]
	dbname := "auction"
	protocol := "tcp"
	dbargs := " "

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", user, pass, protocol, host, port, dbname, dbargs)
}

func PrePareDB() *sql.DB {
	db, err := sql.Open("mysql", makeConnectionString())
	if err != nil {
		panic(err)
	}

	// question.DisposeTable(db)
	question.CreateTableIfNotExists(db)
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
