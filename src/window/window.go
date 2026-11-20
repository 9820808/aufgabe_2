package window

import (
	"errors"
	"fmt"
	"strings"
)

const indent = "  " // S-13

// ErrInvalidWindow wird von NewWindow bei ungültigem Aufbau geliefert.
var ErrInvalidWindow = errors.New("invalid window")

// Window ist das Composite: genau ein Frame und vier Felder, die jeweils
// Glass oder Window sind (S-2, S-3).
type Window struct {
	frame *Frame
	panes [4]Item
}

// NewWindow erzeugt ein Fenster und prüft den Aufbau (S-6).
// Felder sind von 0..3 links nach rechts, oben nach unten nummeriert (S-4).
func NewWindow(frame *Frame, panes [4]Item) (*Window, error) {
	if frame == nil {
		return nil, fmt.Errorf("%w: frame is nil", ErrInvalidWindow)
	}
	for i, p := range panes {
		switch v := p.(type) {
		case *Glass:
			if v == nil {
				return nil, fmt.Errorf("%w: pane %d is nil", ErrInvalidWindow, i)
			}
		case *Window:
			if v == nil {
				return nil, fmt.Errorf("%w: pane %d is nil", ErrInvalidWindow, i)
			}
		case nil:
			return nil, fmt.Errorf("%w: pane %d is nil", ErrInvalidWindow, i)
		default:
			return nil, fmt.Errorf("%w: pane %d must be Glass or Window, got %T", ErrInvalidWindow, i, p)
		}
	}
	return &Window{frame: frame, panes: panes}, nil
}

// Price summiert Rahmen und Felder rekursiv; Window hat keinen
// Eigenpreis (S-10).
func (w *Window) Price() int {
	sum := w.frame.Price()
	for _, p := range w.panes {
		sum += p.Price()
	}
	return sum
}

// Display gibt die eigene Zeile aus, danach Frame und Felder 0..3 in
// Pre-Order, jeweils um eine Ebene eingerückt (S-12, S-13).
func (w *Window) Display() string {
	var b strings.Builder
	b.WriteString("Window\n")
	writeIndented(&b, w.frame.Display())
	for _, p := range w.panes {
		writeIndented(&b, p.Display())
	}
	return b.String()
}

// writeIndented rückt jede Zeile von s um eine Ebene ein.
func writeIndented(b *strings.Builder, s string) {
	for _, line := range strings.SplitAfter(s, "\n") {
		if line == "" {
			continue
		}
		b.WriteString(indent)
		b.WriteString(line)
	}
}
