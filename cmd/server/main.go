package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hurzelpurzel/eso-sops-server/internal/backend"
	"github.com/hurzelpurzel/eso-sops-server/internal/config"
	"github.com/hurzelpurzel/eso-sops-server/internal/decrypt"
	"github.com/hurzelpurzel/eso-sops-server/internal/utils"
)

var git backend.GitBackend
var s3 backend.S3Backend
var other backend.OthersBackend
var oras backend.OrasBackend

func main() {
	cfg, err := config.LoadConfig()
	utils.CheckErr(err)

	users, er := config.LoadUsers()
	utils.CheckErr(er)

	

	// Download all repositories and buckets at startup
	if len(cfg.Repos) > 0 {
		gitback, err := backend.CreateGit(cfg)
		utils.CheckErr(err)
		err = gitback.DownloadAll()
		utils.CheckErr(err)
		git = gitback
	}
	
	if len(cfg.Buckets) > 0 {
		s3back, err := backend.CreateS3(cfg)
		utils.CheckErr(err)
		err = s3back.DownloadAll()
		utils.CheckErr(err)
		s3 = s3back
	}

	if len(cfg.OciRegistrys) > 0 {
		orasback, err := backend.CreateOras(cfg)
		utils.CheckErr(err)
		err = orasback.DownloadAll()
		utils.CheckErr(err)
		oras = orasback
	}

	if len(cfg.Others) > 0 {
		otherback, err := backend.CreateOthers(cfg)
		utils.CheckErr(err)
		err = otherback.DownloadAll()
		utils.CheckErr(err)
		other = otherback
	}

	// Set up Gin router
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
		err := s3.DownloadAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = git.DownloadAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "repositories initialized"})
	})

	authorized.GET("/git/:repo/:filepath", func(c *gin.Context) {
		username := c.MustGet(gin.AuthUserKey).(string)
		repo := c.Param("repo")
		filepath := c.Param("filepath")
		filename := repo + "/" + filepath
		data, err := decrypt.GetDecryptedJson(git, users.GetUserByName(username), filename)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	})

	authorized.GET("/s3/:bucket/:filepath", func(c *gin.Context) {
		username := c.MustGet(gin.AuthUserKey).(string)
		bucket := c.Param("bucket")
		filepath := c.Param("filepath")
		filename := bucket + "/" + filepath
		data, err := decrypt.GetDecryptedJson(s3, users.GetUserByName(username), filename)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	})

    authorized.GET("/oras/:reg/:image/:filepath", func(c *gin.Context) {
		username := c.MustGet(gin.AuthUserKey).(string)
		reg := c.Param("reg")
		image := c.Param("image")
		filepath := c.Param("filepath")
		filename := reg + "/" + image + "/" + filepath
		data, err := decrypt.GetDecryptedJson(oras, users.GetUserByName(username), filename)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	})

	authorized.GET("/other/:folder/:filepath", func(c *gin.Context) {
		username := c.MustGet(gin.AuthUserKey).(string)
		folder := c.Param("folder")
		filepath := c.Param("filepath")
		filename := folder + "/" + filepath
		data, err := decrypt.GetDecryptedJson(other, users.GetUserByName(username), filename)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	err = r.Run(":8080")
	utils.CheckErr(err)
}
