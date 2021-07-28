package canvas

import (
	"image"
)

type Canvas interface {
	Move(image.Point)
	Draw(img image.Image, rect *image.Rectangle)
}

type TextCanvas interface {
	Canvas
	SetFont(name string, size int)
	Text(s string)
}
