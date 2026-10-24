package config

import (
	"os"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

func ReadFile(filename string) (GenConfig, error) {
	// Check if file exists
	if _, err := os.Stat(filename); err != nil {
		return GenConfig{}, ErrFileNotFound
	}
	// initiate new parser
	parser := hclparse.NewParser()
	// parse file
	file, diags := parser.ParseHCLFile(filename)

	if diags.HasErrors() {
		return GenConfig{}, NewError(diags.Error(), ErrFileCantRead)
	}

	var chaosCfg GenConfig

	diags = gohcl.DecodeBody(file.Body, nil, &chaosCfg)

	if diags.HasErrors() {
		return GenConfig{}, NewError(diags.Error(), ErrFileCantRead)
	}
	return chaosCfg, nil
}
