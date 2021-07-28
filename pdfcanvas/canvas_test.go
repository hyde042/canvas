package pdfcanvas_test

import (
	"image"
	"os"
	"testing"

	"github.com/hyde042/canvas/pdfcanvas"
)

func TestCanvas(t *testing.T) {
	c := pdfcanvas.New()

	c.SetFont("", 24)
	for i := 1; i <= 100; i++ {
		c.Move(image.Point{30, i * 30})
		c.Text("Hello, world!")
	}
	if err := os.MkdirAll("temp", 0777); err != nil {
		t.Fatal(err)
	}
	f, err := os.Create("temp/test.pdf")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := c.Write(f); err != nil {
		t.Fatal(err)
	}
}
