//go:build windows

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Microsoft/hcsshim/osversion"
)

func usage() {
	exe := filepath.Base(os.Args[0])
	fmt.Printf("Usage: %s [-b]\n", exe)
	flag.PrintDefaults()
	txt := `
Prints the Windows OS version. Intended for Windows container troubleshooting. (Uses the same code path as Docker/Moby to get the host's os version for matching with os.version in container image manifests).

Examples:
$ %s
10.0.22000
$ %s -b
22000
	`

	fmt.Printf(txt, exe, exe)
}

func main() {
	flag.Usage = usage

	FlagShort := flag.Bool("b", false, "print only the build version (i.e. '2200' instead of '10.0.22000')")
	flag.Parse()
	UseShort := *FlagShort

	// the windows .exe needs to have a manifest, otherwise the output of osversion.Get() is incorrect!
	version := osversion.Get()

	switch UseShort {
	case true:
		fmt.Println(version.Build)
	default:
		fmt.Printf("%d.%d.%d", version.MajorVersion, version.MinorVersion, version.Build)
	}

}
