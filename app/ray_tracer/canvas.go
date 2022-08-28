package ray_tracer

import (
	"fmt"
	"math"
	"os"
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
	if (x < 0 || x >= c.width) || (y < 0 || y >= c.height) {
		fmt.Printf("WARN: pixel coordinate out of bound [%d, %d] with value x: %d, y: %d\n", c.width, c.height, x, y)
		return
	}

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

func (c *Canvas) ToCanvasCoordinates(x float64, y float64) (int, int) {
	canvas_x, canvas_y := int(math.Round(x)), int(math.Round(y))

	// Since Y coordinate in canvas is upside-down, we need to convert
	// "world" Y coordinate to canvas Y coordinate
	canvas_y = c.height - canvas_y
	return canvas_x, canvas_y
}

func (c *Canvas) ppmHeader() string {
	MAGIC_NUMBER := "P3"
	header := fmt.Sprintf("%s\n%d %d\n%d\n", MAGIC_NUMBER, c.width, c.height, MAX_COLORS)

	return header
}

func (c *Canvas) ppmPixelData() string {
	const MAX_PPM_LINE_LEN = 70
	var result strings.Builder
	for y := 0; y < c.height; y++ {
		line := ""
		for x := 0; x < c.width; x++ {
			rgbComponents := c.PixelAt(x, y).ToScaledRgbComponents()
			for _, component := range rgbComponents {
				componentString := strconv.Itoa(component)
				if len(line)+len(componentString) >= MAX_PPM_LINE_LEN {
					line = strings.Trim(line, " ")
					result.WriteString(line + "\n")
					line = ""
				}
				line += componentString + " "
			}
		}
		line = strings.Trim(line, " ")
		result.WriteString(line + "\n")
	}

	return result.String()
}

func (c *Canvas) PpmData() string {
	return c.ppmHeader() + c.ppmPixelData()
}

func (c *Canvas) SavePpm(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic("Failed to create file!")
	}

	f.WriteString(c.PpmData())
	fmt.Printf("Succsessfully written \"%s\"\n", filename)
}
