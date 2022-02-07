package esbuild

import (
	"encoding/json"
	"errors"
	"fmt"

	esbuild "github.com/evanw/esbuild/pkg/api"
)

type OutputFile = esbuild.OutputFile

type Metadata struct {
	Outputs map[string]Output `json:outputs`
}

type Output struct {
	Entrypoint string `json:entryPoint`
}

func Build(entrypoints []string, outDir string, optimize bool) ([]OutputFile, Metadata, error) {
	options := esbuild.BuildOptions{
		EntryPoints: entrypoints,
		Sourcemap:   esbuild.SourceMapInline,
		Target:      esbuild.ESNext,
		Bundle:      true,
		Write:       false,
		Metafile:    true,
		Outdir:      outDir,
	}

	if optimize {
		options.EntryNames = "[name]-[hash]"
		options.Sourcemap = esbuild.SourceMapLinked
		options.Target = esbuild.ES2015
		options.MinifyWhitespace = true
		options.MinifyIdentifiers = true
		options.MinifySyntax = true
	}

	result := esbuild.Build(options)

	if len(result.Errors) > 0 {
		return nil, Metadata{}, errors.New("compiling assets")
	}

	var metadata Metadata
	err := json.Unmarshal([]byte(result.Metafile), &metadata)
	if err != nil {
		return nil, Metadata{}, fmt.Errorf("unmarshalling esbuild metafile: %w", err)
	}

	return result.OutputFiles, metadata, nil
}
