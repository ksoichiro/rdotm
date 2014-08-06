// Copyright (c) 2014 Soichiro Kashima
// Licensed under MIT license.

package main

import (
	"encoding/hex"
	"strings"
)

func hexToInt(hexString string) (a, r, g, b int) {
	raw := hexString
	// Remove prefix '#'
	if strings.HasPrefix(raw, "#") {
		braw := []byte(raw)
		raw = string(braw[1:])
	}

	// Format hex string
	if len(raw) == 8 {
		// AARRGGBB: Do nothing
	} else if len(raw) == 6 {
		// RRGGBB: Insert alpha(FF)
		raw = "FF" + raw
	} else if len(raw) == 4 {
		// ARGB: Duplicate each hex
		braw := []byte(raw)
		sa := string(braw[0:1])
		sr := string(braw[1:2])
		sg := string(braw[2:3])
		sb := string(braw[3:4])
		raw = sa + sa + sr + sr + sg + sg + sb + sb
	} else if len(raw) == 3 {
		// RGB: Insert alpha(F) and duplicate each hex
		raw = "F" + raw
		braw := []byte(raw)
		sa := string(braw[0:1])
		sr := string(braw[1:2])
		sg := string(braw[2:3])
		sb := string(braw[3:4])
		raw = sa + sa + sr + sr + sg + sg + sb + sb
	}
	bytes, _ := hex.DecodeString(raw)
	a = int(bytes[0])
	r = int(bytes[1])
	g = int(bytes[2])
	b = int(bytes[3])
	return
}
