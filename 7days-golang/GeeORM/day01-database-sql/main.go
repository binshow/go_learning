package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go_learning/7days-golang/GeeORM/log"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/6/15 11:43 PM
// @description   : 使用 sqlite 适应 database/sql
// -------------------------------------------

func testsql() {
	//1. 连接数据库，第一个是驱动名称，在上面导包时会注册 sqlite3 的驱动，如果上面的包没有导入就会有问题
	db, _ := sql.Open("sqlite3", "gee.db")
	defer db.Close()

	//2. 执行对应的sql语句
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")

	//3. 占位符防止sql注入
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		fmt.Println(affected)
	}

	//4. 查询一行
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	//5. Scan 获取对应的 column 一列的值
	if err := row.Scan(&name); err == nil {
		//log.Println(name)
		log.Info(name)
	}
}


func main() {
	engine, _ := NewEngine("sqlite3", "gee.db")
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)

}