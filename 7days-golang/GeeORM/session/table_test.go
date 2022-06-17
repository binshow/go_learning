package session

import "testing"

// -------------------------------------------
// @file          : table_test.go
// @author        : binshow
// @time          : 2022/6/17 9:53 AM
// @description   :
// -------------------------------------------


type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {

	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}