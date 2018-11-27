package conf

import "database/sql"

//连接数据库
func DBConnect() *sql.DB {
	dbName := "goyongtest" //数据库名称
	db, err := sql.Open("mysql", "root:zax123@/"+dbName+"?charset=utf8")
	if err != nil {
		panic(err)
		return nil
	}
	return db
}
