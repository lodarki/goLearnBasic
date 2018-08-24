package main

import (
	"fmt"
	"image"
)

//Package image 定义了 Image 接口：
//
//type Image interface {
//	ColorModel() color.Model　//ColorModel　返回图片的 color.Model
//	Bounds() Rectangle　　　　//图片中非０color的区域
//	At(x, y int) color.Color　　//返回指定点的像素color.Color
//}
//*注意*：`Bounds` 方法的 Rectangle 返回值实际上是一个 image.Rectangle， 其定义在 image 包中。
//
//（参阅文档了解全部信息。）
//
//color.Color 和 color.Model 也是接口，但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。这些接口和类型由image/color 包定义。
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
