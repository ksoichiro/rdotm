package main

import "testing"

var hextests = []struct {
	expr string
	a    int
	r    int
	g    int
	b    int
}{
	{"#FF336699", 255, 51, 102, 153},
	{"FF336699", 255, 51, 102, 153},
	{"#336699", 255, 51, 102, 153},
	{"#9369", 153, 51, 102, 153},
	{"#369", 255, 51, 102, 153},
}

func TestHexToInt(t *testing.T) {
	for _, tt := range hextests {
		a, r, g, b := hexToInt(tt.expr)
		if a != tt.a || r != tt.r || g != tt.g || b != tt.b {
			t.Errorf("Expected %s -> (a, r, g, b)=(%d, %d, %d, %d) but was (%d,%d,%d,%d)\n", tt.expr, tt.a, tt.r, tt.g, tt.b, a, r, g, b)
		}
	}
}
