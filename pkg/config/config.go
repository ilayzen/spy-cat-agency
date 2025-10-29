package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadFromFile(path string, config any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config, err: %v", err)
	}

	err = yaml.Unmarshal(b, config)
	if err != nil {
		return fmt.Errorf("failed to unmarchal config, err: %v", err)
	}

	return nil
}