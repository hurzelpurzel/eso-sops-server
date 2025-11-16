package backend

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	cf "github.com/hurzelpurzel/eso-sops-server/internal/config"
)

type S3Backend struct {
	BasePath string
	Type     string
	Buckets  []cf.Bucket
}

func CreateS3(cfg *cf.Config) (*S3Backend, error) {
	git := S3Backend{}
	err := git.Init(cfg)
	return &git, err
}

func (g *S3Backend) Init(cfg *cf.Config) error {
	g.BasePath = cfg.CheckoutDir
	g.Type = "s3"
	var err error
	g.Buckets = cfg.Buckets
	return err
}

func (s *S3Backend) DownloadAll() error {
	for _, bucket := range s.Buckets {
		if err := s.DownloadByName(bucket.Name); err != nil {
			return err
		}
	}
	return nil

}

func (s *S3Backend) DownloadByName(name string) error {
	bucketCfg := s.GetBucketByName(name)
	client, err := getClient(bucketCfg.Profile)
	if err != nil {
		return fmt.Errorf("failed to create S3 client: %w", err)
	}
	bucket := bucketCfg.Name
	resp, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucket,
	})
	if err != nil {
		return fmt.Errorf("failed to list objects: %w", err)
	}

    bucketdir := s.getPath() + "/" + name
	_ = os.RemoveAll(bucketdir) // clean up
	_ = os.MkdirAll(bucketdir, 0755)

	for _, item := range resp.Contents {
		if err := downloadFromS3(client,bucketdir, bucket, item); err != nil {
			log.Printf("error downloading %s: %v", *item.Key, err)
		}
	}

	return nil

}

func (s *S3Backend) GetBucketByName(name string) *cf.Bucket {
	for _, buck := range s.Buckets {
		if buck.Name == name {
			return &buck
		}
	}
	return nil
}

func getClient(profile string) (*s3.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %w", err)
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}

func (g *S3Backend) getPath() string {
	return g.BasePath +"/"+g.Type
}


// downloadFromS3 lädt ein einzelnes Objekt herunter
func downloadFromS3(client *s3.Client,path string, bucket string, item types.Object) error {
	fmt.Printf("Downloading: %s\n", *item.Key)

	getObj, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    item.Key,
	})
	if err != nil {
		log.Printf("failed to get object %s, %v", *item.Key, err)
		return err
	}
	defer func() {
		if err := getObj.Body.Close(); err != nil {
			log.Printf("failed to close object body: %v", err)
		}
	}()

	// Lokale Datei anlegen
    filename := path + "/" + *item.Key
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("failed to create file %s, %v", filename, err)
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}()

	_, err = io.Copy(f, getObj.Body)
	if err != nil {
		log.Printf("failed to copy object %s, %v", *item.Key, err)
		return err
	}

	return nil
}
