package raytracer

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// indexing starts at 1
func getLine(lines string, lineNumber int) string {

	linesSlice := strings.Split(lines, "\n")
	if lineNumber < 1 || lineNumber > len(lines) {
		return ""
	}
	return linesSlice[lineNumber-1]

}

func TestCanvas(t *testing.T) {

	t.Run("NewCanvas() must create a new canvas", func(t *testing.T) {

		tests := []struct {
			height int64
			width  int64
		}{
			{0, 0},
			{10, 12},
			{10, 10},
			{50, 60},
		}

		for _, v := range tests {

			x := NewCanvas(v.width, v.height)

			assert.Equal(t, x.Height, v.height)
			assert.Equal(t, x.Width, v.width)

		}

	})

	t.Run("length of data must be equal to height", func(t *testing.T) {

		width := 800
		height := 600

		canvas := NewCanvas(int64(width), int64(height))

		assert.Equal(t, len(canvas.Data), height)

	})

	t.Run("length of data's element must be equal to width", func(t *testing.T) {

		width := 800
		height := 600

		canvas := NewCanvas(int64(width), int64(height))

		assert.Equal(t, len(canvas.Data[0]), width)
	})

	t.Run("WritePixel must write the pixel at given coords", func(t *testing.T) {

		width := 800
		height := 600

		canvas := NewCanvas(int64(width), int64(height))

		color := NewColor(255, 128, 0)

		canvas.WritePixel(100, 120, color)

		assert.Equal(t, canvas.Data[120][100].R, color.R)
		assert.Equal(t, canvas.Data[120][100].G, color.G)
		assert.Equal(t, canvas.Data[120][100].B, color.B)

	})

	t.Run("GetPixel must write the pixel at given coords", func(t *testing.T) {

		width := 800
		height := 600

		canvas := NewCanvas(int64(width), int64(height))

		color := NewColor(255, 128, 0)

		canvas.WritePixel(100, 120, color)

		pixel := canvas.GetPixel(100, 120)

		assert.Equal(t, pixel.R, color.R)
		assert.Equal(t, pixel.G, color.G)
		assert.Equal(t, pixel.B, color.B)

	})

	t.Run("ToPPM must return a PPM file format", func(t *testing.T) {

		width := 5
		height := 2

		canvas := NewCanvas(int64(width), int64(height))

		ppmData := canvas.ToPPMString()

		firstLine := getLine(ppmData, 1)

		secondLine := getLine(ppmData, 2)
		expectedSecondLine := fmt.Sprintf("%d %d", width, height)

		thirdLine := getLine(ppmData, 3)
		expectedThirdLine := "255"

		fourthLine := getLine(ppmData, 4)
		expectedFourthLine := "0 0 0"

		// format is 3 line for headers + data + 1 empty line
		expectedLineCount := 1 + 1 + 1 + (width * height) + 1

		assert.Equal(t, firstLine, "P3")
		assert.Equal(t, secondLine, expectedSecondLine)
		assert.Equal(t, thirdLine, expectedThirdLine)
		assert.Equal(t, fourthLine, expectedFourthLine)
		assert.Equal(t, len(strings.Split(ppmData, "\n")), expectedLineCount)

	})

}
