package test

import (
	"errors"
	"testing"

	"composite/src/window"
)

// mustWindow baut ein Fenster oder bricht den Test ab.
func mustWindow(t *testing.T, panes [4]window.Item) *window.Window {
	t.Helper()
	w, err := window.NewWindow(window.NewFrame(), panes)
	if err != nil {
		t.Fatalf("NewWindow: %v", err)
	}
	return w
}

func g() window.Item { return window.NewGlass() }

func simple(t *testing.T) *window.Window {
	return mustWindow(t, [4]window.Item{g(), g(), g(), g()})
}

func split(t *testing.T) *window.Window {
	return mustWindow(t, [4]window.Item{g(), g(), g(), simple(t)})
}

func muntin(t *testing.T) *window.Window {
	return mustWindow(t, [4]window.Item{simple(t), g(), split(t), g()})
}

// T-1
func TestLeafPrices(t *testing.T) {
	if got := window.NewGlass().Price(); got != 30 {
		t.Errorf("Glass.Price() = %d, want 30", got)
	}
	if got := window.NewFrame().Price(); got != 50 {
		t.Errorf("Frame.Price() = %d, want 50", got)
	}
}

// T-2
func TestExamplePrices(t *testing.T) {
	tests := []struct {
		name string
		root window.Item
		want int
	}{
		{"simple", simple(t), 170},
		{"split", split(t), 310},
		{"muntin", muntin(t), 590},
	}
	for _, tt := range tests {
		if got := tt.root.Price(); got != tt.want {
			t.Errorf("%s: Price() = %d, want %d", tt.name, got, tt.want)
		}
	}
}

// T-3
func TestLeafDisplay(t *testing.T) {
	if got := window.NewGlass().Display(); got != "Glass Panel\n" {
		t.Errorf("Glass.Display() = %q", got)
	}
	if got := window.NewFrame().Display(); got != "Frame\n" {
		t.Errorf("Frame.Display() = %q", got)
	}
}

// T-3
func TestWindowDisplay(t *testing.T) {
	want := "Window\n" +
		"  Frame\n" +
		"  Glass Panel\n" +
		"  Glass Panel\n" +
		"  Glass Panel\n" +
		"  Window\n" +
		"    Frame\n" +
		"    Glass Panel\n" +
		"    Glass Panel\n" +
		"    Glass Panel\n" +
		"    Glass Panel\n"
	if got := split(t).Display(); got != want {
		t.Errorf("Display() =\n%s\nwant\n%s", got, want)
	}
}

// T-4
func TestNewWindowInvalid(t *testing.T) {
	var nilGlass *window.Glass
	tests := []struct {
		name  string
		frame *window.Frame
		panes [4]window.Item
	}{
		{"nil frame", nil, [4]window.Item{g(), g(), g(), g()}},
		{"nil pane", window.NewFrame(), [4]window.Item{g(), nil, g(), g()}},
		{"typed nil pane", window.NewFrame(), [4]window.Item{g(), g(), nilGlass, g()}},
		{"frame as pane", window.NewFrame(), [4]window.Item{g(), g(), g(), window.NewFrame()}},
	}
	for _, tt := range tests {
		w, err := window.NewWindow(tt.frame, tt.panes)
		if !errors.Is(err, window.ErrInvalidWindow) {
			t.Errorf("%s: err = %v, want ErrInvalidWindow", tt.name, err)
		}
		if w != nil {
			t.Errorf("%s: window = %v, want nil", tt.name, w)
		}
	}
}

// T-5: Leaf und Composite werden über dieselbe Schnittstelle benutzt.
func TestUniformTreatment(t *testing.T) {
	items := []window.Item{window.NewGlass(), window.NewFrame(), simple(t)}
	sum := 0
	for _, it := range items {
		sum += it.Price()
	}
	if sum != 30+50+170 {
		t.Errorf("sum = %d, want 250", sum)
	}
}
