package backend

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/hurzelpurzel/eso-sops-server/internal/config"
)


type GitBackend struct {
	BasePath  string
	Type      string
	GitCredentials *config.GitCredentials
	Repos     []config.Repo
}

func CreateGit (cfg *config.Config) (*GitBackend, error){
	git := GitBackend{}
	err := git.Init(cfg)
	return &git,err
}

func (g *GitBackend) Init(cfg *config.Config) error {
	g.BasePath = cfg.CheckoutDir
	g.Type = "git"
	var err error
	g.GitCredentials, err = config.LoadGitConfig()
	g.Repos = cfg.Repos
	return err
}

func (g *GitBackend) getPath() string {
	return g.BasePath +"/"+g.Type
}

func (g *GitBackend) DownloadAll() error {
	for _, rep := range g.Repos {
		if err := g.DownloadByName(rep.Name); err != nil {
			return err
		}
	}
	return nil
}

func (g *GitBackend) GetRepoByName(name string) *config.Repo {
	for _, rep := range g.Repos {
		if rep.Name == name {
			return &rep
		}
	}
	return nil
}

func (g *GitBackend) DownloadByName(name string) error {
	repodir := g.getPath() + "/" + name
	_ = os.RemoveAll(repodir) // clean up
	_ = os.MkdirAll(repodir, 0755)
	repcfg := g.GetRepoByName(name)
	sec := g.GitCredentials.GetSecretByName(repcfg.Profile)	
	fmt.Printf("Cloning repository to...%s", repodir)	
	repo, err := git.PlainClone(repodir, false, &git.CloneOptions{
		URL: repcfg.URL,
		Auth: &http.BasicAuth{
			Username: sec.GitUser, // anything except an empty string
			Password: sec.GitToken,
		},
		Progress:        os.Stdout,
		RemoteName:      "origin",
		InsecureSkipTLS: true,
		SingleBranch:    true,
		ReferenceName:   plumbing.ReferenceName(repcfg.Branch),
	})
	_, _ = repo.Head()

	return err
}
