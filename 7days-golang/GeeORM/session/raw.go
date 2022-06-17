package session

import (
	"database/sql"
	"go_learning/7days-golang/GeeORM/clause"
	"go_learning/7days-golang/GeeORM/dialect"
	"go_learning/7days-golang/GeeORM/log"
	"go_learning/7days-golang/GeeORM/schema"
	"strings"
)

// -------------------------------------------
// @file          : raw.go
// @author        : binshow
// @time          : 2022/6/16 9:49 AM
// @description   :
// -------------------------------------------

// Session 用于实现和数据库的交互
type Session struct {
	db 	 	 *sql.DB					// 连接数据库成功之后返回的指针
	dialect  dialect.Dialect
	refTable *schema.Schema				// 操作的那个表
	sql  	 strings.Builder			// sql语句
	sqlVars  []interface{}			// 占位符 对应值
	clause   clause.Clause
}

func New(db *sql.DB , dialect dialect.Dialect) *Session {
	return &Session{
		db:       db,
		dialect:  dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string , values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars , values...)
	return s
}

// 下面三个方法的封装：可以在sql执行完情况 session 中的两个变量，这样session是可以复用的，开启
//					一次会话，可以执行多次 SQL

// Exec 执行对应的sql语句
func (s *Session) Exec() (result sql.Result , err error) {
	defer s.Clear()
	log.Info("session.exec" , s.sql.String() , s.sqlVars)
	if result , err = s.DB().Exec(s.sql.String() , s.sqlVars...); err != nil{
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}