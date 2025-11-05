package config

import (
	
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
    RepoURL string `yaml:"repo_url"`
	RepoDir string `yaml:"repo_dir"`
}

type Secret struct {
    GitUser string `yaml:"git_user"`
    GitToken string `yaml:"git_token"`
	AgeKey string `yaml:"age_key"`
}

func LoadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var cfg Config
    err = yaml.Unmarshal(data, &cfg)
    return &cfg, err
}


func LoadSecret(path string) (*Secret, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var sec Secret
    err = yaml.Unmarshal(data, &sec)
    return &sec, err
}

