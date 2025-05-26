package util

import (
    "os"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

type ElasticsearchInfo struct {
    Address  string `yaml:"address"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
}

type GCSInfo struct {
    Bucket string `yaml:"bucket"`
}

type TokenInfo struct {
    Secret string `yaml:"secret"`
}

type StripeInfo struct {
	APIKey string `yaml:"api_key"`
}

type ApplicationConfig struct {
    ElasticsearchConfig *ElasticsearchInfo `yaml:"elasticsearch"`
    GCSConfig           *GCSInfo           `yaml:"gcs"`
    TokenConfig         *TokenInfo         `yaml:"token"`
	StripeConfig        *StripeInfo        `yaml:"stripe"`
}

func LoadApplicationConfig(configDir, configFile string) (*ApplicationConfig, error) {
    content, err := os.ReadFile(filepath.Join(configDir, configFile))
    if err != nil {
        return nil, err
    }

    var config ApplicationConfig
    err = yaml.Unmarshal(content, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}