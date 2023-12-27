package raytracer

import (
	"fmt"
	"io"
)

type Canvas struct {
	Width  int64
	Height int64
	Data   [][]*Color
}

func NewCanvas(w, h int64) *Canvas {

	data := make([][]*Color, 0, h)

	for colNum := int64(0); colNum < h; colNum++ {
		row := make([]*Color, 0, w)
		for rowNum := int64(0); rowNum < w; rowNum++ {
			row = append(row, BlackColor)
		}
		data = append(data, row)

	}

	return &Canvas{
		Width:  w,
		Height: h,
		Data:   data,
	}

}

func (c *Canvas) WritePixel(x, y int, co *Color) {

	if y > int(c.Height)-1 || y < 0 {
		return
	}

	if x > int(c.Width)-1 || x < 0 {
		return
	}

	c.Data[y][x] = co
}

func (c *Canvas) GetPixel(x, y int) *Color {
	return c.Data[y][x]
}

const PPMHeader = `P3
%d %d
255
`

func (c *Canvas) ToPPMString() string {

	ppmString := ""

	ppmString += fmt.Sprintf(PPMHeader, c.Width, c.Height)

	for y := range c.Data {

		for x := range c.Data[y] {

			color := c.Data[y][x]
			colorClamped := color.Clamp(255)

			ppmString += fmt.Sprintf("%d %d %d\n", int64(colorClamped.R), int64(colorClamped.G), int64(colorClamped.B))
		}

	}

	return ppmString

}

func (c *Canvas) ToPPM(w io.Writer) {

	fmt.Fprintf(w, PPMHeader, c.Width, c.Height)

	for y := range c.Data {

		for x := range c.Data[y] {

			color := c.Data[y][x]
			colorClamped := color.Clamp(255)

			fmt.Fprintf(w, "%d %d %d\n", int64(colorClamped.R), int64(colorClamped.G), int64(colorClamped.B))
		}

	}
}
