// Copyright (c) 2014 Soichiro Kashima
// Licensed under MIT license.

package main

import "flag"

// Command line options
type Options struct {
	ResDir          string
	OutDir          string
	Class           string
	Clean           bool
	Localize        bool
	PrefixStrings   string
	PrefixIntegers  string
	PrefixColors    string
	PrefixDrawables string
}

// Resource model structure
type Resources struct {
	Language  string     `xml:"-"`
	Strings   []String   `xml:"string"`
	Integers  []Integer  `xml:"integer"`
	Colors    []Color    `xml:"color"`
	Drawables []Drawable `xml:"-"`
	Items     []Item     `xml:"item"`
}

type String struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Integer struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Item struct {
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type Color struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Drawable struct {
	Name string
}

func main() {
	// Get command line options
	var (
		resDir   = flag.String("res", "", "Resource(res) directory path. Required.")
		outDir   = flag.String("out", "", "Output directory path. Required.")
		class    = flag.String("class", "R", "Class name to overwrite default value(R). Optional.")
		clean    = flag.Bool("clean", false, "Clean output directory before execution.")
		localize = flag.Bool("localize", false, "Enable localization using NSLocalizedStringFromTable.")
		ps       = flag.String("ps", "string_", "Prefix for generated string methods.")
		pi       = flag.String("pi", "integer_", "Prefix for generated integer methods.")
		pc       = flag.String("pc", "color_", "Prefix for generated color methods.")
		pd       = flag.String("pd", "drawable_", "Prefix for generated drawable methods.")
	)
	flag.Parse()
	if *resDir == "" || *outDir == "" {
		// Exit if the required options are empty
		flag.Usage()
		return
	}

	// Parse resource XML files and generate source code
	parse(&Options{
		ResDir:          *resDir,
		OutDir:          *outDir,
		Class:           *class,
		Clean:           *clean,
		Localize:        *localize,
		PrefixStrings:   *ps,
		PrefixIntegers:  *pi,
		PrefixColors:    *pc,
		PrefixDrawables: *pd})
}
