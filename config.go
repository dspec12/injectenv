package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const exampleYAML = `---
profile1:
    key1: value
    key2: value2

`

func (env *envConfig) loadConfigFile() {
	homedir := os.Getenv("HOME")
	if homedir == "" {
		log.Fatalf("Environmental variable for $HOME is not set.")
	}

	configFile := homedir + "/.injectenv.yaml"
	if !fileExists(configFile) {
		log.Println("No config file detected. An example will be created at:", configFile)
		err := os.WriteFile(configFile, []byte(exampleYAML), 0660)
		if err != nil {
			log.Fatal(err)
		}
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(data), &env); err != nil {
		log.Fatal(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
