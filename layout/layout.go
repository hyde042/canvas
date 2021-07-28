package layout

import (
	"fmt"
	"strconv"
	"strings"
)

type Layout []interface{}

func (t Layout) Image(p Point, path string) Layout {
	return append(t, Image{el(p), path})
}

func (t Layout) Font(family string, size int) Layout {
	return append(t, Font{family, size})
}

func (t Layout) Text(p Point, s string) Layout {
	return append(t, Text{el(p), s})
}

func el(p Point) Element {
	return Element{Position: p}
}

func (t Layout) ResolvePages(pageSize int) Layout {

	// ???

	panic("TODO")
}

type Point []float64

func (t Point) X() float64 { return t.At(0) }
func (t Point) Y() float64 { return t.At(1) }
func (t Point) Z() float64 { return t.At(2) }

func (t Point) At(i int) float64 {
	if len(t) > i {
		return t[i]
	}
	return 0.0
}

func (t Point) String() string {
	var sb strings.Builder
	for i, n := range t {
		if i > 0 {
			sb.WriteRune('/')
		}
		sb.WriteString(strconv.FormatFloat(n, 'f', 6, 64)) // TODO: trim trailing zeros?
	}
	return sb.String()
}

type Element struct {
	Position Point
}

func (t Element) String() string {
	return t.Position.String()
}

type Image struct {
	Element
	Path string
}

func (t Image) String() string {
	return fmt.Sprintf("I %v %s", t.Element, t.Path)
}

type Font struct {
	Family string
	Size   int
}

func (t Font) String() string {
	if t.Family == "" {
		return fmt.Sprintf("F %d", t.Size)
	}
	return fmt.Sprintf("F %s %d", t.Family, t.Size)
}

type Text struct {
	Element
	Text string
}

func (t Text) String() string {
	return fmt.Sprintf("I %v %s", t.Element, t.Text)
}

type Page struct{}

func (t Page) String() string {
	return "P"
}

// TODO: parsing
