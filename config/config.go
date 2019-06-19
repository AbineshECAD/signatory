package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// ServerConfig contains the information necessary to the tezos signing server
type ServerConfig struct {
	Port        int `yaml:"port"`
	UtilityPort int `yaml:"utility_port"`
}

// YubiConfig contains the information necessary to use the Yubi HSM backend
type YubiConfig struct {
	Host      string `yaml:"host"`
	Password  string `yaml:"password"`
	AuthKeyID uint16 `yaml:"auth_key_id"`
}

// AzureConfig contains the information necessary to use the Azure Key Vault backend
type AzureConfig struct {
	ClientID       string   `yaml:"client_id"`
	ClientSecret   string   `yaml:"client_secret"`
	DirectoryID    string   `yaml:"tenant_id"`
	SubscriptionID string   `yaml:"subscription"`
	ResourceGroup  string   `yaml:"resource_group"`
	Vault          string   `yaml:"vault"`
	Keys           []string `yaml:"keys"`
}

// TezosConfig contains the configuration related to tezos network
type TezosConfig struct {
	Keys              []string
	AllowedOperations []string `yaml:"allowed_operations"`
	AllowedKinds      []string `yaml:"allowed_kinds"`
	LogPayloads       bool     `yaml:"log_payloads"`
}

// Config contains all the configuration necessary to run the signatory
type Config struct {
	Yubi   []YubiConfig  `yaml:"yubi"`
	Azure  []AzureConfig `yaml:"azure"`
	Tezos  TezosConfig   `yaml:"tezos"`
	Server ServerConfig  `yaml:"server"`
}

// Read read the config from a file
func (c *Config) Read(file string) error {
	yamlFile, _ := ioutil.ReadFile(file)
	err := yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return fmt.Errorf("Unmarshal: %v", err)
	}

	return nil
}
