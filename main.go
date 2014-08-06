// Copyright (c) 2014 Soichiro Kashima
// Licensed under MIT license.

package main

import (
	"flag"
	"os"
)

// Command line options
type Options struct {
	ResDir   string
	OutDir   string
	Class    string
	Clean    bool
	Localize bool
}

// Resource model structure
type Resources struct {
	Language  string     `xml:"-"`
	Strings   []String   `xml:"string"`
	Colors    []Color    `xml:"color"`
	Drawables []Drawable `xml:"-"`
}

type String struct {
	Name  string `xml:"name,attr"`
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
	)
	flag.Parse()
	if *resDir == "" || *outDir == "" {
		// Exit if the required options are empty
		flag.Usage()
		os.Exit(1)
	}

	// Parse resource XML files and generate source code
	parse(&Options{
		ResDir:   *resDir,
		OutDir:   *outDir,
		Class:    *class,
		Clean:    *clean,
		Localize: *localize})
}
