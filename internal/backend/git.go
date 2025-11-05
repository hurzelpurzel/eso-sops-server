package backend

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/hurzelpurzel/eso-sops-server/internal/config"
	"os"
)

func CloneRepo(config *config.Config, sec *config.Secret) error {
	defer os.RemoveAll(config.RepoDir) // clean up

	fmt.Println("Cloning repository...")
	_, err := git.PlainClone(config.RepoDir, false, &git.CloneOptions{
		URL: config.RepoURL,
		Auth: &http.BasicAuth{
			Username: sec.GitUser, // anything except an empty string
			Password: sec.GitToken,
		},
		Progress:        os.Stdout,
		InsecureSkipTLS: true,
	})
	return err

}
