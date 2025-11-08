package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hurzelpurzel/eso-sops-server/internal/backend"
	"github.com/hurzelpurzel/eso-sops-server/internal/config"
	"github.com/hurzelpurzel/eso-sops-server/internal/decrypt"
)

func initRepo(config *config.Config, sec *config.Secret) {
	for _, rep := range config.Repos {
		if err := backend.CloneRepo(config, sec, &rep); err != nil {
			panic(err)
		}

	}
}

func main() {
	cfg, _ := config.LoadConfig("config.yaml")
	secret, _ := config.LoadSecret("git.yaml")
	users, _ := config.LoadUsers("users.yaml")

	initRepo(cfg, secret)

	r := gin.Default()

	// Define authorized users
	authorized := r.Group("/", gin.BasicAuth(*users.ToAccounts()))

	authorized.GET("/init", func(c *gin.Context) {
		username := c.MustGet(gin.AuthUserKey).(string)
		user := users.GetUserByName(username)
		if !user.HasRole("admin") {
			c.JSON(403, gin.H{"error": "forbidden"})
			return
		}
		initRepo(cfg, secret)
		c.JSON(200, gin.H{"status": "repository initialized"})
	})

	authorized.GET("/git/:repo/:filepath", func(c *gin.Context) {
		username := c.MustGet(gin.AuthUserKey).(string)
		repo := c.Param("repo")
		filepath := c.Param("filepath")
		filename := repo + "/" + filepath
		data, err := decrypt.GetDecryptedJson(cfg, users.GetUserByName(username), filename)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}
