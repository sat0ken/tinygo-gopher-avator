package avatar

// Gaze represents the direction of the eye gaze.
type Gaze struct {
	vertical   float32
	horizontal float32
}

// NewGaze creates a new Gaze with vertical and horizontal values.
func NewGaze(vertical, horizontal float32) Gaze {
	return Gaze{vertical: vertical, horizontal: horizontal}
}

// GetVertical returns the vertical component of the gaze.
func (g Gaze) GetVertical() float32 {
	return g.vertical
}

// GetHorizontal returns the horizontal component of the gaze.
func (g Gaze) GetHorizontal() float32 {
	return g.horizontal
}
