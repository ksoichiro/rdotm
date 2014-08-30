// Copyright (c) 2014 Soichiro Kashima
// Licensed under MIT license.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	OutputHeader = `// DO NOT EDIT.
// This file is automatically generated by rdotm tool.
// https://github.com/ksoichiro/rdotm

`
)

func printLocalizableStrings(res *Resources, opt *Options, lang string) {
	class := opt.Class

	// Create language separated directory
	langDir := filepath.Join(opt.OutDir, lang+".lproj")
	os.MkdirAll(langDir, 0777)

	filename := filepath.Join(langDir, class+".strings")
	f, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	defer f.Close()

	f.WriteString(OutputHeader)

	// String
	for _, s := range res.Strings {
		f.WriteString(fmt.Sprintf("\"%s\" = \"%s\";\n", s.Name, s.Value))
	}

	f.Close()
}

func printAsObjectiveC(res *Resources, opt *Options) {
	class := opt.Class

	// Print header file(.h)
	filename := filepath.Join(opt.OutDir, class+".h")
	f, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	defer f.Close()

	f.WriteString(OutputHeader)
	f.WriteString(fmt.Sprintf(`#import <UIKit/UIKit.h>

@interface %s : NSObject

`, class))

	// String
	for _, s := range res.Strings {
		// Method definition
		f.WriteString(fmt.Sprintf(`/** %s */
+ (NSString *)%s%s;
`, s.Value, opt.PrefixStrings, s.Name))
	}

	// Integer
	for _, i := range res.Integers {
		// Method definition
		f.WriteString(fmt.Sprintf(`/** %s */
+ (NSInteger)%s%s;
`, i.Value, opt.PrefixIntegers, i.Name))
	}

	// Color
	for _, c := range res.Colors {
		// Method definition
		f.WriteString(fmt.Sprintf(`/** %s */
+ (UIColor *)%s%s;
`, c.Value, opt.PrefixColors, c.Name))
	}

	// Drawable
	for _, d := range res.Drawables {
		// Method definition
		f.WriteString(fmt.Sprintf(`+ (UIImage *)%s%s;
`, opt.PrefixDrawables, d.Name))
	}

	// Integer array
	for _, i := range res.IntegerArrays {
		// Method definition
		f.WriteString(fmt.Sprintf(`+ (NSArray *)%s%s;
`, opt.PrefixIntegerArrays, i.Name))
	}

	f.WriteString(`
@end
`)
	f.Close()

	// Print implementation file(.m)
	filename = filepath.Join(opt.OutDir, class+".m")
	f, _ = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	defer f.Close()

	// Import header file
	f.WriteString(OutputHeader)
	f.WriteString(fmt.Sprintf(`#import "%s.h"

@implementation %s

`, class, class))

	// String
	for _, s := range res.Strings {
		// Method implementation
		if opt.Localize {
			// Read from LANG.lproj/R.strings
			f.WriteString(fmt.Sprintf("+ (NSString *)%s%s { return NSLocalizedStringFromTable(@\"%s\", @\"%s\", nil); }\n", opt.PrefixStrings, s.Name, s.Name, class))
		} else {
			f.WriteString(fmt.Sprintf("+ (NSString *)%s%s { return @\"%s\"; }\n", opt.PrefixStrings, s.Name, s.Value))
		}
	}

	// Integer
	for _, i := range res.Integers {
		// Method implementation
		f.WriteString(fmt.Sprintf("+ (NSInteger)%s%s { return %s; }\n", opt.PrefixIntegers, i.Name, i.Value))
	}

	// Color
	for _, c := range res.Colors {
		// Method implementation
		a, r, g, b := hexToInt(c.Value)
		f.WriteString(fmt.Sprintf("+ (UIColor *)%s%s { return [UIColor colorWithRed:%d/255.0 green:%d/255.0 blue:%d/255.0 alpha:%d/255.0]; }\n", opt.PrefixColors, c.Name, r, g, b, a))
	}

	// Drawable
	for _, d := range res.Drawables {
		// Method implementation
		f.WriteString(fmt.Sprintf("+ (UIImage *)%s%s { return [UIImage imageNamed:@\"%s\"]; }\n", opt.PrefixDrawables, d.Name, d.Name))
	}

	// Integer array
	for _, i := range res.IntegerArrays {
		var v = ""
		for _, n := range i.Items {
			if v != "" {
				v += ", "
			}
			v += "@" + n.Value
		}
		// Method implementation
		f.WriteString(fmt.Sprintf("+ (NSArray *)%s%s { return @[%s]; }\n", opt.PrefixIntegerArrays, i.Name, v))
	}

	f.WriteString(`
@end
`)
	f.Close()
}
