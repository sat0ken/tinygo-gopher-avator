package avatar

// Mouth represents a mouth drawable.
type Mouth struct {
	minWidth  int16
	maxWidth  int16
	minHeight int16
	maxHeight int16
}

// NewMouth creates a new Mouth with the given dimensions.
func NewMouth(minWidth, maxWidth, minHeight, maxHeight int16) *Mouth {
	return &Mouth{
		minWidth:  minWidth,
		maxWidth:  maxWidth,
		minHeight: minHeight,
		maxHeight: maxHeight,
	}
}

// Draw draws the mouth on the canvas.
func (m *Mouth) Draw(canvas *Canvas, rect BoundingRect, ctx *DrawContext) {
	var primaryColor uint16
	if ctx.GetColorDepth() == 1 {
		primaryColor = 1
	} else {
		primaryColor = ctx.GetColorPalette().Get(ColorPrimary)
	}

	breath := ctx.GetBreath()
	if breath > 1.0 {
		breath = 1.0
	}
	openRatio := ctx.GetMouthOpenRatio()

	h := m.minHeight + int16(float32(m.maxHeight-m.minHeight)*openRatio)
	w := m.minWidth + int16(float32(m.maxWidth-m.minWidth)*(1-openRatio))
	x := rect.GetLeft() - w/2
	y := rect.GetTop() - h/2 + int16(breath*2)
	canvas.FillRect(x, y, w, h, primaryColor)
}
