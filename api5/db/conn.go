package db

import (
	"database/sql"
	"fmt"
	"log"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const driverName = "mysql"

// Conn 各repositoryで利用するDB接続(Connection)情報
var Conn *sql.DB

func init() {
	/* ===== データベースへ接続する. ===== */
	// ユーザ
	user := "resns" //os.Getenv("MYSQL_USER")
	// パスワード
	password := "resns" //os.Getenv("MYSQL_PASSWORD")
	// 接続先ホスト
	host := "127.0.0.1" //Mac
	//host := "192.168.99.100" //os.Getenv("MYSQL_HOST")
	// 接続先ポート
	port := "3306" //os.Getenv("MYSQL_PORT")
	// 接続先データベース
	database := "resns_app" //os.Getenv("MYSQL_DATABASE")

	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database
	var err error
	Conn, err = sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}
}
