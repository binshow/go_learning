package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"		// 0. 导入数据库驱动包
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/6/18 7:07 PM
// @description   : 使用 go 连接 mysql
// -------------------------------------------

func initDriver() (db *sql.DB , err error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "root", "bin123456", "127.0.0.1", "3306", "binshow")
	db, err = sql.Open("mysql", url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func insert(db *sql.DB) {
	sql := `INSERT INTO t_user_info(user_id, username, account_id, account_name, gender, phone, avatar) values(?,?,?,?,?,?,?)`
	_, err := db.Exec(sql, "1", "测试用户", "10", "测试账号", 1, "18756978264", "https://moose-plus.oss-cn-shenzhen.aliyuncs.com/avatar.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("添加成功")
}

func update(db *sql.DB) {
	sql := `UPDATE t_user_info SET username = ? WHERE user_id = ?`
	result, err := db.Exec(sql, "修改名字", "1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func delete(db *sql.DB) {
	sql := `DELETE FROM t_user_info WHERE user_id = ?`
	result, err := db.Exec(sql, "1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

// UserInfo 映射binshow库中的 t_user_info 表
type UserInfo struct {
	UserId   string  `json:"user_id"`
	UserName string	  `json:"user_name"`
	Avatar   string   `json:"avatar"`
	Phone    string	  `json:"phone"`
}

func query(db *sql.DB) {
	// 查询 SQL
	sql := "select user_id, avatar, username , phone from t_user_info"
	var userInfo UserInfo
	err := db.QueryRow(sql).Scan(&userInfo.UserId, &userInfo.UserName, &userInfo.Avatar , &userInfo.Phone)
	if err != nil {
		fmt.Println("sacn error :: ", err)
		return
	}
	fmt.Println(userInfo)
}

// 查询多行数据
func queryMulti(db *sql.DB) {
	sql := "select user_id, avatar, username from t_user_info where user_id IN(?,?)"
	rows, err := db.Query(sql, "785919644501544960", "790883082524954600")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 判断结果中有没有下一行数据
	for rows.Next() {
		var userInfo UserInfo
		err = rows.Scan(&userInfo.UserId, &userInfo.UserName, &userInfo.Avatar)
		if err != nil {
			fmt.Println("发送错误")
			return
		}
		fmt.Println(userInfo)
	}
}



func main() {
	//1. 连接mysql
	db , err := initDriver()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//insert(db)
	//update(db)
	query(db)
	//delete(db)





}