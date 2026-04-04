package avatar

// Drawable is the interface for face parts that can be drawn.
type Drawable interface {
	Draw(canvas *Canvas, rect BoundingRect, ctx *DrawContext)
}
