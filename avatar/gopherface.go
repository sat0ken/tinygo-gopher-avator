package avatar

type GopherFace struct {
	boundingRect   BoundingRect
	display        Displayer
	rightEyeCanvas *Canvas
	leftEyeCanvas  *Canvas
	baseDrawn      bool
}

const (
	gopherRightEyeCX = int16(104)
	gopherRightEyCY  = int16(78)
	gopherLeftEyeCX  = int16(205)
	gopherLeftEyCY   = int16(73)
	gopherEyeWhiteR  = int16(30)
	gopherPupilR     = int16(10)
	eyeHalf          = int16(44)
	eyeCanvasW       = int16(88)
)

func NewGopherFace(display Displayer) *GopherFace {
	return &GopherFace{
		boundingRect: NewBoundingRectSize(0, 0, GopherBitmapWidth, GopherBitmapHeight),
		display:      display,
	}
}

func (f *GopherFace) GetBoundingRect() *BoundingRect { return &f.boundingRect }

func (f *GopherFace) Draw(ctx *DrawContext) {
	if f.rightEyeCanvas == nil { f.rightEyeCanvas = NewCanvas(eyeCanvasW, eyeCanvasW) }
	if f.leftEyeCanvas == nil  { f.leftEyeCanvas  = NewCanvas(eyeCanvasW, eyeCanvasW) }
	if !f.baseDrawn {
		f.display.DrawRGBBitmap565(0, 0, gopherBitmap[:], GopherBitmapWidth, GopherBitmapHeight)
		f.baseDrawn = true
	}

	// Inline copy to keep bitmaps in Flash (no []uint16 function parameter)
	copy(f.rightEyeCanvas.buf, gopherRightEyeBitmap[:])
	f.drawEye(f.rightEyeCanvas, gopherRightEyeCX, gopherRightEyCY, ctx.GetRightGaze(), ctx.GetRightEyeOpenRatio())

	copy(f.leftEyeCanvas.buf, gopherLeftEyeBitmap[:])
	f.drawEye(f.leftEyeCanvas, gopherLeftEyeCX, gopherLeftEyCY, ctx.GetLeftGaze(), ctx.GetLeftEyeOpenRatio())
}

func (f *GopherFace) drawEye(c *Canvas, cx, cy int16, g Gaze, openRatio float32) {
	localCX, localCY := eyeHalf, eyeHalf
	// Fill only the white-of-eye circle (ring outline stays from bitmap)
	c.FillCircle(localCX, localCY, gopherEyeWhiteR, eyeWhiteRGB565)
	if openRatio <= 0 {
		c.FillRect(localCX-gopherEyeWhiteR, localCY-2, gopherEyeWhiteR*2, 4, pupilRGB565)
	} else {
		maxOffset := float32(gopherEyeWhiteR - gopherPupilR - 2)
		px := localCX + int16(g.GetHorizontal()*maxOffset)
		py := localCY + int16(g.GetVertical()*maxOffset)
		c.FillCircle(px, py, gopherPupilR, pupilRGB565)
		c.FillCircle(px+3, py-3, 3, highlightRGB565)
		if openRatio < 1.0 {
			coverH := int16(float32(gopherEyeWhiteR*2) * (1.0 - openRatio))
			c.FillRect(localCX-gopherEyeWhiteR, localCY-gopherEyeWhiteR, gopherEyeWhiteR*2, coverH, eyeWhiteRGB565)
		}
	}
	ox := cx - eyeHalf
	oy := cy - eyeHalf
	f.display.DrawRGBBitmap565(ox, oy, c.buf, eyeCanvasW, eyeCanvasW)
}
