package backend

import (
	"context"
	"fmt"
	"os"
	"log"

	"github.com/hurzelpurzel/eso-sops-server/internal/config"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
)

type OrasBackend struct {
	BasePath       string
	Type           string
	OciRegistrys    []config.OciRegistry
}

func CreateOras(cfg *config.Config) (OrasBackend, error) {
	be := OrasBackend{}
	err := be.Init(cfg)
	return be, err
}

func (g *OrasBackend) Init(cfg *config.Config) error {
	g.BasePath = cfg.CheckoutDir
	g.Type = "oras"
	var err error
	g.OciRegistrys = cfg.OciRegistrys
	return err
}

func (g OrasBackend) GetPath() string {
	return g.BasePath + "/" + g.Type
}

func (g OrasBackend) DownloadAll() error {
	for _, rep := range g.OciRegistrys {
		if err := g.DownloadByName(rep.Name); err != nil {
			return err
		}
	}
	return nil
}

func (g OrasBackend) GetOciRegistryByName(name string) *config.OciRegistry {
	for _, rep := range g.OciRegistrys {
		if rep.Name == name {
			return &rep
		}
	}
	return nil
}

func (g OrasBackend) DownloadByName(name string) error {
	repodir := g.GetPath() + "/" + name
	_ = os.RemoveAll(repodir) // clean up
	_ = os.MkdirAll(repodir, 0755)
	repcfg := g.GetOciRegistryByName(name)
    ctx := context.Background()

	// Beispiel: Referenz auf ein Artefakt in einer Registry
	// Format: <REGISTRY>/<PROJECT>/<REPO>/<NAME>:<TAG>
	artifactRef := repcfg.Hostname+"/"+repcfg.Repository+"/"+repcfg.Image+":"+repcfg.Tag

	
	// Lokales File-Store Backend
	fs, err := file.New(repodir)
	if err != nil {
		return fmt.Errorf("failed to create local folder: %w", err)
	}
	
	defer func() {
        if cerr := fs.Close(); cerr != nil {
            fmt.Printf("failed to close file: %v\n", cerr)
        }
    }()

	// Remote-Repository erstellen
	repo, err := remote.NewRepository(artifactRef)
	if err != nil {
	    return fmt.Errorf("failed to create remote: %w", err)
	}

	// Registry-Login konfigurieren
	// Für Google Artifact Registry: Access Token über gcloud holen
	//   gcloud auth print-access-token
	// und als Passwort setzen, Benutzername = "oauth2accesstoken"
	accessToken,err := config.LoadOrasToken()
	if err != nil {
		return fmt.Errorf("unable to load token: %w", err)
	}
	repo.Client = &auth.Client{
		Credential: func(ctx context.Context, host string) (auth.Credential, error) {
			return auth.Credential{
				Username: "oauth2accesstoken",
				Password: accessToken,
			}, nil
		},
	}

	// Artefakt kopieren: Registry → lokales Filesystem
	desc, err := oras.Copy(ctx, repo, artifactRef, fs, artifactRef, oras.DefaultCopyOptions)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	log.Printf("Artefakt succcessful copied: %s\n", desc.Digest)
	log.Printf("Files are now in: %s\n", repodir)
	return err
}
