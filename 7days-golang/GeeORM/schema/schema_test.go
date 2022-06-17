package schema

import (
	"go_learning/7days-golang/GeeORM/dialect"
	"testing"
)

// -------------------------------------------
// @file          : schema_test.go
// @author        : binshow
// @time          : 2022/6/17 9:49 AM
// @description   :
// -------------------------------------------


type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}