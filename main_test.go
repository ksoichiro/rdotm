package main

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

const testOutDir = "_out"

func TestMain(t *testing.T) {
	defer os.RemoveAll(testOutDir)
	os.Args = []string{os.Args[0], "-res=" + filepath.Join("testdata", "res"), "-out=" + testOutDir}
	os.RemoveAll(testOutDir)
	main()
	if !mainExecuted() {
		t.Errorf("Expected main process executed but not executed\n")
	}
}

func TestMainLocalizeClean(t *testing.T) {
	defer os.RemoveAll(testOutDir)
	flag.CommandLine = flag.NewFlagSet("LocalizeClean", flag.PanicOnError)
	os.Args = []string{os.Args[0], "-res=" + filepath.Join("testdata", "res2"), "-out=" + testOutDir, "-localize", "-clean"}
	os.RemoveAll(testOutDir)
	main()
	if !mainExecuted() {
		t.Errorf("Expected main process executed but not executed\n")
	}
}

func TestMainLackOfRequiredOptionsRes(t *testing.T) {
	defer os.RemoveAll(testOutDir)
	flag.CommandLine = flag.NewFlagSet("LackOfRequiredOptionsRes", flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-out=" + testOutDir}
	os.RemoveAll(testOutDir)
	main()
	if mainExecuted() {
		t.Errorf("Expected main process not executed but executed\n")
	}
}

func TestMainLackOfRequiredOptionsOut(t *testing.T) {
	defer os.RemoveAll(testOutDir)
	flag.CommandLine = flag.NewFlagSet("LackOfRequiredOptionsOut", flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-res=" + filepath.Join("testdata", "res")}
	os.RemoveAll(testOutDir)
	main()
	if mainExecuted() {
		t.Errorf("Expected main process not executed but executed\n")
	}
}

func TestMainInvalidType(t *testing.T) {
	defer os.RemoveAll(testOutDir)
	flag.CommandLine = flag.NewFlagSet("InvalidType", flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-res=" + filepath.Join("testdata", "res"), "-out=" + testOutDir, "-types=dimen"}
	os.RemoveAll(testOutDir)
	main()
	if mainExecuted() {
		t.Errorf("Expected main process not executed but executed\n")
	}
}

func mainExecuted() bool {
	if _, err := os.Stat(testOutDir); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
