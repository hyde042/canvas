package canvas

import (
	"image"
)

const DefaultFontSize = 16

type Canvas interface {
	Move(p image.Point)
	Draw(img image.Image, rect *image.Rectangle)
}

type TextCanvas interface {
	Canvas
	SetFont(name string, size int)
	Text(s string)
}
