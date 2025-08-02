package storage

import (
	"bytes"
	"context"
	"fmt"
	"gym-map/config"
	"gym-map/store"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3FileInfo struct {
	name    string
	size    int64
	modTime time.Time
}

// Implement FileInfo interface
func (sfi S3FileInfo) Name() string       { return sfi.name }
func (sfi S3FileInfo) Size() int64        { return sfi.size }
func (sfi S3FileInfo) Mode() os.FileMode  { return 0 }
func (sfi S3FileInfo) ModTime() time.Time { return sfi.modTime }
func (sfi S3FileInfo) IsDir() bool        { return false }
func (sfi S3FileInfo) Sys() any           { return nil }

type S3Storage struct {
	Config config.Config
	Client *s3.Client
}

type bufferReadSeekCloser struct {
	*bytes.Reader
}

func (b *bufferReadSeekCloser) Close() error {
	return nil
}

func GetS3Client(config config.Config) *s3.Client {
	awsConfig, err := awsconfig.LoadDefaultConfig(
		context.TODO(),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				config.StorageS3AccessKey,
				config.StorageS3SecretKey,
				"",
			),
		),
		awsconfig.WithRegion(config.StorageS3Region),
	)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return s3.NewFromConfig(awsConfig, func(o *s3.Options) {
		o.BaseEndpoint = &config.StorageS3Endpoint
	})
}

func (ss S3Storage) getPath(fileType store.FileType, name string) string {
	return filepath.Join(string(fileType), name)
}

func (ss S3Storage) Read(fileType store.FileType, name string) (*store.StorageObject, error) {
	path := ss.getPath(fileType, name)
	objectInput := s3.GetObjectInput{
		Bucket: &ss.Config.StorageS3BucketName,
		Key:    &path,
	}
	output, err := ss.Client.GetObject(context.Background(), &objectInput)
	if err != nil {
		return nil, err
	}

	fileInfo := S3FileInfo{
		name:    name,
		size:    *output.ContentLength,
		modTime: *output.LastModified,
	}

	// TODO: only send the reader
	data, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	return &store.StorageObject{
		ReadSeekCloser: &bufferReadSeekCloser{bytes.NewReader(data)},
		FileInfo:       fileInfo,
	}, nil
}

func (ss S3Storage) Write(fileType store.FileType, data io.ReadSeeker, name string) error {
	path := ss.getPath(fileType, name)

	uploader := manager.NewUploader(ss.Client)
	objectInput := s3.PutObjectInput{
		Bucket: &ss.Config.StorageS3BucketName,
		Key:    &path,
		Body:   data,
	}

	_, err := uploader.Upload(context.Background(), &objectInput)
	if err != nil {
		return fmt.Errorf("S3 PutObject failed: %w", err)
	}
	return nil
}

func (ss S3Storage) Remove(fileType store.FileType, name string) error {
	path := ss.getPath(fileType, name)

	deleteInput := s3.DeleteObjectInput{
		Bucket: &ss.Config.StorageS3BucketName,
		Key:    &path,
	}

	_, err := ss.Client.DeleteObject(context.Background(), &deleteInput)
	if err != nil {
		return err
	}

	return nil
}
