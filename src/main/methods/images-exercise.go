package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

//还记得之前编写的图片生成器吗？现在来另外编写一个，不过这次将会返回 image.Image 来代替 slice 的数据。
//
//自定义的 Image 类型，要实现必要的方法，并且调用 `pic.ShowImage`。
//
//Bounds 应当返回一个 `image.Rectangle`，例如 `image.Rect(0, 0, w, h)`。
//
//ColorModel 应当返回 `color.RGBAModel`。
//
//At 应当返回一个颜色；在这个例子里，在最后一个图片生成器的值 v 匹配 `color.RGBA{v, v, 255, 255}`。
// 1  新建构造体

type Image struct {
	x0, y0, x1, y1 int
}

// 2 实现官方image的三个方法

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(i.x0, i.y0, i.x1, i.y1)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}

func main() {
	m := Image{0, 0, 200, 200}
	pic.ShowImage(m)
}
