package backend
/*
import (
    "context"
    "fmt"
    "io"
    "os"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func main() {
    bucket := "your-bucket-name"
    key := "path/to/your/file.txt"
    destination := "local-file.txt"

    // Load AWS config
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        panic("unable to load SDK config, " + err.Error())
    }

    // Create S3 client
    client := s3.NewFromConfig(cfg)

    // Get the object
    output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
        Bucket: &bucket,
        Key:    &key,
    })
    if err != nil {
        panic("failed to get object, " + err.Error())
    }
    defer output.Body.Close()

    // Create local file
    file, err := os.Create(destination)
    if err != nil {
        panic("failed to create file, " + err.Error())
    }
    defer file.Close()

    // Write content to local file
    _, err = io.Copy(file, output.Body)
    if err != nil {
        panic("failed to copy content, " + err.Error())
    }

    fmt.Println("File downloaded successfully:", destination)
}
*/