// Package window implementiert das Composite-Pattern am Beispiel eines
// verschachtelbaren Fensters (siehe spec.md).
package window

// Item ist die gemeinsame Schnittstelle (Component) für Leaves und
// Composites. Der Aufrufer behandelt einen ganzen Baum wie ein einzelnes
// Item (S-1).
type Item interface {
	Price() int
	Display() string
}
