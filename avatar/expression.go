package avatar

// Expression represents the facial expression of the avatar.
type Expression int

const (
	ExpressionHappy   Expression = iota
	ExpressionAngry
	ExpressionSad
	ExpressionDoubt
	ExpressionSleepy
	ExpressionNeutral
)
