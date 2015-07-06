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
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	if DB_HOST == "" {
		DB_HOST = "localhost"
	}
	if DB_PORT == "" {
		DB_PORT = "3306"
	}
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASS := os.Getenv("MYSQL_PASS")
	if MYSQL_USER == "" {
		MYSQL_USER = "root"
	}
	if MYSQL_PASS == "" {
		MYSQL_PASS = "password"
	}
	protocol := "tcp"
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s", MYSQL_USER, MYSQL_PASS, protocol, DB_HOST, DB_HOST, dbname)
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
