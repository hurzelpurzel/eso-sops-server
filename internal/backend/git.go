package backend

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/hurzelpurzel/eso-sops-server/internal/config"
)

func CloneRepo(config * config.Config ) error {
    if _, err := os.Stat(config.RepoDir); os.IsNotExist(err) {
        fmt.Println("Cloning repository...")
        _, err := git.PlainClone(config.RepoDir, false, &git.CloneOptions{
            URL:      config.RepoURL,
            Progress: os.Stdout,
        })
        return err
    }
    fmt.Println("Repository already cloned.")
    return nil
}