package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Canvas struct {
	width  int
	height int
	pixels [][]Color
}

const MAX_COLORS = 255

// Adressing is NOT like in the "usual" matrices with rows and columns
func newCanvas(width int, height int) Canvas {
	canvas := Canvas{width: width, height: height}
	canvas.pixels = make([][]Color, width)
	for i := 0; i < width; i++ {
		canvas.pixels[i] = make([]Color, height)
	}
	return canvas
}

func (c *Canvas) WritePixel(x int, y int, color Color) {
	c.pixels[x][y] = color
}

func (c *Canvas) PixelAt(x int, y int) Color {
	return c.pixels[x][y]
}

func (c *Canvas) Fill(color Color) {
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			c.WritePixel(x, y, color)
		}
	}
}

func (c *Canvas) ppmHeader() string {
	MAGIC_NUMBER := "P3"
	header := fmt.Sprintf("%s\n%d %d\n%d\n", MAGIC_NUMBER, c.width, c.height, MAX_COLORS)

	return header
}

func (c *Canvas) ppmPixelData() string {
	const MAX_PPM_LINE_LEN = 70
	result := ""
	for y := 0; y < c.height; y++ {
		line := ""
		for x := 0; x < c.width; x++ {
			rgbComponents := c.PixelAt(x, y).ToScaledRgbComponents()
			for _, component := range rgbComponents {
				componentString := strconv.Itoa(component)
				if len(line)+len(componentString) >= MAX_PPM_LINE_LEN {
					line = strings.Trim(line, " ")
					result += line + "\n"
					line = ""
				}
				line += componentString + " "
			}
		}
		line = strings.Trim(line, " ")
		result += line + "\n"
	}
	return result
}

func (c *Canvas) PpmData() string {
	return c.ppmHeader() + c.ppmPixelData()
}
