package config

import "github.com/gin-gonic/gin"

type Config struct {
	CheckoutDir    string `yaml:"checkout_dir"`
	Repos	[] Repo  `yaml:"repos"`
}

type Repo struct{
	URL    string `yaml:"url"`
	Branch string `yaml:"branch"`
	Name  string `yaml:"name"`

}

type Secret struct {
	GitUser  string `yaml:"git_user"`
	GitToken string `yaml:"git_token"`
}

type Users struct {
	Users []User `yaml:"users"`
}

type User struct {
	Name     string   `yaml:"name"`
	Password string   `yaml:"password"`
	AgeKey   string   `yaml:"age_key"`
	Roles    []string `yaml:"roles"`
}

func (u *Users) GetUserByName(name string) *User {
	for _, user := range u.Users {
		if user.Name == name {
			return &user
		}
	}
	return nil
}

func (cfg *Config ) GetRepoByName(name string) *Repo {
	for _, rep := range cfg.Repos {
		if rep.Name == name {
			return &rep
		}
	}
	return nil
}

func (u *Users) ToAccounts() *gin.Accounts {
	accounts := make(gin.Accounts)
	for _, user := range u.Users {
		accounts[user.Name] = user.Password
	}
	return &accounts
}

func (u User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}
