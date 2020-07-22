package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

//指向image/color的White只用最后一个包单词
var palette = []color.Color{color.White, color.Black}

//常量是指在程序编译后运行时始终都不会变化的值
//常量声明和变量声明一般都会出现在包级别，所以这些常量在整个包中都是可以共享的，
// 或者你也可以把常量声明定义在函数体内部，那么这种常量就只能在函数体内用
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

//go build lissajous.go
//./lissajous >out.gif 然后浏览器打开out.gif
//func main() {
//	rand.Seed(time.Now().UTC().UnixNano())
//	lissajous(os.Stdout)
//}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	//复合声明，实例化复合类型
	//struct是一组值或者叫字段的集合，不同的类型集合在一个struct可以让我们以一个统一的单元进行处理
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		//struct内部的变量可以以一个点"."来进行访问
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
