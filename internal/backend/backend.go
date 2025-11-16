package backend

import "github.com/hurzelpurzel/eso-sops-server/internal/config"



type Backend interface {
    Init( cfg *config.Config )  error
    DownloadAll() error
    DownloadByName(name string) error
}



