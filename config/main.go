package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type EnvConfig map[string]map[string]string

var EnvMap EnvConfig

func (env *EnvConfig) LoadConfigFile(filePath string) error {
	if !fileExists(filePath) {
		return errors.New("no config file found")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal([]byte(data), &env); err != nil {
		return err
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
