package decrypt

import (
    "bytes"
    "fmt"
    
    "os"
    "os/exec"
    
    "path/filepath"
    "strings"

   
    "github.com/hurzelpurzel/eso-sops-server/internal/config"
)



/* Expects sops binary in container*/
func decryptSOPS(filePath string, agekey string) (string, error) {
    err := os.Setenv("SOPS_AGE_KEY", agekey)
    if err != nil {
        return "", err
    }

    cmd := exec.Command( "sops", "-d", filePath)
    
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
        return "", err
    }
    _ = os.Setenv("SOPS_AGE_KEY", "" )

    return out.String(), nil
}

func GetDecryptedJson(config *config.Config, user *config.User, filename string) ( * string, error) {
   
    fullPath := filepath.Join(config.CheckoutDir, filename)

    if !strings.HasSuffix(fullPath, ".json") {
        return nil, fmt.Errorf("file %s is not a JSON file", fullPath)
    }

    content, err := decryptSOPS(fullPath, user.AgeKey)
    if err != nil {
        return nil, fmt.Errorf("error decrypting %s: %w", fullPath, err)
    }

    
    return &content, nil
}


/*
func GetDecryptedJson(config *config.Config , user *config.User, filename string) (map[string]string, error) {
    result := make(map[string]string)
    err := filepath.WalkDir(filepath.Join(config.CheckoutDir, filename), func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() && strings.HasSuffix(path, ".json") {
            content, err := decryptSOPS(path,user.AgeKey)
            if err != nil {
                return fmt.Errorf("error on decrypt %s: %w", path, err)
            }
            result[filepath.Base(path)] = content
        }
        return nil
    })
    return result, err
}*/