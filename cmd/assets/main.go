package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/itsliamegan/assets"
	"github.com/itsliamegan/assets/manifest"
)

func main() {
	var outDir string
	var manifestFile string
	var optimize bool

	flag.StringVar(&outDir, "outdir", "", "directory to place output files in")
	flag.StringVar(&manifestFile, "manifest", "", "file to write the asset manifest to")
	flag.BoolVar(&optimize, "optimize", false, "perform optimizations to minimize the size of output files")

	flag.Usage = func() {
		fmt.Println("assets - bundle and transform frontend assets, using sensible defaults")
		fmt.Println()
		fmt.Println("usage: assets [options] <entrypoint>...")
		fmt.Println()
		fmt.Println("options: ")
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("assets: no entrypoints specified")
		os.Exit(1)
	}

	if outDir == "" {
		fmt.Println("assets: output directory not specified")
		os.Exit(1)
	}

	if manifestFile == "" {
		fmt.Println("assets: manifest file not specified")
		os.Exit(1)
	}

	entrypoints := flag.Args()

	files, mfest, err := assets.Build(entrypoints, outDir, optimize)
	exitIf(err)

	err = assets.WriteAll(files)
	exitIf(err)

	err = manifest.Write(mfest, manifestFile)
	exitIf(err)
}

func exitIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
