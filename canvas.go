package canvas

import (
	"image"

	"github.com/hyde042/canvas/layout"
)

const DefaultFontSize = 16

type Canvas interface {
	Draw(p layout.Point, img image.Image, rect *image.Rectangle)
}

type TextCanvas interface {
	Canvas
	SetFont(name string, size int)
	DrawText(p layout.Point, s string)
}

type PagedCanvas interface {
	Canvas
	AddPage()
}
