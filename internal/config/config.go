package config

import (
	"encoding/json"
	"os"
	"fmt"
)

// GetEnvironment returns the environment with the given name.
//
// If the environment does not exist, an error is returned.
func (c *Config) GetEnvironment(name string) (*Environment, error) {
	env, ok := c.Environments[name]
	if !ok {
		return nil, fmt.Errorf("environment %q not found", name)
	}
	return &env, nil
}

// SaveToFile saves the configuration to the file at the given path.
//
// The file contents are a JSON representation of the configuration, indented with two spaces.
// The file is saved with 644 permissions.
func (c *Config) SaveToFile(filePath string) error {
	marshaledConfig, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, marshaledConfig, 0o644)
}

// Merge merges the environments from another configuration into the current configuration.
// 
// For each environment in the other configuration, if the environment does not exist in the
// current configuration, it is added. If it exists, its properties (APIKey, Username, Password, 
// BaseURL) are updated only if they are non-empty in the other configuration.

func (c *Config) Merge(other *Config) {
	for envName, otherEnv := range other.Environments {
		env, exists := c.Environments[envName]
		if !exists {
			c.Environments[envName] = otherEnv
			continue
		}

		if otherEnv.Auth.APIKey != "" {
			env.Auth.APIKey = otherEnv.Auth.APIKey
		}

		if otherEnv.Auth.Username != "" {
			env.Auth.Username = otherEnv.Auth.Username
		}

		if otherEnv.Auth.Password != "" {
			env.Auth.Password = otherEnv.Auth.Password
		}

		if otherEnv.BaseURL != "" {
			env.BaseURL = otherEnv.BaseURL
		}

		c.Environments[envName] = env
	}
}
