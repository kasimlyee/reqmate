package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// LoadConfig loads a configuration from a given path.
//
// If the path is empty, it will attempt to load a configuration from the
// following locations:
//
//   - The current working directory
//   - $HOME/.config/reqmate
//
// The environment variables prefixed with "REQMATE_" will override the
// values in the configuration.
//
// If the configuration file does not exist, it will return a default configuration.
//
func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetEnvPrefix("REQMATE")

	if path != "" {
		v.SetConfigFile(path)
	} else {
		v.SetConfigName(".reqmate")
		v.AddConfigPath(".")
		v.AddConfigPath("$HOME/.config/reqmate")
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	config = processEnvOverrides(config)

	return &config, nil
}

// processEnvOverrides processes the environment variables and overrides the
// values in the given configuration.
//
// The only variable that is overridden is the API key for each environment.
// The environment variable name is generated by calling envAPIKeyVarName for
// the given environment name.
//
// If the environment variable is not set, the value in the configuration is
// left unchanged.
func processEnvOverrides(cfg Config) Config {
	for envName, env := range cfg.Environments {
		auth := env.Auth

		if apiKey := os.Getenv(envAPIKeyVarName(envName)); apiKey != "" {
			auth.APIKey = apiKey
		}

		cfg.Environments[envName] = env
	}

	return cfg
}

// envAPIKeyVarName returns the environment variable name for the API key of the given environment name.
//
// The environment variable name is in the format "REQMATE_<envName>_APIKEY", where <envName> is the environment name.
func envAPIKeyVarName(envName string) string {
	return fmt.Sprintf("REQMATE_%s_APIKEY", envName)
}
