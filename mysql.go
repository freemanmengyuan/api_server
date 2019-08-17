package main

import (
	"database/sql"  //相关扩展
	"fmt"
	// 引入数据库驱动注册及初始化
	_"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:469312@tcp(144.34.144.200:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	//新增
	/*smt, err := db.Prepare("insert into user_info set name=?, department=?, create_time=?")
	if err != nil {
		panic(err)
	}
	res, err := smt.Exec("lisi", "技术部", "2019-09-20")
	if err != nil {
		panic(err)
	}
	id,err := res.LastInsertId()*/

	//修改
	smt, err := db.Prepare("update user_info set name=?  where id=?")
	if err != nil {
		panic(err)
	}
	res, err := smt.Exec("wangwu", 2)
	if err != nil {
		panic(err)
	}
	id,err := res.RowsAffected()

	fmt.Print(id)
}
