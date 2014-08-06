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
		for i := range resSubDirs {
			// Get only values directories
			valuesDir := resSubDirs[i]
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
	for j := range files {
		entry := files[j]
		if matched, _ := regexp.MatchString(".xml$", entry.Name()); !matched {
			continue
		}
		entryPath := filepath.Join(valuesDir, entry.Name())
		r := parseXml(entryPath)
		if 0 < len(r.Strings) {
			res.Strings = append(res.Strings, r.Strings...)
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

	for i := range resSubDirs {
		// Get only drawable directory
		drawableDir := resSubDirs[i]
		if matched, _ := regexp.MatchString("^drawable", drawableDir.Name()); !matched {
			continue
		}

		files, _ := ioutil.ReadDir(filepath.Join(opt.ResDir, drawableDir.Name()))
		for j := range files {
			entry := files[j]
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
