package backend

import (
	"os"

	"github.com/hurzelpurzel/eso-sops-server/internal/config"
)

type OthersBackend struct {
	BasePath  string
	Type      string
	Others     []config.Other
}

func CreateOthers (cfg *config.Config) (OthersBackend, error){
	ot := OthersBackend{}
	err := ot.Init(cfg)
	return ot,err
}

func (g *OthersBackend) Init(cfg *config.Config) error {
	g.BasePath = cfg.CheckoutDir
	g.Type = "other"
	g.Others = cfg.Others
	return nil
}

func (g OthersBackend) DownloadAll() error{
	for _, oth := range g.Others {
		if err := g.DownloadByName(oth.Name); err != nil {
			return err
		}
	}	
 	return nil
}

func (g OthersBackend) DownloadByName(name string) error{
	otherdir := g.GetPath() + "/" + name
	if !dirExists(otherdir) {
		return os.MkdirAll(otherdir, 0755)	
	}
 	return nil
}

func dirExists(path string) bool {
    info, err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    }
    return err == nil && info.IsDir()
}

func (g OthersBackend) GetPath() string {
	return g.BasePath +"/"+g.Type
}
