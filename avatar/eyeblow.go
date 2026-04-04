package avatar

// Eyeblow represents an eyebrow drawable.
type Eyeblow struct {
	width  int16
	height int16
	isLeft bool
}

// NewEyeblow creates a new Eyeblow with the given dimensions and side.
func NewEyeblow(width, height int16, isLeft bool) *Eyeblow {
	return &Eyeblow{width: width, height: height, isLeft: isLeft}
}

// Draw draws the eyebrow on the canvas.
func (eb *Eyeblow) Draw(canvas *Canvas, rect BoundingRect, ctx *DrawContext) {
	if eb.width == 0 || eb.height == 0 {
		return
	}
	exp := ctx.GetExpression()
	x := rect.GetLeft()
	y := rect.GetTop()

	var primaryColor uint16
	if ctx.GetColorDepth() == 1 {
		primaryColor = 1
	} else {
		primaryColor = ctx.GetColorPalette().Get(ColorPrimary)
	}

	if exp == ExpressionAngry || exp == ExpressionSad {
		var a int16
		if eb.isLeft != (exp == ExpressionSad) {
			a = -1
		} else {
			a = 1
		}
		dx := a * 3
		dy := a * 5

		x1 := x - eb.width/2
		x2 := x1 - dx
		x4 := x + eb.width/2
		x3 := x4 + dx
		y1 := y - eb.height/2 - dy
		y2 := y + eb.height/2 - dy
		y3 := y - eb.height/2 + dy
		y4 := y + eb.height/2 + dy
		canvas.FillTriangle(x1, y1, x2, y2, x3, y3, primaryColor)
		canvas.FillTriangle(x2, y2, x3, y3, x4, y4, primaryColor)
	} else {
		x1 := x - eb.width/2
		y1 := y - eb.height/2
		if exp == ExpressionHappy {
			y1 -= 5
		}
		canvas.FillRect(x1, y1, eb.width, eb.height, primaryColor)
	}
}
