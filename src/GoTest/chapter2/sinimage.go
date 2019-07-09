package chapter2

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

// 图片大大小
const size  = 300

func main() {
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值范围在0-2PI之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素，向下偏移一般像素并翻转
		y := size/2 - math.Sin(s)*size/2
        // 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})

	}

	file, err := os.Create("C:\\Users\\Administrator\\Desktop\\oeasy\\data\\mine\\Project\\Go_Space\\src\\GoTest\\two\\sin.png")

	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, pic)

	file.Close()


}