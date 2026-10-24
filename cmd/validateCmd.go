package cmd

import (
	"fmt"

	"github.com/navigator-systems/go-chaos/config"
)

func ValidateCmd(file string) error {
	cfg, err := config.ReadFile(file)
	if err != nil {
		return fmt.Errorf("I'm unable to read or unmarshal file %q: %w", file, err)
	}

	fmt.Printf("Application: %s\n", cfg.App)
	fmt.Printf("Description: %s\n", cfg.Description)
	if cfg.Function != "" {
		fmt.Printf("Function: %s\n", cfg.Function)
	}

	fmt.Printf("Jobs: %d\n", len(cfg.Job))
	for i, job := range cfg.Job {
		fmt.Printf("\nJob %d:\n", i+1)
		fmt.Printf("  Cloud: %s\n", job.Cloud)
		fmt.Printf("  Service: %s\n", job.Service)
		if job.Region != "" {
			fmt.Printf("  Region: %s\n", job.Region)
		}
		if job.Namespace != "" {
			fmt.Printf("  Namespace: %s\n", job.Namespace)
		}
		if job.Project != "" {
			fmt.Printf("  Project: %s\n", job.Project)
		}
		fmt.Printf("  Chaos: %s\n", job.Chaos.Chaos)
		fmt.Printf("  Tag: %s\n", job.Chaos.Tag)
		fmt.Printf("  Count: %d\n", job.Chaos.Count)
	}

	fmt.Printf("\nNotifications: %d\n", len(cfg.Notifications))
	return nil
}
