package main

import (
	"bytes"
	"strings"
	"testing"
)

// -------------------------------------------
// @file          : string_test.go
// @author        : binshow
// @time          : 2022/6/19 3:16 PM
// @description   : 测试string的拼接
// -------------------------------------------

func Plus(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

func StrBuilder(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func ByteBuffer(n int, str string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
	return buf.String()
}

func PreStrBuilder(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func PreByteBuffer(n int, str string) string {
	buf := new(bytes.Buffer)
	buf.Grow(n * len(str))
	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
	return buf.String()
}

func BenchmarkPlus(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Plus(1000, "string")
	}
}

func BenchmarkStrBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StrBuilder(1000, "string")
	}
}

func BenchmarkByteBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ByteBuffer(1000, "string")
	}
}

func BenchmarkPreStrBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PreStrBuilder(1000, "string")
	}
}

func BenchmarkPreByteBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PreByteBuffer(1000, "string")
	}
}