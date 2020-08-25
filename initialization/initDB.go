package initialization

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var Db *sql.DB
const DbName = "goBlog"

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "foo.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt, err := Db.Prepare("CREATE TABLE IF NOT EXISTS `"+ DbName +"` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `email` VARCHAR(64) NULL, `password` VARCHAR(128) NULL, `nickname` VARCHAR(128) NULL)")
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
	// 最大连接数
	Db.SetMaxOpenConns(20)
	// 最大空闲连接数
	Db.SetMaxIdleConns(20)
}
