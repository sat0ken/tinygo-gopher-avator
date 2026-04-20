//go:build tn

package avatar

// Colors for TN panels (no DisplayInversion).
// Values are sent directly to display without hardware inversion.
const eyeWhiteRGB565  = uint16(0xFFFF) // white
const pupilRGB565     = uint16(0x0000) // black
const highlightRGB565 = uint16(0xFFFF) // white
