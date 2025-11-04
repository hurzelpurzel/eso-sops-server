package decrypt

import (
    "bytes"
    "fmt"
    "io/fs"
    "os"
    "os/exec"
    
    "path/filepath"
    "strings"

   
    "github.com/hurzelpurzel/eso-sops-server/internal/config"
)

/* Expects sops binary in container*/
func decryptSOPS(filePath string) (string, error) {
    cmd := exec.Command("sops", "-d", filePath)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        return "", err
    }
    return out.String(), nil
}

func GetDecryptedYAMLs(config *config.Config ) (map[string]string, error) {
    result := make(map[string]string)
    err := filepath.WalkDir(filepath.Join(config.RepoDir, config.YamlDir), func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() && strings.HasSuffix(path, ".yaml") {

            content, err := decryptSOPS(path)
            if err != nil {
                return fmt.Errorf("fehler beim entschluesseln von %s: %w", path, err)
            }
            result[filepath.Base(path)] = content
        }
        return nil
    })
    return result, err
}