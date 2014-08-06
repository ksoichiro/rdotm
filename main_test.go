package main

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{os.Args[0], "-res=" + filepath.Join("testdata", "res"), "-out=_out"}
	main()
	os.RemoveAll("_out")
}

func TestMainLocalizeClean(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("LocalizeClean", flag.PanicOnError)
	os.Args = []string{os.Args[0], "-res=" + filepath.Join("testdata", "res2"), "-out=_out", "-localize", "-clean"}
	main()
	os.RemoveAll("_out")
}

func TestMainLackOfRequiredOptionsRes(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("LackOfRequiredOptionsRes", flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-out=_out"}
	main()
	os.RemoveAll("_out")
}

func TestMainLackOfRequiredOptionsOut(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("LackOfRequiredOptionsOut", flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-res=" + filepath.Join("testdata", "res")}
	main()
	os.RemoveAll("_out")
}
