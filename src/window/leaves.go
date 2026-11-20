package window

import "fmt"

const (
	glassPrice = 30 // S-8
	framePrice = 50 // S-9
)

// Glass ist eine Glasscheibe (Leaf).
type Glass struct{}

// NewGlass erzeugt eine Glasscheibe.
func NewGlass() *Glass {
	return &Glass{}
}

// Price liefert den Preis der Scheibe (S-8).
func (g *Glass) Price() int {
	return glassPrice
}

// Display liefert die eigene Zeile mit Preis ohne Einrückung (S-11, S-14).
func (g *Glass) Display() string {
	return fmt.Sprintf("Glass Panel (%d)\n", g.Price())
}

// Frame ist ein Fensterrahmen (Leaf).
type Frame struct{}

// NewFrame erzeugt einen Rahmen.
func NewFrame() *Frame {
	return &Frame{}
}

// Price liefert den Preis des Rahmens (S-9).
func (f *Frame) Price() int {
	return framePrice
}

// Display liefert die eigene Zeile mit Preis ohne Einrückung (S-11, S-14).
func (f *Frame) Display() string {
	return fmt.Sprintf("Frame (%d)\n", f.Price())
}
