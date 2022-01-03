package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 创建 context
	ctx, cancel := chromedp.NewContext(context.Background())

	str := getCurrentDirectory()
	// 读取用户输入
	fmt.Println("请输入要爬取的网址地址(http(s)://等开头需要加上): ")
	reader := bufio.NewReader(os.Stdin)

	url, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入地址有误")
		return
	}

	defer cancel()
	// 生成pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(url, &buf)); err != nil {
		fmt.Println("输入地址有误")
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(str+"/output.pdf", buf, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("任务完成,输出文件至:", str)

}

func printToPDF(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url), //浏览指定的页面
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx) // 通过cdp执行PrintToPDF
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
