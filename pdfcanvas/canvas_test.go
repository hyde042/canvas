package pdfcanvas_test

import (
	"os"
	"testing"

	"github.com/hyde042/canvas/layout"
	"github.com/hyde042/canvas/pdfcanvas"
)

func TestCanvas(t *testing.T) {
	c := pdfcanvas.New()

	c.SetFont("", 24)
	for i := 1; i <= 100; i++ {
		c.DrawText(layout.Point{30, float64(i * 30)}, "Hello, world!")
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
