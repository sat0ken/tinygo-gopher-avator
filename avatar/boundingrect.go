package avatar

// BoundingRect represents the bounding rectangle for drawing.
type BoundingRect struct {
	top    int16
	left   int16
	width  int16
	height int16
}

// NewBoundingRect creates a BoundingRect with default size (0x0).
func NewBoundingRect(top, left int16) BoundingRect {
	return BoundingRect{top: top, left: left}
}

// NewBoundingRectSize creates a BoundingRect with explicit size.
func NewBoundingRectSize(top, left, width, height int16) BoundingRect {
	return BoundingRect{top: top, left: left, width: width, height: height}
}

func (r *BoundingRect) GetTop() int16    { return r.top }
func (r *BoundingRect) GetLeft() int16   { return r.left }
func (r *BoundingRect) GetWidth() int16  { return r.width }
func (r *BoundingRect) GetHeight() int16 { return r.height }
func (r *BoundingRect) GetRight() int16  { return r.left + r.width }
func (r *BoundingRect) GetBottom() int16 { return r.top + r.height }
func (r *BoundingRect) GetCenterX() int16 { return r.left + r.width/2 }
func (r *BoundingRect) GetCenterY() int16 { return r.top + r.height/2 }

func (r *BoundingRect) SetPosition(top, left int16) {
	r.top = top
	r.left = left
}

func (r *BoundingRect) SetSize(width, height int16) {
	r.width = width
	r.height = height
}
