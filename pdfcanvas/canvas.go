package pdfcanvas

import (
	"image"
	"io"

	"github.com/hyde042/canvas"
	"github.com/hyde042/canvas/layout"

	"github.com/signintech/gopdf"
)

var _ interface {
	canvas.Canvas
	canvas.TextCanvas
} = &Canvas{}

type Canvas struct {
	doc gopdf.GoPdf
}

func New( /* TODO: options */ ) *Canvas {
	var c Canvas
	c.doc.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	c.doc.AddPage()

	// TODO: opetation logging and proper page building
	// TODO: load default font(s) lazily

	fontData, err := canvas.FontFS.ReadFile("font/LiberationSerif-Regular.ttf")
	if err != nil {
		panic(err)
	}
	if err := c.doc.AddTTFFontData("", fontData); err != nil {
		panic(err)
	}
	c.SetFont("", canvas.DefaultFontSize)
	return &c
}

func (t *Canvas) Draw(p layout.Point, img image.Image, rect *image.Rectangle) {

	// TODO

}

func (t *Canvas) SetFont(name string, size int) {
	t.doc.SetFont(name, "", size)
}

func (t *Canvas) DrawText(p layout.Point, s string) {
	t.move(p)
	t.doc.Text(s)
}

func (t *Canvas) Write(w io.Writer) error {
	return t.doc.Write(w)
}

func (t *Canvas) move(p layout.Point) {
	t.doc.SetX(p.X())
	t.doc.SetY(p.Y())
}
