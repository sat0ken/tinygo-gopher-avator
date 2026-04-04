package avatar

// BatteryIcon draws the battery status icon.
type BatteryIcon struct{}

// NewBatteryIcon creates a new BatteryIcon.
func NewBatteryIcon() *BatteryIcon {
	return &BatteryIcon{}
}

func (bi *BatteryIcon) drawBatteryIcon(canvas *Canvas, x, y int16, fgColor, bgColor uint16, batteryStatus BatteryIconStatus, batteryLevel int32) {
	canvas.DrawRect(x, y+5, 5, 5, fgColor)
	canvas.DrawRect(x+5, y, 30, 15, fgColor)
	batteryWidth := int16(30 * float32(batteryLevel) / 100.0)
	canvas.FillRect(x+5+30-batteryWidth, y, batteryWidth, 15, fgColor)
	if batteryStatus == BatteryCharging {
		canvas.FillTriangle(x+20, y, x+15, y+8, x+20, y+8, bgColor)
		canvas.FillTriangle(x+18, y+7, x+18, y+15, x+23, y+7, bgColor)
		canvas.DrawLine(x+20, y, x+15, y+8, fgColor)
		canvas.DrawLine(x+20, y, x+20, y+7, fgColor)
		canvas.DrawLine(x+18, y+15, x+23, y+7, fgColor)
		canvas.DrawLine(x+18, y+8, x+18, y+15, fgColor)
	}
}

// Draw draws the battery icon on the canvas.
func (bi *BatteryIcon) Draw(canvas *Canvas, rect BoundingRect, ctx *DrawContext) {
	if ctx.GetBatteryStatus() == BatteryInvisible {
		return
	}
	var fgColor, bgColor uint16
	if ctx.GetColorDepth() == 1 {
		fgColor = 1
		bgColor = 0
	} else {
		fgColor = ctx.GetColorPalette().Get(ColorPrimary)
		bgColor = ctx.GetColorPalette().Get(ColorBackground)
	}
	bi.drawBatteryIcon(canvas, 285, 5, fgColor, bgColor, ctx.GetBatteryStatus(), ctx.GetBatteryLevel())
}
