package main

import (
	"os"
	"path/filepath"
	"testing"
)

var allTypes = map[string]bool{
	"string":   true,
	"integer":  true,
	"color":    true,
	"drawable": true,
}

func TestParseDrawable(t *testing.T) {
	var tests = []struct {
		opt               Options
		expectedResources int
	}{
		{Options{ResDir: filepath.Join("testdata", "res"), OutDir: "_out", Types: allTypes}, 1},
		{Options{ResDir: filepath.Join("testdata", "res2"), OutDir: "_out", Types: allTypes}, 3},
	}
	for _, tt := range tests {
		res := parseDrawables(&tt.opt)
		if len(res.Drawables) != tt.expectedResources {
			t.Errorf("Expected %d drawables but was %d\n", tt.expectedResources, len(res.Drawables))
		}
		os.RemoveAll(tt.opt.OutDir)
	}
	for _, sw := range []bool{true, false} {
		opt := &Options{ResDir: filepath.Join("testdata", "res"), OutDir: "_out", Types: map[string]bool{"drawable": sw}}
		res := parseDrawables(opt)
		if sw && len(res.Drawables) == 0 {
			t.Errorf("Expected some drawables but was nothingd\n")
		} else if !sw && len(res.Drawables) != 0 {
			t.Errorf("Expected no drawables but was %d drawables\n", len(res.Drawables))
		}
		os.RemoveAll(opt.OutDir)
	}
}

func TestParseXml(t *testing.T) {
	res := parseXml(filepath.Join("testdata", "res", "values", "strings.xml"))
	if len(res.Strings) != 2 {
		t.Errorf("Expected %d strings but was %d\n", 2, len(res.Strings))
	}
	if res.Strings[0].Name != "title_main" {
		t.Errorf("Expected name '%s' but was '%s'\n", "title_main", res.Strings[0].Name)
	}
	if res.Strings[0].Value != "Main" {
		t.Errorf("Expected name '%s' but was '%s'\n", "Main", res.Strings[0].Value)
	}
	if res.Strings[1].Name != "label_next" {
		t.Errorf("Expected name '%s' but was '%s'\n", "label_next", res.Strings[1].Name)
	}
	if res.Strings[1].Value != "Next" {
		t.Errorf("Expected name '%s' but was '%s'\n", "Next", res.Strings[1].Value)
	}
	res = parseXml("invalid")
	if len(res.Strings) != 0 {
		t.Errorf("Expected %d strings but was %d\n", 0, len(res.Strings))
	}
}

func TestParseLang(t *testing.T) {
	res := parseLang(&Options{Types: allTypes}, filepath.Join("testdata", "res", "values"))
	if len(res.Integers) != 2 {
		t.Errorf("Expected %d strings but was %d\n", 2, len(res.Integers))
	}
	// <integer> is prioritized
	if res.Integers[0].Name != "min_length_age" {
		t.Errorf("Expected name '%s' but was '%s'\n", "min_length_age", res.Integers[0].Name)
	}
	if res.Integers[0].Value != "1" {
		t.Errorf("Expected name '%s' but was '%s'\n", "1", res.Integers[0].Value)
	}
	if res.Integers[1].Name != "max_length_name" {
		t.Errorf("Expected name '%s' but was '%s'\n", "max_length_name", res.Integers[1].Name)
	}
	if res.Integers[1].Value != "20" {
		t.Errorf("Expected name '%s' but was '%s'\n", "20", res.Integers[1].Value)
	}
	res = parseXml("invalid")
	if len(res.Integers) != 0 {
		t.Errorf("Expected %d strings but was %d\n", 0, len(res.Integers))
	}
}

func TestParseLangPartial(t *testing.T) {
	var allTypesTests = []map[string]bool{
		{"string": true, "integer": false, "color": false, "drawable": false},
		{"string": false, "integer": true, "color": false, "drawable": false},
		{"string": false, "integer": false, "color": true, "drawable": false},
	}

	for _, tests := range allTypesTests {
		res := parseLang(&Options{Types: tests}, filepath.Join("testdata", "res", "values"))
		if tests["string"] {
			if len(res.Strings) == 0 {
				t.Errorf("Expected some strings but was nothing\n")
			}
		} else {
			if len(res.Strings) != 0 {
				t.Errorf("Expected no strings but was %d\n", len(res.Strings))
			}
		}
		if tests["integer"] {
			if len(res.Integers) == 0 {
				t.Errorf("Expected some integers but was nothing\n")
			}
		} else {
			if len(res.Integers) != 0 {
				t.Errorf("Expected no integers but was %d\n", len(res.Integers))
			}
		}
		if tests["color"] {
			if len(res.Colors) == 0 {
				t.Errorf("Expected some colors but was nothing\n")
			}
		} else {
			if len(res.Colors) != 0 {
				t.Errorf("Expected no colors but was %d\n", len(res.Colors))
			}
		}
	}
}
