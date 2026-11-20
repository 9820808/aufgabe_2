package window

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

// Display liefert die eigene Zeile ohne Einrückung (S-11, S-14).
func (g *Glass) Display() string {
	return "Glass Panel\n"
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

// Display liefert die eigene Zeile ohne Einrückung (S-11, S-14).
func (f *Frame) Display() string {
	return "Frame\n"
}
