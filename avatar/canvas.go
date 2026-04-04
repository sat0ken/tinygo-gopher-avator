package avatar

import (
	"image/color"
	"math"
)

// Displayer is the interface for a display device.
// Compatible with tinygo.org/x/drivers display drivers.
type Displayer interface {
	// FillRectangle fills a rectangle on the display with the given color.
	FillRectangle(x, y, width, height int16, c color.RGBA) error
	// DrawRGBBitmap565 transfers a RGB565 pixel buffer to the display.
	// This enables fast bulk transfer instead of pixel-by-pixel calls.
	DrawRGBBitmap565(x, y int16, data []uint16, w, h int16) error
	// Size returns the display width and height.
	Size() (x, y int16)
}

// Canvas is a software-rendered offscreen buffer.
// It implements drawing primitives similar to M5Canvas (LovyanGFX).
type Canvas struct {
	buf       []uint16
	width     int16
	height    int16
	baseColor uint16
}

// NewCanvas creates a new Canvas with given dimensions.
func NewCanvas(width, height int16) *Canvas {
	buf := make([]uint16, int(width)*int(height))
	return &Canvas{buf: buf, width: width, height: height}
}

func (c *Canvas) Width() int16  { return c.width }
func (c *Canvas) Height() int16 { return c.height }

// SetBaseColor sets the color used for Clear().
func (c *Canvas) SetBaseColor(color uint16) {
	c.baseColor = color
}

// Clear fills the canvas with the base color.
func (c *Canvas) Clear() {
	c.FillSprite(c.baseColor)
}

// FillSprite fills the entire canvas with the given color.
func (c *Canvas) FillSprite(color uint16) {
	for i := range c.buf {
		c.buf[i] = color
	}
}

func (c *Canvas) setPixel(x, y int16, color uint16) {
	if x < 0 || y < 0 || x >= c.width || y >= c.height {
		return
	}
	c.buf[int(y)*int(c.width)+int(x)] = color
}

func (c *Canvas) getPixel(x, y int16) uint16 {
	if x < 0 || y < 0 || x >= c.width || y >= c.height {
		return 0
	}
	return c.buf[int(y)*int(c.width)+int(x)]
}

// FillRect fills a rectangle.
func (c *Canvas) FillRect(x, y, w, h int16, color uint16) {
	for dy := int16(0); dy < h; dy++ {
		for dx := int16(0); dx < w; dx++ {
			c.setPixel(x+dx, y+dy, color)
		}
	}
}

// DrawRect draws a rectangle outline.
func (c *Canvas) DrawRect(x, y, w, h int16, color uint16) {
	for dx := int16(0); dx < w; dx++ {
		c.setPixel(x+dx, y, color)
		c.setPixel(x+dx, y+h-1, color)
	}
	for dy := int16(0); dy < h; dy++ {
		c.setPixel(x, y+dy, color)
		c.setPixel(x+w-1, y+dy, color)
	}
}

// DrawLine draws a line using Bresenham's algorithm.
func (c *Canvas) DrawLine(x0, y0, x1, y1 int16, color uint16) {
	dx := abs16(x1 - x0)
	dy := abs16(y1 - y0)
	sx := int16(1)
	if x0 > x1 {
		sx = -1
	}
	sy := int16(1)
	if y0 > y1 {
		sy = -1
	}
	err := dx - dy
	for {
		c.setPixel(x0, y0, color)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

// FillCircle fills a circle using midpoint circle algorithm.
func (c *Canvas) FillCircle(cx, cy, r int16, color uint16) {
	for y := -r; y <= r; y++ {
		x := int16(math.Sqrt(float64(r*r - y*y)))
		c.FillRect(cx-x, cy+y, 2*x+1, 1, color)
	}
}

// DrawCircle draws a circle outline.
func (c *Canvas) DrawCircle(cx, cy, r int16, color uint16) {
	x := r
	y := int16(0)
	err := int16(0)
	for x >= y {
		c.setPixel(cx+x, cy+y, color)
		c.setPixel(cx+y, cy+x, color)
		c.setPixel(cx-y, cy+x, color)
		c.setPixel(cx-x, cy+y, color)
		c.setPixel(cx-x, cy-y, color)
		c.setPixel(cx-y, cy-x, color)
		c.setPixel(cx+y, cy-x, color)
		c.setPixel(cx+x, cy-y, color)
		y++
		if err <= 0 {
			err += 2*y + 1
		} else {
			x--
			err += 2*(y-x) + 1
		}
	}
}

// FillEllipse fills an ellipse.
func (c *Canvas) FillEllipse(cx, cy, rx, ry int16, color uint16) {
	if rx == 0 || ry == 0 {
		return
	}
	for y := -ry; y <= ry; y++ {
		// x^2/rx^2 + y^2/ry^2 <= 1
		// x <= rx * sqrt(1 - y^2/ry^2)
		ratio := float64(y) / float64(ry)
		x := int16(float64(rx) * math.Sqrt(1.0-ratio*ratio))
		c.FillRect(cx-x, cy+y, 2*x+1, 1, color)
	}
}

// FillTriangle fills a triangle using scanline algorithm.
func (c *Canvas) FillTriangle(x0, y0, x1, y1, x2, y2 int16, color uint16) {
	// Sort vertices by y coordinate
	if y0 > y1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}
	if y0 > y2 {
		x0, x2 = x2, x0
		y0, y2 = y2, y0
	}
	if y1 > y2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}

	if y0 == y2 {
		// Degenerate case: horizontal line
		minX := min16(x0, min16(x1, x2))
		maxX := max16(x0, max16(x1, x2))
		c.FillRect(minX, y0, maxX-minX+1, 1, color)
		return
	}

	for y := y0; y <= y2; y++ {
		var xa, xb int16
		if y < y1 {
			if y1 != y0 {
				xa = x0 + (x1-x0)*int16(y-y0)/int16(y1-y0)
			} else {
				xa = x0
			}
		} else {
			if y2 != y1 {
				xa = x1 + (x2-x1)*int16(y-y1)/int16(y2-y1)
			} else {
				xa = x1
			}
		}
		if y2 != y0 {
			xb = x0 + (x2-x0)*int16(y-y0)/int16(y2-y0)
		} else {
			xb = x0
		}
		if xa > xb {
			xa, xb = xb, xa
		}
		c.FillRect(xa, y, xb-xa+1, 1, color)
	}
}

// PushToDisplay copies the canvas buffer to the display at position (x, y)
// using a single bulk RGB565 transfer for maximum performance.
func (c *Canvas) PushToDisplay(disp Displayer, x, y int16) {
	disp.DrawRGBBitmap565(x, y, c.buf, c.width, c.height)
}

// PushRotateZoom applies rotation and zoom from src canvas into this canvas.
// cx, cy is the center of rotation in the source canvas coordinates.
// This is used to implement the rotation/zoom feature of m5stack-avatar.
func (c *Canvas) PushRotateZoom(src *Canvas, cx, cy int16, rotation, scaleX, scaleY float32) {
	cosR := float32(math.Cos(float64(rotation)))
	sinR := float32(math.Sin(float64(rotation)))

	dstW := c.width
	dstH := c.height
	dstCX := dstW / 2
	dstCY := dstH / 2

	for dy := int16(0); dy < dstH; dy++ {
		for dx := int16(0); dx < dstW; dx++ {
			// Map destination pixel to source pixel
			rx := float32(dx-dstCX) / scaleX
			ry := float32(dy-dstCY) / scaleY
			// Apply inverse rotation
			sx := cosR*rx + sinR*ry + float32(cx)
			sy := -sinR*rx + cosR*ry + float32(cy)
			// Sample from source
			sx16 := int16(sx)
			sy16 := int16(sy)
			if sx16 >= 0 && sx16 < src.width && sy16 >= 0 && sy16 < src.height {
				pixel := src.getPixel(sx16, sy16)
				if pixel != src.baseColor {
					c.setPixel(dx, dy, pixel)
				}
			}
		}
	}
}

func abs16(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}

func min16(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

func max16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}
