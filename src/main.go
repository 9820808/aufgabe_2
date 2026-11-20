// Demonstrator für das Composite-Pattern: baut drei Fensterbäume und gibt
// jeweils Display() und Price() aus (S-15, S-16).
package main

import (
	"fmt"
	"os"

	"composite/src/window"
)

// example ist ein benannter Beispielbaum.
type example struct {
	name string
	root window.Item
}

func main() {
	examples, err := buildExamples()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for i, ex := range examples {
		fmt.Printf("=== Beispiel %d: %s ===\n", i+1, ex.name)
		fmt.Print(ex.root.Display())
		fmt.Printf("Preis: %d\n\n", ex.root.Price())
	}
}

// buildExamples erzeugt die drei Beispielbäume aus spec.md, Abschnitt 7.
func buildExamples() ([]example, error) {
	simple, err := simpleWindow()
	if err != nil {
		return nil, fmt.Errorf("simple window: %w", err)
	}

	split, err := splitWindow()
	if err != nil {
		return nil, fmt.Errorf("split window: %w", err)
	}

	sub0, err := simpleWindow()
	if err != nil {
		return nil, fmt.Errorf("muntin window pane 0: %w", err)
	}
	sub2, err := splitWindow()
	if err != nil {
		return nil, fmt.Errorf("muntin window pane 2: %w", err)
	}
	muntin, err := window.NewWindow(window.NewFrame(), [4]window.Item{
		sub0, window.NewGlass(), sub2, window.NewGlass(),
	})
	if err != nil {
		return nil, fmt.Errorf("muntin window: %w", err)
	}

	return []example{
		{"Einfaches Fenster", simple},
		{"Geteiltes Feld", split},
		{"Sprossenfenster", muntin},
	}, nil
}

// simpleWindow: Fenster mit vier Glasfeldern.
func simpleWindow() (*window.Window, error) {
	return window.NewWindow(window.NewFrame(), [4]window.Item{
		window.NewGlass(), window.NewGlass(), window.NewGlass(), window.NewGlass(),
	})
}

// splitWindow: Felder 0..2 Glass, Feld 3 ein Unterfenster (Beispiel aus der Aufgabe).
func splitWindow() (*window.Window, error) {
	sub, err := simpleWindow()
	if err != nil {
		return nil, fmt.Errorf("pane 3: %w", err)
	}
	return window.NewWindow(window.NewFrame(), [4]window.Item{
		window.NewGlass(), window.NewGlass(), window.NewGlass(), sub,
	})
}
