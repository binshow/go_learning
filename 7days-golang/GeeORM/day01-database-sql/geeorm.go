package main

import (
	"database/sql"
	"go_learning/7days-golang/GeeORM/dialect"
	"go_learning/7days-golang/GeeORM/log"
	"go_learning/7days-golang/GeeORM/session"
)

// -------------------------------------------
// @file          : geeorm.go
// @author        : binshow
// @time          : 2022/6/16 10:01 AM
// @description   :	负责交互前和交互后的一些操作
// -------------------------------------------


type Engine struct {
	db 		*sql.DB
	dialect dialect.Dialect		// 适配不同的数据库，屏蔽数据库差异
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db , engine.dialect)
}