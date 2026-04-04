package avatar

// BatteryIconStatus represents the state of the battery icon.
type BatteryIconStatus int

const (
	BatteryDischarging BatteryIconStatus = iota
	BatteryCharging
	BatteryInvisible
	BatteryUnknown
)

// DrawContext carries all the state needed for drawing a face frame.
type DrawContext struct {
	expression       Expression
	breath           float32
	leftGaze         Gaze
	leftEyeOpenRatio float32
	rightGaze        Gaze
	rightEyeOpenRatio float32
	mouthOpenRatio   float32
	palette          *ColorPalette
	speechText       string
	rotation         float32
	scale            float32
	colorDepth       int
	batteryStatus    BatteryIconStatus
	batteryLevel     int32
}

// NewDrawContext creates a new DrawContext.
func NewDrawContext(
	expression Expression,
	breath float32,
	palette *ColorPalette,
	rightGaze Gaze,
	rightEyeOpenRatio float32,
	leftGaze Gaze,
	leftEyeOpenRatio float32,
	mouthOpenRatio float32,
	speechText string,
	rotation float32,
	scale float32,
	colorDepth int,
	batteryStatus BatteryIconStatus,
	batteryLevel int32,
) *DrawContext {
	return &DrawContext{
		expression:        expression,
		breath:            breath,
		palette:           palette,
		rightGaze:         rightGaze,
		rightEyeOpenRatio: rightEyeOpenRatio,
		leftGaze:          leftGaze,
		leftEyeOpenRatio:  leftEyeOpenRatio,
		mouthOpenRatio:    mouthOpenRatio,
		speechText:        speechText,
		rotation:          rotation,
		scale:             scale,
		colorDepth:        colorDepth,
		batteryStatus:     batteryStatus,
		batteryLevel:      batteryLevel,
	}
}

func (dc *DrawContext) GetExpression() Expression       { return dc.expression }
func (dc *DrawContext) GetBreath() float32              { return dc.breath }
func (dc *DrawContext) GetLeftGaze() Gaze               { return dc.leftGaze }
func (dc *DrawContext) GetLeftEyeOpenRatio() float32    { return dc.leftEyeOpenRatio }
func (dc *DrawContext) GetRightGaze() Gaze              { return dc.rightGaze }
func (dc *DrawContext) GetRightEyeOpenRatio() float32   { return dc.rightEyeOpenRatio }
func (dc *DrawContext) GetMouthOpenRatio() float32      { return dc.mouthOpenRatio }
func (dc *DrawContext) GetColorPalette() *ColorPalette  { return dc.palette }
func (dc *DrawContext) GetSpeechText() string           { return dc.speechText }
func (dc *DrawContext) GetRotation() float32            { return dc.rotation }
func (dc *DrawContext) GetScale() float32               { return dc.scale }
func (dc *DrawContext) GetColorDepth() int              { return dc.colorDepth }
func (dc *DrawContext) GetBatteryStatus() BatteryIconStatus { return dc.batteryStatus }
func (dc *DrawContext) GetBatteryLevel() int32          { return dc.batteryLevel }
