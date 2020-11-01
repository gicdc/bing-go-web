package utils

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

//init函数先于main函数执行，初始化操作
func init() {
	var err error
	//	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	Db, err = sqlx.Open(`sqlserver`,
		`server=127.0.0.1;user id=sa;password=;port=1433;database=bings2;`)
	if err != nil {
		panic("连接错误")
	}
	if err = Db.Ping(); err != nil {
		panic("运行错误")
	}
}
