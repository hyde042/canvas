package layout

import (
	"fmt"
	"strconv"
	"strings"
)

type Point []float64

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
