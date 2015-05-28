package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/zenazn/goji"
	// "golang.org/x/net/context"

	// "github.com/hachibeeDI/tiny-akasha/view"
	"github.com/hachibeeDI/tiny-akasha/controller"
	"github.com/hachibeeDI/tiny-akasha/model/entity"
	"github.com/hachibeeDI/tiny-akasha/model/entity/answer"
	"github.com/hachibeeDI/tiny-akasha/model/entity/question"
)

func makeConnectionString(dbname string) string {
	flag.Parse()
	host := "localhost"
	port := "3306"
	user := os.Args[1]
	pass := os.Args[2]
	protocol := "tcp"
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s", user, pass, protocol, host, port, dbname)
}

func PrePareDB() *sql.DB {
	connst := makeConnectionString("")
	fmt.Println(connst)
	mysql, err := sql.Open("mysql", connst)
	if err != nil {
		panic(err)
	}
	mysql.Exec(`CREATE DATABASE IF NOT EXISTS akasha DEFAULT CHARACTER SET utf8; `)
	mysql.Close()

	conToDb := makeConnectionString("akasha")
	fmt.Println(conToDb)
	db, err := sql.Open("mysql", conToDb)
	if err != nil {
		panic(err)
	}

	// question.DisposeTable(db)
	question.CreateTableIfNotExists(db)
	answer.CreateTableIfNotExists(db)
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
