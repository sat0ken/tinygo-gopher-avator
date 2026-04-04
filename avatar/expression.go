package avatar

// Expression represents the eye and eyebrow shape of the avatar's face.
type Expression int

const (
	// EyeShapeNormal draws eyes as plain circles and eyebrows as horizontal bars.
	EyeShapeNormal Expression = iota
	// EyeShapeHalfOpen draws eyes with the lower half covered and a small inner circle, eyebrows raised.
	EyeShapeHalfOpen
	// EyeShapeHalfClosed draws eyes with the lower half covered and eyebrows as horizontal bars.
	EyeShapeHalfClosed
	// EyeShapeInnerSlant cuts the inner-upper corner of each eye with a triangle and tilts eyebrows inward.
	EyeShapeInnerSlant
	// EyeShapeOuterSlant cuts the outer-upper corner of each eye with a triangle and tilts eyebrows outward.
	EyeShapeOuterSlant
)
