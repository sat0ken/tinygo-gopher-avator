package avatar

// Eye represents a single eye drawable.
type Eye struct {
	r      int16
	isLeft bool
}

// NewEye creates a new Eye with the given radius and side.
func NewEye(r int16, isLeft bool) *Eye {
	return &Eye{r: r, isLeft: isLeft}
}

// Draw draws the eye on the canvas.
func (e *Eye) Draw(canvas *Canvas, rect BoundingRect, ctx *DrawContext) {
	exp := ctx.GetExpression()
	x := rect.GetCenterX()
	y := rect.GetCenterY()

	var g Gaze
	var openRatio float32
	if e.isLeft {
		g = ctx.GetLeftGaze()
		openRatio = ctx.GetLeftEyeOpenRatio()
	} else {
		g = ctx.GetRightGaze()
		openRatio = ctx.GetRightEyeOpenRatio()
	}

	offsetX := int16(g.GetHorizontal() * 3)
	offsetY := int16(g.GetVertical() * 3)

	var primaryColor, backgroundColor uint16
	if ctx.GetColorDepth() == 1 {
		primaryColor = 1
		backgroundColor = 0
	} else {
		primaryColor = ctx.GetColorPalette().Get(ColorPrimary)
		backgroundColor = ctx.GetColorPalette().Get(ColorBackground)
	}

	if openRatio > 0 {
		canvas.FillCircle(x+offsetX, y+offsetY, e.r, primaryColor)

		if exp == EyeShapeInnerSlant || exp == EyeShapeOuterSlant {
			x0 := x + offsetX - e.r
			y0 := y + offsetY - e.r
			x1 := x0 + e.r*2
			y1 := y0
			var x2 int16
			if (!e.isLeft) != (exp == EyeShapeOuterSlant) {
				x2 = x0
			} else {
				x2 = x1
			}
			y2 := y0 + e.r
			canvas.FillTriangle(x0, y0, x1, y1, x2, y2, backgroundColor)
		}

		if exp == EyeShapeHalfOpen || exp == EyeShapeHalfClosed {
			x0 := x + offsetX - e.r
			y0 := y + offsetY - e.r
			w := e.r*2 + 4
			h := e.r + 2
			if exp == EyeShapeHalfOpen {
				y0 += e.r
				canvas.FillCircle(x+offsetX, y+offsetY, e.r*2/3, backgroundColor)
			}
			canvas.FillRect(x0, y0, w, h, backgroundColor)
		}
	} else {
		x1 := x - e.r + offsetX
		y1 := y - 2 + offsetY
		w := e.r * 2
		h := int16(4)
		canvas.FillRect(x1, y1, w, h, primaryColor)
	}
}
