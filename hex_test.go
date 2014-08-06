package main

import (
	"fmt"
	"testing"
)

func TestHexToInt(t *testing.T) {
	if fail, msg := testHexToInt(t, "#FF336699", 255, 51, 102, 153); fail {
		t.Errorf(msg)
	}
	if fail, msg := testHexToInt(t, "FF336699", 255, 51, 102, 153); fail {
		t.Errorf(msg)
	}
	if fail, msg := testHexToInt(t, "#336699", 255, 51, 102, 153); fail {
		t.Errorf(msg)
	}
	if fail, msg := testHexToInt(t, "#9369", 153, 51, 102, 153); fail {
		t.Errorf(msg)
	}
	if fail, msg := testHexToInt(t, "#369", 255, 51, 102, 153); fail {
		t.Errorf(msg)
	}
}

func testHexToInt(t *testing.T, expr string, ea, er, eg, eb int) (fail bool, msg string) {
	fail = false
	a, r, g, b := hexToInt(expr)
	if a != ea || r != er || g != eg || b != eb {
		msg = fmt.Sprintf("Expected %s -> (a, r, g, b)=(%d, %d, %d, %d) but was (%d,%d,%d,%d)", expr, ea, er, eg, eb, a, r, g, b)
		fail = true
	}
	return
}
