package config

import (
	"os"
   
	"gopkg.in/yaml.v2"
)






func GetConfigPath() string {
    if value, exists := os.LookupEnv("CONFIG_PATH"); exists {
        return value
    }
    return "/home/ludger/testdir/config"
}

func LoadConfig(file string) (* Config, error) {
    path := GetConfigPath() + "/" + file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}

func LoadSecret(file string) (*Secret, error) {
    path := GetConfigPath() + "/" + file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var sec Secret
	err = yaml.Unmarshal(data, &sec)
	return &sec, err
}

func LoadUsers(file string) (*Users, error) {
    path := GetConfigPath() + "/" + file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var users Users
	err = yaml.Unmarshal(data, &users)
	return &users, err
}

