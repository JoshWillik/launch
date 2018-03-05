package config

import (
	"encoding/json"
	"os"
	"path"
)

// Config is the deploy configuration struct
type Config struct {
	ContextRoot string
}

// Read the settings for the current project or return defaults
func Read(dir string) Config {
	config := Config{ContextRoot: dir}
	configFile, err := os.Open(path.Join(dir, "/.launch-config.json"))
	if err != nil {
		return config
	}
	json.NewDecoder(configFile).Decode(&config) // ignore error
	return config
}
