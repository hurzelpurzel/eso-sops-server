package backend

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/hurzelpurzel/eso-sops-server/internal/config"
	"os"
)


func CloneRepo(config *config.Config, sec *config.Secret, repcfg *config.Repo ) error {
    repodir := config.CheckoutDir + "/"+ repcfg.Name
	os.RemoveAll(repodir) // clean up
    os.MkdirAll(repodir, 0755)
	
	fmt.Printf("Cloning repository to...%s", repodir)
	repo , err := git.PlainClone(repodir, false, &git.CloneOptions{
		URL: repcfg.URL,
		Auth: &http.BasicAuth{
			Username: sec.GitUser, // anything except an empty string
			Password: sec.GitToken,
		},
		Progress:        os.Stdout,
        RemoteName: "origin",
		InsecureSkipTLS: true,
        SingleBranch: true ,
		ReferenceName: plumbing.ReferenceName(repcfg.Branch),
	})
    repo.Head()

	return err

}
