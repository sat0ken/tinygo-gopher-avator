//go:build tinygo && m5stack

// Basic example for tinygo-avatar on M5Stack Core (ILI9341/ILI9342C 320x240).
// Build with: tinygo flash -target m5stack ./examples/basics/
package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/notnow/tinygo-avatar/avatar"
	"tinygo.org/x/drivers/ili9341"
)

// ili9341Wrapper wraps ili9341.Device to implement avatar.Displayer.
type ili9341Wrapper struct {
	dev *ili9341.Device
}

func (d *ili9341Wrapper) FillRectangle(x, y, width, height int16, c color.RGBA) error {
	return d.dev.FillRectangle(x, y, width, height, c)
}

func (d *ili9341Wrapper) DrawRGBBitmap565(x, y int16, data []uint16, w, h int16) error {
	return d.dev.DrawRGBBitmap(x, y, data, w, h)
}

func (d *ili9341Wrapper) Size() (int16, int16) {
	return d.dev.Size()
}

func main() {
	machine.SPI2.Configure(machine.SPIConfig{
		SCK:       machine.SPI0_SCK_PIN,
		SDO:       machine.SPI0_SDO_PIN,
		SDI:       machine.SPI0_SDI_PIN,
		Frequency: 40e6,
	})

	// Configure backlight
	backlight := machine.LCD_BL_PIN
	backlight.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Initialize ILI9341/ILI9342C display
	display := ili9341.NewSPI(
		machine.SPI2,
		machine.LCD_DC_PIN,
		machine.LCD_SS_PIN,
		machine.LCD_RST_PIN,
	)
	display.Configure(ili9341.Config{
		Width:            320,
		Height:           240,
		DisplayInversion: displayInversion,
	})
	backlight.High()
	display.SetRotation(ili9341.Rotation0Mirror)
	display.FillScreen(color.RGBA{210, 210, 210, 255})

	// Create and start the avatar with Gopher face
	disp := &ili9341Wrapper{dev: display}
	gopherFace := avatar.NewGopherFace(disp)
	av := avatar.NewAvatarWithFace(gopherFace)
	av.Start(16) // 16-bit color

	// Change expressions every 3 seconds
	expressions := []avatar.Expression{
		avatar.EyeShapeNormal,
		avatar.EyeShapeHalfOpen,
		avatar.EyeShapeOuterSlant,
		avatar.EyeShapeInnerSlant,
		avatar.EyeShapeNormal,
		avatar.EyeShapeHalfClosed,
	}

	i := 0
	for {
		av.SetExpression(expressions[i])
		i = (i + 1) % len(expressions)
		time.Sleep(3 * time.Second)
	}
}
