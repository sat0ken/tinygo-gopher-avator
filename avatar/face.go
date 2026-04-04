package avatar

// FaceDrawer is the interface for any face that can be drawn by Avatar.
type FaceDrawer interface {
	Draw(ctx *DrawContext)
	GetBoundingRect() *BoundingRect
}

// Face holds all drawable parts and renders them to the display.
type Face struct {
	mouth       Drawable
	eyeR        Drawable
	eyeL        Drawable
	eyeblowR    Drawable
	eyeblowL    Drawable
	mouthPos    BoundingRect
	eyeRPos     BoundingRect
	eyeLPos     BoundingRect
	eyeblowRPos BoundingRect
	eyeblowLPos BoundingRect
	boundingRect BoundingRect
	balloon     *Balloon
	effect      *Effect
	battery     *BatteryIcon
	sprite      *Canvas
	tmpSprite   *Canvas
	display     Displayer
}

// NewFace creates a default Face with standard parts and positions.
func NewFace(display Displayer) *Face {
	br := NewBoundingRectSize(0, 0, 320, 240)
	return &Face{
		mouth:        NewMouth(50, 90, 4, 60),
		mouthPos:     NewBoundingRect(148, 163),
		eyeR:         NewEye(8, false),
		eyeRPos:      NewBoundingRect(93, 90),
		eyeL:         NewEye(8, true),
		eyeLPos:      NewBoundingRect(96, 230),
		eyeblowR:     NewEyeblow(32, 0, false),
		eyeblowRPos:  NewBoundingRect(67, 96),
		eyeblowL:     NewEyeblow(32, 0, true),
		eyeblowLPos:  NewBoundingRect(72, 230),
		boundingRect: br,
		balloon:      NewBalloon(),
		effect:       NewEffect(),
		battery:      NewBatteryIcon(),
		display:      display,
	}
}

// NewFaceCustom creates a Face with custom parts and positions.
func NewFaceCustom(
	mouth Drawable,
	mouthPos BoundingRect,
	eyeR Drawable,
	eyeRPos BoundingRect,
	eyeL Drawable,
	eyeLPos BoundingRect,
	eyeblowR Drawable,
	eyeblowRPos BoundingRect,
	eyeblowL Drawable,
	eyeblowLPos BoundingRect,
	boundingRect BoundingRect,
	display Displayer,
) *Face {
	return &Face{
		mouth:        mouth,
		mouthPos:     mouthPos,
		eyeR:         eyeR,
		eyeRPos:      eyeRPos,
		eyeL:         eyeL,
		eyeLPos:      eyeLPos,
		eyeblowR:     eyeblowR,
		eyeblowRPos:  eyeblowRPos,
		eyeblowL:     eyeblowL,
		eyeblowLPos:  eyeblowLPos,
		boundingRect: boundingRect,
		balloon:      NewBalloon(),
		effect:       NewEffect(),
		battery:      NewBatteryIcon(),
		display:      display,
	}
}

// GetBoundingRect returns the bounding rectangle of the face.
func (f *Face) GetBoundingRect() *BoundingRect {
	return &f.boundingRect
}

// SetMouth sets a custom mouth drawable.
func (f *Face) SetMouth(mouth Drawable) { f.mouth = mouth }

// SetLeftEye sets a custom left eye drawable.
func (f *Face) SetLeftEye(eye Drawable) { f.eyeL = eye }

// SetRightEye sets a custom right eye drawable.
func (f *Face) SetRightEye(eye Drawable) { f.eyeR = eye }

// GetMouth returns the mouth drawable.
func (f *Face) GetMouth() Drawable { return f.mouth }

// GetLeftEye returns the left eye drawable.
func (f *Face) GetLeftEye() Drawable { return f.eyeL }

// GetRightEye returns the right eye drawable.
func (f *Face) GetRightEye() Drawable { return f.eyeR }

// Draw renders the face to the display.
func (f *Face) Draw(ctx *DrawContext) {
	w := f.boundingRect.GetWidth()
	h := f.boundingRect.GetHeight()

	// Create sprite if not exists
	if f.sprite == nil {
		f.sprite = NewCanvas(w, h)
	}

	// Set background color
	var bgColor uint16
	if ctx.GetColorDepth() == 1 {
		bgColor = 0
	} else {
		bgColor = ctx.GetColorPalette().Get(ColorBackground)
	}
	f.sprite.FillSprite(bgColor)

	breath := ctx.GetBreath()
	if breath > 1.0 {
		breath = 1.0
	}
	breathOffset := int16(breath * 3)

	// Draw each part with breath offset
	emptyRect := BoundingRect{}

	rect := f.mouthPos
	rect.SetPosition(rect.GetTop()+breathOffset, rect.GetLeft())
	f.mouth.Draw(f.sprite, rect, ctx)

	rect = f.eyeRPos
	rect.SetPosition(rect.GetTop()+breathOffset, rect.GetLeft())
	f.eyeR.Draw(f.sprite, rect, ctx)

	rect = f.eyeLPos
	rect.SetPosition(rect.GetTop()+breathOffset, rect.GetLeft())
	f.eyeL.Draw(f.sprite, rect, ctx)

	rect = f.eyeblowRPos
	rect.SetPosition(rect.GetTop()+breathOffset, rect.GetLeft())
	f.eyeblowR.Draw(f.sprite, rect, ctx)

	rect = f.eyeblowLPos
	rect.SetPosition(rect.GetTop()+breathOffset, rect.GetLeft())
	f.eyeblowL.Draw(f.sprite, rect, ctx)

	f.balloon.Draw(f.sprite, emptyRect, ctx)
	f.effect.Draw(f.sprite, emptyRect, ctx)
	f.battery.Draw(f.sprite, emptyRect, ctx)

	// Apply rotation and scale, then push to display
	scale := ctx.GetScale()
	rotation := ctx.GetRotation()

	if rotation == 0 && scale == 1.0 {
		// Fast path: no transform
		f.sprite.PushToDisplay(f.display, f.boundingRect.GetLeft(), f.boundingRect.GetTop())
	} else {
		// Transform path using tmp sprite
		if f.tmpSprite == nil {
			f.tmpSprite = NewCanvas(w, h)
		}
		f.tmpSprite.SetBaseColor(bgColor)
		f.tmpSprite.Clear()
		cx := w / 2
		cy := h / 2
		f.tmpSprite.PushRotateZoom(f.sprite, cx, cy, rotation, scale, scale)
		f.tmpSprite.PushToDisplay(f.display, f.boundingRect.GetLeft(), f.boundingRect.GetTop())
	}
}
