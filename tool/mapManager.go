package tool

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"net/http"
)

// 下面是一组生成柏林噪声图像有关的算法函数
// 根据算法生成噪声，加入随机数种子带来更多的随机性
func noise(x, y int64, seedA int64) float64 {
	n := x + y*57 + seedA*7
	m := (n << 13) ^ n
	return (1.0 - float64((m*(m*m*15731+789221)+1376312589)&0x7fffffff)/float64(0x40000000))
}

// 使噪声平滑一些
func smoothedNoise(x float64, y float64, seedA int64) float64 {
	xInt := int64(math.Trunc(x))
	yInt := int64(math.Trunc(y))
	cornersT := (noise(xInt-1, yInt-1, seedA) + noise(xInt+1, yInt-1, seedA) + noise(xInt-1, yInt+1, seedA) + noise(xInt+1, yInt+1, seedA)) / 16
	sidesT := (noise(xInt-1, yInt, seedA) + noise(xInt+1, yInt, seedA) + noise(xInt, yInt-1, seedA) + noise(xInt, yInt+1, seedA)) / 8
	centerT := noise(xInt, yInt, seedA) / 4
	return cornersT + sidesT + centerT
}

// 用于调整的函数
func interpolate(a, b, x float64) float64 {
	c := x * math.Pi
	d := (1 - math.Cos(c)) * 0.5
	return a*(1-d) + b*d
}

// 进一步调整噪声
func interpolatedNoise(x, y float64, seedA int64) float64 {
	xInt := math.Trunc(x)
	xFrac := x - xInt
	yInt := math.Trunc(y)
	yFrac := y - yInt
	v1 := smoothedNoise(xInt, yInt, seedA)
	v2 := smoothedNoise(xInt+1, yInt, seedA)
	v3 := smoothedNoise(xInt, yInt+1, seedA)
	v4 := smoothedNoise(xInt+1, yInt+1, seedA)
	i1 := interpolate(v1, v2, xFrac)
	i2 := interpolate(v3, v4, xFrac)
	return interpolate(i1, i2, yFrac)
}

// PerlinNoise2D 根据指定参数生成二维的柏林噪声
func PerlinNoise2D(x, y float64, seedA int64, alphaA float64, betaA float64, octavesA int) float64 {
	// 存放结果的变量
	resultT := 0.0
	// 放大系数
	scaleT := 1.0
	// 循环调整octavesA次
	for i := 0; i < octavesA; i++ {
		resultT += interpolatedNoise(x, y, seedA) / scaleT
		// *= 操作符与 += 类似，表示 scaleT = scaleT * alphaA
		scaleT *= alphaA
		x *= betaA
		y *= betaA
	}
	return resultT
}

// GetPerlinNoise2DColor 用于计算坐标为 (x, y)点的颜色
func GetPerlinNoise2DColor(x, y float64, alphaA, betaA float64, octavesA int, seedA int64, currentColorA color.Color) color.Color {
	// 先计算噪声（即变化量）
	noiseT := PerlinNoise2D(x, y, seedA, alphaA, betaA, octavesA)
	// 获取当前颜色，准备进行改变
	r, g, b, a := currentColorA.RGBA()
	r = r % 256
	g = g % 256
	b = b % 256
	a = a % 256
	// 放大系数
	scaleT := 9.1
	// 用取模法随机确定修改RGB三原色中的哪个
	rgbSelectorT := int(noiseT*scaleT/100) % 3
	// 根据选择来对某一种原色进行变化
	switch rgbSelectorT {
	case 0:
		r = r + uint32(noiseT*scaleT)
		r = r % 256
	case 1:
		g = g + uint32(noiseT*scaleT)
		g = g % 256
	case 2:
		b = b + uint32(noiseT*scaleT)
		b = b % 256
	}
	return &color.RGBA{byte(r), byte(g), byte(b), byte(a)}
}

// 处理根路径请求的函数，将返回一个随机图片
func HandleImage(respA http.ResponseWriter, reqA *http.Request) {
	// 设置图片的大小
	widthT := 1027.0
	heightT := 768.0
	// 设置生成柏林噪声的参数，再加入一些随机性
	alphaT := 0.31 + rand.Float64()*0.2 - 0.1
	betaT := 0.22 + rand.Float64()*0.2 - 0.1
	octavesT := 3 + rand.Intn(10)
	seedT := rand.Int()
	// 准备新的图片
	imageT := image.NewNRGBA(image.Rect(0, 0, int(widthT), int(heightT)))
	// 准备存放不断变化颜色的colorT变量，并确定一个初始颜色startColorT
	var colorT color.Color
	colorT = color.RGBA{byte(rand.Int() % 256), byte(rand.Int() % 256), byte(rand.Int() % 256), 0xFF}
	startColorT := colorT
	// 循环计算并设置每个点的颜色
	for i := 0.0; i < heightT; i++ {
		for j := 0.0; j < widthT; j++ {
			// 变化的颜色根据柏林噪声基于初始颜色进行变化
			colorT = GetPerlinNoise2DColor(j, i, alphaT, betaT, octavesT, int64(seedT), startColorT)
			// 设置点的颜色
			imageT.Set(int(j), int(i), colorT)
		}
	}
	// 写入网页响应头中的内容类型，表示是png格式的图片
	respA.Header().Set("Content-Type", "image/png")
	// 进行png格式的图形编码，并以流式方法写入http响应中
	png.Encode(respA, imageT)
}

// func createMap() {
// 	// 初始化随机数种子
// 	rand.Seed(time.Now().Unix())
// 	// 设定根路由处理函数
// 	http.HandleFunc("/", HandleImage)
// 	// 在指定端口上监听
// 	http.ListenAndServe(":8837", nil)
// }
