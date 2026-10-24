package cmd

import (
	"fmt"

	"github.com/navigator-systems/go-chaos/config"
)

func DestroyCmd(file string) error {
	cfg, err := config.ReadFile(file)
	if err != nil {
		return fmt.Errorf("I'm unable to read or unmarshal file %q: %w", file, err)
	}
	selectFunction(&cfg)

	return nil
}
