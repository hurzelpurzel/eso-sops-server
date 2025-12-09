package config

import (
	"os"

	"github.com/hurzelpurzel/eso-sops-server/internal/utils"
	"gopkg.in/yaml.v2"
)

const EnvGitCredFile = "GIT_CREDS_FILE"
const EnvGitUsersYaml = "USERS_FILE"
const EnvConfigYaml = "CONFIG_FILE"
const EnvOrasTokenFile = "ORAS_TOKEN_FILE"



func LoadConfig() (* Config, error) {
    path, err := utils.GetEnvOrFail(EnvConfigYaml) 
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}

func LoadOrasToken() (string, error) {
	path, err := utils.GetEnvOrFail(EnvOrasTokenFile)	
    if err != nil {
		return "", err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func LoadUsers() (*Users, error) {
    path, err := utils.GetEnvOrFail(EnvGitUsersYaml) 
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var users Users
	err = yaml.Unmarshal(data, &users)
	return &users, err
}

/*-----------------------------------*/

func LoadGitConfig() (*GitCredentials, error) {
    path, err := utils.GetEnvOrFail(EnvGitCredFile) 
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var sec GitCredentials
	err = yaml.Unmarshal(data, &sec)
	return &sec, err
}

