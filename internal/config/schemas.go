package config

type Config struct {
	Version    string             `yaml:"version" json:"version"`
	Environments map[string]Environment `yaml:"environments" json:"environments"`
	Requests     []RequestDefinition   `yaml:"requests" json:"requests"`
	Variables    map[string]interface{} `yaml:"variables" json:"variables"`
}

type Environment struct {
	BaseURL    string            `yaml:"base_url" json:"base_url"`
	Headers    map[string]string `yaml:"headers" json:"headers"`
	Auth       AuthConfig        `yaml:"auth" json:"auth"`
	Variables  map[string]interface{} `yaml:"variables" json:"variables"`
}

type AuthConfig struct {
	Type     string `yaml:"type" json:"type"`
	APIKey   string `yaml:"api_key" json:"api_key"`
	Token    string `yaml:"token" json:"token"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type RequestDefinition struct {
	Name        string            `yaml:"name" json:"name"`
	Method      string            `yaml:"method" json:"method"`
	Endpoint    string            `yaml:"endpoint" json:"endpoint"`
	Headers     map[string]string `yaml:"headers" json:"headers"`
	Body        interface{}       `yaml:"body" json:"body"`
	//Tests       []TestAssertion   `yaml:"tests" json:"tests"`
	Extract     map[string]string `yaml:"extract" json:"extract"`
}