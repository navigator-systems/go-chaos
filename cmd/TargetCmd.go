package cmd

import (
	"errors"
	"log"
	"strings"

	"github.com/navigator-systems/go-chaos/config"
)

func TargetCmd(file string, target string) (*config.GenConfig, error) {

	targetVars := strings.Split(target, ".")
	if len(targetVars) < 3 {
		err := errors.New("target must contain 3 values, separated by dot, example, aws.ec2.app:prod")
		log.Println("error: ", err)
		return nil, err
	}
	cloud, service, tag := targetVars[0], targetVars[1], targetVars[2]

	cfg, err := config.ReadFile(file)

	if err != nil {
		log.Println("failed to load configuration: ", err)
		return nil, err
	}

	for j := 0; j < len(cfg.Job); j++ {
		if cfg.Job[j].Cloud == cloud && cfg.Job[j].Service == service && cfg.Job[j].Chaos.Tag == tag {
			switchService(cfg.Job[j], false)
			break
		}
	}
	return &cfg, nil
}
