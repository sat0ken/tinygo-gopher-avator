// Package avatar provides a TinyGo port of m5stack-avatar.
// It draws an animated face on an M5Stack display (or any compatible display).
package avatar

import (
	"math"
	"math/rand"
	"time"
)

// Avatar manages the animated face and its state.
type Avatar struct {
	face FaceDrawer

	isDrawing bool
	expression Expression
	breath     float32

	rightEyeOpenRatio float32
	rightGazeV        float32
	rightGazeH        float32

	leftEyeOpenRatio float32
	leftGazeV        float32
	leftGazeH        float32

	isAutoBlink bool

	mouthOpenRatio float32
	rotation       float32
	scale          float32
	palette        ColorPalette
	speechText     string
	colorDepth     int
	batteryStatus  BatteryIconStatus
	batteryLevel   int32
}

// NewAvatar creates an Avatar with a default face using the given display.
func NewAvatar(display Displayer) *Avatar {
	return NewAvatarWithFace(NewFace(display))
}

// NewAvatarWithFace creates an Avatar with a custom face (any FaceDrawer).
func NewAvatarWithFace(face FaceDrawer) *Avatar {
	return &Avatar{
		face:              face,
		isDrawing:         false,
		expression:        ExpressionNeutral,
		breath:            0,
		rightEyeOpenRatio: 1.0,
		rightGazeV:        0,
		rightGazeH:        0,
		leftEyeOpenRatio:  1.0,
		leftGazeV:         0,
		leftGazeH:         0,
		isAutoBlink:       true,
		mouthOpenRatio:    0,
		rotation:          0,
		scale:             1.0,
		palette:           NewColorPalette(),
		speechText:        "",
		colorDepth:        1,
		batteryStatus:     BatteryInvisible,
		batteryLevel:      0,
	}
}

// GetFace returns the current face as a FaceDrawer interface.
func (a *Avatar) GetFace() FaceDrawer { return a.face }

// SetFace sets a new face (any FaceDrawer implementation).
func (a *Avatar) SetFace(face FaceDrawer) { a.face = face }

// GetColorPalette returns the current color palette.
func (a *Avatar) GetColorPalette() ColorPalette { return a.palette }

// SetColorPalette sets the color palette.
func (a *Avatar) SetColorPalette(cp ColorPalette) { a.palette = cp }

// GetExpression returns the current expression.
func (a *Avatar) GetExpression() Expression { return a.expression }

// SetExpression sets the facial expression.
func (a *Avatar) SetExpression(exp Expression) { a.expression = exp }

// GetBreath returns the current breath value.
func (a *Avatar) GetBreath() float32 { return a.breath }

// SetBreath sets the breath animation value.
func (a *Avatar) SetBreath(f float32) { a.breath = f }

// SetRotation sets the rotation in radians.
func (a *Avatar) SetRotation(radian float32) { a.rotation = radian }

// SetScale sets the scale factor.
func (a *Avatar) SetScale(scale float32) { a.scale = scale }

// SetPosition sets the top-left position of the face.
func (a *Avatar) SetPosition(top, left int16) {
	a.face.GetBoundingRect().SetPosition(top, left)
}

// SetMouthOpenRatio sets the mouth open ratio (0.0 closed, 1.0 fully open).
func (a *Avatar) SetMouthOpenRatio(ratio float32) { a.mouthOpenRatio = ratio }

// SetEyeOpenRatio sets both eye open ratios.
func (a *Avatar) SetEyeOpenRatio(ratio float32) {
	a.leftEyeOpenRatio = ratio
	a.rightEyeOpenRatio = ratio
}

// SetLeftEyeOpenRatio sets the left eye open ratio.
func (a *Avatar) SetLeftEyeOpenRatio(ratio float32) { a.leftEyeOpenRatio = ratio }

// GetLeftEyeOpenRatio returns the left eye open ratio.
func (a *Avatar) GetLeftEyeOpenRatio() float32 { return a.leftEyeOpenRatio }

// SetRightEyeOpenRatio sets the right eye open ratio.
func (a *Avatar) SetRightEyeOpenRatio(ratio float32) { a.rightEyeOpenRatio = ratio }

// GetRightEyeOpenRatio returns the right eye open ratio.
func (a *Avatar) GetRightEyeOpenRatio() float32 { return a.rightEyeOpenRatio }

// SetIsAutoBlink sets whether auto-blink is enabled.
func (a *Avatar) SetIsAutoBlink(b bool) { a.isAutoBlink = b }

// GetIsAutoBlink returns whether auto-blink is enabled.
func (a *Avatar) GetIsAutoBlink() bool { return a.isAutoBlink }

// SetRightGaze sets the right eye gaze direction.
func (a *Avatar) SetRightGaze(vertical, horizontal float32) {
	a.rightGazeV = vertical
	a.rightGazeH = horizontal
}

// GetRightGaze returns the right eye gaze direction.
func (a *Avatar) GetRightGaze() (vertical, horizontal float32) {
	return a.rightGazeV, a.rightGazeH
}

// SetLeftGaze sets the left eye gaze direction.
func (a *Avatar) SetLeftGaze(vertical, horizontal float32) {
	a.leftGazeV = vertical
	a.leftGazeH = horizontal
}

// GetLeftGaze returns the left eye gaze direction.
func (a *Avatar) GetLeftGaze() (vertical, horizontal float32) {
	return a.leftGazeV, a.leftGazeH
}

// GetGaze returns the mean gaze of both eyes.
func (a *Avatar) GetGaze() (vertical, horizontal float32) {
	return 0.5*a.leftGazeV + 0.5*a.rightGazeV,
		0.5*a.leftGazeH + 0.5*a.rightGazeH
}

// SetSpeechText sets the speech balloon text.
func (a *Avatar) SetSpeechText(text string) { a.speechText = text }

// SetBatteryIcon enables or disables the battery icon.
func (a *Avatar) SetBatteryIcon(show bool) {
	if !show {
		a.batteryStatus = BatteryInvisible
	} else {
		a.batteryStatus = BatteryUnknown
	}
}

// SetBatteryStatus sets battery charging state and level (0-100).
func (a *Avatar) SetBatteryStatus(isCharging bool, level int32) {
	if a.batteryStatus == BatteryInvisible {
		return
	}
	if isCharging {
		a.batteryStatus = BatteryCharging
	} else {
		a.batteryStatus = BatteryDischarging
	}
	a.batteryLevel = level
}

// IsDrawing returns true if the avatar is currently animating.
func (a *Avatar) IsDrawing() bool { return a.isDrawing }

// Draw renders one frame of the avatar.
func (a *Avatar) Draw() {
	rightGaze := NewGaze(a.rightGazeV, a.rightGazeH)
	leftGaze := NewGaze(a.leftGazeV, a.leftGazeH)
	ctx := NewDrawContext(
		a.expression, a.breath, &a.palette,
		rightGaze, a.rightEyeOpenRatio,
		leftGaze, a.leftEyeOpenRatio,
		a.mouthOpenRatio, a.speechText,
		a.rotation, a.scale, a.colorDepth,
		a.batteryStatus, a.batteryLevel,
	)
	a.face.Draw(ctx)
}

// Start begins the animation goroutines.
// colorDepth: 1 for 1-bit (monochrome), 16 for 16-bit (full color)
func (a *Avatar) Start(colorDepth int) {
	if a.isDrawing {
		return
	}
	a.colorDepth = colorDepth
	a.isDrawing = true
	go a.drawLoop()
	go a.facialLoop()
}

// Stop stops the animation goroutines.
func (a *Avatar) Stop() {
	a.isDrawing = false
}

func (a *Avatar) drawLoop() {
	for a.isDrawing {
		a.Draw()
		time.Sleep(10 * time.Millisecond)
	}
}

func (a *Avatar) facialLoop() {
	count := 0
	saccadeInterval := time.Second
	blinkInterval := time.Second
	lastSaccade := time.Now()
	lastBlink := time.Now()
	eyeOpen := true

	for a.isDrawing {
		now := time.Now()

		// Saccade (random gaze shift)
		if now.Sub(lastSaccade) > saccadeInterval {
			vertical := rand.Float32()*2 - 1
			horizontal := rand.Float32()*2 - 1
			a.SetRightGaze(vertical, horizontal)
			a.SetLeftGaze(vertical, horizontal)
			saccadeInterval = time.Duration(500+rand.Intn(2000)) * time.Millisecond
			lastSaccade = now
		}

		// Auto blink
		if a.isAutoBlink {
			if now.Sub(lastBlink) > blinkInterval {
				if eyeOpen {
					a.SetEyeOpenRatio(1.0)
					blinkInterval = time.Duration(2500+rand.Intn(2000)) * time.Millisecond
				} else {
					a.SetEyeOpenRatio(0.0)
					blinkInterval = time.Duration(300+rand.Intn(200)) * time.Millisecond
				}
				eyeOpen = !eyeOpen
				lastBlink = now
			}
		}

		// Breath animation (~30fps cycle)
		count = (count + 1) % 100
		a.breath = float32(math.Sin(float64(count) * 2 * math.Pi / 100.0))

		time.Sleep(33 * time.Millisecond)
	}
}
