package avatar

const (
	balloonCX = 240
	balloonCY = 220
)

// Balloon represents a speech balloon drawable.
type Balloon struct{}

// NewBalloon creates a new Balloon.
func NewBalloon() *Balloon {
	return &Balloon{}
}

// Draw draws the speech balloon on the canvas.
func (b *Balloon) Draw(canvas *Canvas, rect BoundingRect, ctx *DrawContext) {
	text := ctx.GetSpeechText()
	if len(text) == 0 {
		return
	}
	cp := ctx.GetColorPalette()
	primaryColor := cp.Get(ColorBalloonForeground)
	backgroundColor := cp.Get(ColorBalloonBackground)

	// Estimate text width (approximate: 8 pixels per character at scale 2)
	textWidth := int16(len(text)) * 8 * 2
	textHeight := int16(8 * 2)

	cx := int16(balloonCX)
	cy := int16(balloonCY)

	canvas.FillEllipse(cx-20, cy, textWidth/2+2, textHeight+2, primaryColor)
	canvas.FillTriangle(cx-62, cy-42, cx-8, cy-10, cx-41, cy-8, primaryColor)
	canvas.FillEllipse(cx-20, cy, textWidth/2, textHeight, backgroundColor)
	canvas.FillTriangle(cx-60, cy-40, cx-10, cy-10, cx-40, cy-10, backgroundColor)
	// Note: Text rendering is not implemented (requires font support).
	// To add text, use a font library compatible with TinyGo.
}
