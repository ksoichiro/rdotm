package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDrawable(t *testing.T) {
	opt := Options{ResDir: filepath.Join("testdata", "res"), OutDir: "_out"}
	res := parseDrawables(&opt)
	if len(res.Drawables) != 1 {
		t.Errorf("Expected %d strings but was %d\n", 1, len(res.Drawables))
	}
	os.RemoveAll(opt.OutDir)

	opt = Options{ResDir: filepath.Join("testdata", "res2"), OutDir: "_out"}
	res = parseDrawables(&opt)
	if len(res.Drawables) != 3 {
		t.Errorf("Expected %d strings but was %d\n", 3, len(res.Drawables))
	}
	os.RemoveAll(opt.OutDir)
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
