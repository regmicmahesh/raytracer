package main

import (
	"fmt"
	"math"
	"os"

	"github.com/regmicmahesh/raytracer"
)

func main() {

	origin := raytracer.NewPoint(0, 0, 0)
	radius := 10
	canvasWidth, canvasHeight := 200, 200

	color := raytracer.NewColor(1, 1, 1)

	canvas := raytracer.NewCanvas(int64(canvasWidth), int64(canvasHeight))
	originX, originY := canvasWidth/2, canvasHeight/2

	canvas.WritePixel(originX, originY, color)

	point := raytracer.NewTranslationMatrix4(0, 0, float64(radius)).ApplyPoint(origin)
	canvas.WritePixel(originX+int(point.X), originY-int(point.Z), color)

	data, err := os.Create("data.ppm")
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 11; i++ {
		point = raytracer.NewRotationYMatrix(math.Pi / 6).ApplyPoint(point)
		fmt.Printf("%d: %+v\n", i, point)
		canvas.WritePixel(originX+int(point.X), originY-int(point.Z), color)
	}

	canvas.ToPPM(data)

}
