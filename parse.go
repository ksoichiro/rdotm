// Copyright (c) 2014 Soichiro Kashima
// Licensed under MIT license.

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func parse(opt *Options) {
	// Clean and create output directory if needed
	if opt.Clean {
		// Discard all files in the output directory
		os.RemoveAll(opt.OutDir)
	}
	os.MkdirAll(opt.OutDir, 0777)

	// Parse all of the files in res/values/*.xml
	var res Resources
	if opt.Localize {
		resSubDirs, _ := ioutil.ReadDir(opt.ResDir)
		for _, valuesDir := range resSubDirs {
			// Get only values directories
			if matched, _ := regexp.MatchString("^values", valuesDir.Name()); !matched {
				continue
			}

			var lang string
			var r Resources
			if valuesDir.Name() == "values" {
				// Base language
				lang = "Base"
				r = parseLang(filepath.Join(opt.ResDir, valuesDir.Name()))
				// Output only base language to Objective-C source
				res = r
			} else {
				re := regexp.MustCompile("values-([a-zA-Z]+)")
				groups := re.FindStringSubmatch(valuesDir.Name())
				if groups == nil {
					// Not supported
					continue
				} else {
					// Maybe supported language
					lang = groups[1]
				}

				r = parseLang(filepath.Join(opt.ResDir, valuesDir.Name()))
			}
			// Create R.strings
			printLocalizableStrings(&r, opt, lang)
		}
	} else {
		valuesDir := filepath.Join(opt.ResDir, "values")
		res = parseLang(valuesDir)
	}
	resD := parseDrawables(opt)
	if 0 < len(resD.Drawables) {
		res.Drawables = append(res.Drawables, resD.Drawables...)
	}
	printAsObjectiveC(&res, opt)
}

func parseLang(valuesDir string) (res Resources) {
	files, _ := ioutil.ReadDir(valuesDir)
	// Regular expressions for format replacement
	expStr := regexp.MustCompile("%[0-9]\\$s")
	expOther := regexp.MustCompile("%[0-9]\\$([a-z])")
	for _, entry := range files {
		if matched, _ := regexp.MatchString(".xml$", entry.Name()); !matched {
			continue
		}
		entryPath := filepath.Join(valuesDir, entry.Name())
		r := parseXml(entryPath)
		if 0 < len(r.Strings) {
			// Replacing Android format to that of Objective-C
			for _, s := range r.Strings {
				// Usually, we use NSString so %1$s should be converted to '%@'
				s.Value = expStr.ReplaceAllLiteralString(s.Value, "%@")
				// Otherwise, pattern chars can be also used for Objective-C
				s.Value = expOther.ReplaceAllString(s.Value, "%$1")
				res.Strings = append(res.Strings, s)
			}
		}
		if 0 < len(r.Colors) {
			res.Colors = append(res.Colors, r.Colors...)
		}
	}
	return res
}

func parseDrawables(opt *Options) (res Resources) {
	resSubDirs, _ := ioutil.ReadDir(opt.ResDir)
	drawables := make(map[string]string)

	for _, drawableDir := range resSubDirs {
		// Get only drawable directory
		if matched, _ := regexp.MatchString("^drawable", drawableDir.Name()); !matched {
			continue
		}

		files, _ := ioutil.ReadDir(filepath.Join(opt.ResDir, drawableDir.Name()))
		for _, entry := range files {
			if matched, _ := regexp.MatchString(".(png|jpeg|jpg)$", strings.ToLower(entry.Name())); !matched {
				continue
			}
			// Identify drawables without modifiers(@*)
			basename := regexp.MustCompile("(@[^@]+)?\\.[a-zA-Z]+$").ReplaceAllString(entry.Name(), "")

			if _, ok := drawables[basename]; ok {
				// Already found
				continue
			}
			drawables[basename] = basename

			// Append new drawable name
			res.Drawables = append(res.Drawables, Drawable{Name: basename})
		}
	}
	return res
}

func parseXml(filename string) (res Resources) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(b, &res)
	if err != nil {
		fmt.Println("Error unmarshaling XML file", err)
		return
	}

	return res
}
