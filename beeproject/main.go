package main

import (
	_ "lhfproject/beeproject/routers"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"lhfproject/beeproject/conf"
	"github.com/astaxie/beego"
)

func main() {
	//insert()
	//update()
	//remove()
	query()
	// 运行时
	beego.Run()
}

//增加数据
func insert() {
	db := conf.DBConnect()
	stmt, err := db.Prepare("INSERT usertable (Name,Password) values (?,?)")
	checkErr(err)
	res, err := stmt.Exec("Mary","zax123456")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

//删除数据
func remove() {
	db := conf.DBConnect()
	stmt, err := db.Prepare("DELETE FROM usertable WHERE Id=?")
	checkErr(err)
	res, err := stmt.Exec(3)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

//更新数据
func update() {
	db := conf.DBConnect()
	stmt, err := db.Prepare("UPDATE usertable SET Name=?,Password=? WHERE Id=?")
	checkErr(err)
	res, err := stmt.Exec("liuhaifeng", "liuhaifeng", 2)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

//查询数据
func query() {
	db := conf.DBConnect()
	rows, err := db.Query("SELECT Id,Name,Password FROM usertable ")
	checkErr(err)

	//普通demo
	for rows.Next() {
		var Id int
		var Name string
		var Password string

		rows.Columns()
		err = rows.Scan(&Id, &Name, &Password)
		checkErr(err)

		fmt.Println(Id)
		fmt.Println(Name)
		fmt.Println(Password)
	}
}

 //出现错误时的处理办法
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}



