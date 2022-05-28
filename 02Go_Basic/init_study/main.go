package main

import (
	"image"
	_ "image/gif" // 空导包的形式
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	// 支持 png 、 jpeg 、 gif 三种格式的图片
	//width, height, err := imageSize(os.Args[1])
	//if err != nil {
	//	fmt.Println("get image size err: " , err)
	//	return
	//}
	//fmt.Printf("image size : %d , %d \n" , width , height)
}

func imageSize(imageFile string) (int, int, error) {
	// 打开图文文件
	f , _ := os.Open(imageFile)
	defer f.Close()

	// 对文件解码，得到图片实例
	image, _, err := image.Decode(f)
	if err != nil {
		return 0 , 0 , err
	}

	// 返回图片区域
	b := image.Bounds()
	return b.Max.X , b.Max.Y , nil
}