package main

import (
	"fmt"
	"os"
)

const filePath1 = "binshow.txt"
const filePath2 = "05-GoSDK-study/zkd.txt"

func main() {
	// 拿到 fileInfo
	stat, err := os.Stat(filePath2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf( "%T" , stat)
	fmt.Println(stat.IsDir())
	fmt.Println(stat.Name()) // nil
	fmt.Println(stat.Mode()) // -rw-r--r--
}
