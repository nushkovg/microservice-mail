package config

import (
	"encoding/json"
	"io/ioutil"
)

// TemplateConfig parsed data from configuration
type TemplateConfig struct {
	TemplateName string            `json:"template,omitempty"`
	Data         map[string]string `json:"data,omitempty"`
}

// Config holds the microservice full configuration.
type Config struct {
	// VerificationURL contains the url to the verification action
	Template map[string]TemplateConfig `json:"template"`

	// Mail is a map of <property>:<value>. For example,
	// "host": "smtp.example.com"
	Mail map[string]string `json:"mail"`

	// RabbitMQ holds information about the rabbitmq server
	RabbitMQ map[string]string `json:"rabbitmq"`
}

// LoadConfig loads a Config from a configuration JSON file.
func LoadConfig(confFile string) (*Config, error) {
	confBytes, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(confBytes, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
