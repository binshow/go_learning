package main

import "errors"

// -------------------------------------------
// @file          : error_string.go
// @author        : binshow
// @time          : 2022/6/25 3:57 PM
// @description   :
// -------------------------------------------

type errorString string

func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")
