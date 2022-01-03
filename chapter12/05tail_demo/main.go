package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tailf 的用法实例 : 查看日志
func main() {
	fileName := "chapter12/my.log"
	config := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 3,
		}, //从文件的哪个地方开始读
		ReOpen:    true,  // 重新打开
		MustExist: false, // 文件不存在不报错
		Poll:      true,  // 轮询
		Follow:    true,  // 是否跟随
	}

	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed ,err : ", err)
		return
	}

	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen ,filename: %s \n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line: ", line.Text)
	}

}
