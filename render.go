package canvas

import "github.com/hyde042/canvas/layout"

func Render(c Canvas, l layout.Layout) {
	for _, v := range l {
		switch v := v.(type) {
		case layout.Font:
			setFont(c, v)
		case layout.Text:
			drawText(c, v)
		case layout.Page:
			addPage(c)
		}

		// TODO: image
	}
}

// TODO: renderer type with state and image based fallbacks for everything

func setFont(c Canvas, f layout.Font) {
	if tc, ok := c.(TextCanvas); ok {
		tc.SetFont(f.Family, f.Size)
	}
}

func drawText(c Canvas, t layout.Text) {
	if tc, ok := c.(TextCanvas); ok {
		tc.DrawText(t.Position, t.Text)
	}
}

func addPage(c Canvas) {
	if pc, ok := c.(PagedCanvas); ok {
		pc.AddPage()
	}
}
