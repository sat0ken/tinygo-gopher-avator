package avatar

// Color key indices
const (
	ColorPrimary           = 0
	ColorSecondary         = 1
	ColorBackground        = 2
	ColorBalloonForeground = 3
	ColorBalloonBackground = 4
	colorCount             = 5
)

// RGB565 color constants
const (
	White uint16 = 0xFFFF
	Black uint16 = 0x0000
)

// RGB565 converts 8-bit R, G, B values to RGB565 format.
func RGB565(r, g, b uint8) uint16 {
	return (uint16(r>>3) << 11) | (uint16(g>>2) << 5) | uint16(b>>3)
}

// ColorPalette holds colors for drawing a face.
type ColorPalette struct {
	colors [colorCount]uint16
}

// NewColorPalette creates a default ColorPalette (white on black).
func NewColorPalette() ColorPalette {
	cp := ColorPalette{}
	cp.colors[ColorPrimary] = White
	cp.colors[ColorSecondary] = White
	cp.colors[ColorBackground] = Black
	cp.colors[ColorBalloonForeground] = Black
	cp.colors[ColorBalloonBackground] = White
	return cp
}

// Get returns the color for the given key index.
func (cp *ColorPalette) Get(key int) uint16 {
	if key < 0 || key >= colorCount {
		return 0
	}
	return cp.colors[key]
}

// Set sets the color for the given key index.
func (cp *ColorPalette) Set(key int, value uint16) {
	if key < 0 || key >= colorCount {
		return
	}
	cp.colors[key] = value
}
