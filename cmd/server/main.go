package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hurzelpurzel/eso-sops-server/internal/backend"
	"github.com/hurzelpurzel/eso-sops-server/internal/config"
	"github.com/hurzelpurzel/eso-sops-server/internal/decrypt"
)

func main() {
	cfg, _ := config.LoadConfig("/home/ludger/testdir/config/config.yaml")
	secret, _ := config.LoadSecret("/home/ludger/testdir/config/secret.yaml")
	if err := backend.CloneRepo(cfg, secret); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/keys", func(c *gin.Context) {
		data, err := decrypt.GetDecryptedJson(cfg, secret)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	})

	r.Run(":8080")
}
