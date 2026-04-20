//go:build !tn

package avatar

// Colors pre-inverted to compensate for DisplayInversion (INVON).
// INVON: sent value is hardware-inverted before display.
// 0x0000 → displays as white, 0xFFFF → displays as black.
const eyeWhiteRGB565  = uint16(0x0000)
const pupilRGB565     = uint16(0xFFFF)
const highlightRGB565 = uint16(0x0000)
