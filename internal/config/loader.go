package config

import (
    "os"
    "gopkg.in/yaml.v2"
)

type Config struct {
    RepoURL string `yaml:"repo_url"`
	RepoDir string `yaml:"repo_dir"`
    YamlDir string `yaml:"yaml_dir"`
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


/*
repoURL     = "git@github.com:dein-user/dein-repo.git" // oder HTTPS
    repoDir     = "./repo"
    yamlDir     = "configs" // Pfad im Repo zu den YAML-Dateien
	// 
	// */