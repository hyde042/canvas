package pdfcanvas

import (
	"image"

	"github.com/hyde042/canvas"
)

var _ interface {
	canvas.Canvas
	canvas.TextCanvas
} = &Canvas{}

type Canvas struct {

	// ???

}

func New( /* TODO: page size argument */ ) *Canvas {
	return &Canvas{

		// ???

	}
}

func (c *Canvas) Move(image.Point) {
	panic("TODO")
}

func (c *Canvas) Draw(img image.Image, rect *image.Rectangle) {
	panic("TODO")
}

func (c *Canvas) SetFont(name string, size int) {
	panic("TODO")
}

func (c *Canvas) Text(s string) {
	panic("TODO")
}
