package avatar

import "math"

// Effect draws expression-specific visual effects.
type Effect struct{}

// NewEffect creates a new Effect.
func NewEffect() *Effect {
	return &Effect{}
}

func (e *Effect) drawBubbleMark(canvas *Canvas, x, y, r int16, color uint16, offset float32) {
	rr := r + int16(float32(r)*0.2*offset)
	canvas.DrawCircle(x, y, rr, color)
	canvas.DrawCircle(x-rr/4, y-rr/4, rr/4, color)
}

func (e *Effect) drawSweatMark(canvas *Canvas, x, y, r int16, color uint16, offset float32) {
	y += int16(5 * offset)
	r += int16(float32(r) * 0.2 * offset)
	canvas.FillCircle(x, y, r, color)
	a := int16(float64(r) * math.Sqrt(3) / 2)
	canvas.FillTriangle(x, y-r*2, x-a, y-r/2, x+a, y-r/2, color)
}

func (e *Effect) drawChillMark(canvas *Canvas, x, y, r int16, color uint16, offset float32) {
	h := r + int16(math.Abs(float64(float32(r)*0.2*offset)))
	canvas.FillRect(x-r/2, y, 3, h/2, color)
	canvas.FillRect(x, y, 3, h*3/4, color)
	canvas.FillRect(x+r/2, y, 3, h, color)
}

func (e *Effect) drawAngerMark(canvas *Canvas, x, y, r int16, color, bColor uint16, offset float32) {
	r += int16(math.Abs(float64(float32(r) * 0.4 * offset)))
	canvas.FillRect(x-r/3, y-r, r*2/3, r*2, color)
	canvas.FillRect(x-r, y-r/3, r*2, r*2/3, color)
	canvas.FillRect(x-r/3+2, y-r, r*2/3-4, r*2, bColor)
	canvas.FillRect(x-r, y-r/3+2, r*2, r*2/3-4, bColor)
}

func (e *Effect) drawHeartMark(canvas *Canvas, x, y, r int16, color uint16, offset float32) {
	r += int16(float32(r) * 0.4 * offset)
	canvas.FillCircle(x-r/2, y, r/2, color)
	canvas.FillCircle(x+r/2, y, r/2, color)
	a := int16(float64(r) * math.Sqrt(2) / 4.0)
	canvas.FillTriangle(x, y, x-r/2-a, y+a, x+r/2+a, y+a, color)
	canvas.FillTriangle(x, y+r/2+2*a, x-r/2-a, y+a, x+r/2+a, y+a, color)
}

// Draw draws the effect on the canvas.
func (e *Effect) Draw(canvas *Canvas, rect BoundingRect, ctx *DrawContext) {
	var primaryColor, bgColor uint16
	if ctx.GetColorDepth() == 1 {
		primaryColor = 1
		bgColor = 0
	} else {
		primaryColor = ctx.GetColorPalette().Get(ColorPrimary)
		bgColor = ctx.GetColorPalette().Get(ColorBackground)
	}
	offset := ctx.GetBreath()
	exp := ctx.GetExpression()

	switch exp {
	case ExpressionDoubt:
		e.drawSweatMark(canvas, 290, 110, 7, primaryColor, -offset)
	case ExpressionAngry:
		e.drawAngerMark(canvas, 280, 50, 12, primaryColor, bgColor, offset)
	case ExpressionHappy:
		e.drawHeartMark(canvas, 280, 50, 12, primaryColor, offset)
	case ExpressionSad:
		e.drawChillMark(canvas, 270, 0, 30, primaryColor, offset)
	case ExpressionSleepy:
		e.drawBubbleMark(canvas, 290, 40, 10, primaryColor, offset)
		e.drawBubbleMark(canvas, 270, 52, 6, primaryColor, -offset)
	}
}
