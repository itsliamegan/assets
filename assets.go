package assets

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/itsliamegan/assets/esbuild"
	"github.com/itsliamegan/assets/manifest"
)

type OutputFile = esbuild.OutputFile

func Build(entrypoints []string, outDir string, optimize bool) ([]OutputFile, manifest.Manifest, error) {
	files, metadata, err := esbuild.Build(entrypoints, outDir, optimize)
	if err != nil {
		return nil, nil, err
	}

	mfest := make(manifest.Manifest)
	for outputPath, output := range metadata.Outputs {
		if output.Entrypoint == "" {
			continue
		}

		entrypointBasename := filepath.Base(output.Entrypoint)
		outputFileBasename := filepath.Base(outputPath)

		mfest[entrypointBasename] = outputFileBasename
	}

	return files, mfest, nil
}

func WriteAll(files []OutputFile) error {
	for _, file := range files {
		parent := filepath.Dir(file.Path)
		err := os.MkdirAll(parent, 0744)
		if err != nil {
			return fmt.Errorf("writing output file: %w", err)
		}

		err = os.WriteFile(file.Path, file.Contents, 0644)
		if err != nil {
			return fmt.Errorf("writing output file: %w", err)
		}
	}

	return nil
}
