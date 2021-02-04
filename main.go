package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func main() {
	bucket := os.Getenv("BUCKET_NAME")
	localObjectPath := os.Getenv("LOCAL_OBJECT_PATH")
	remoteObjectPath := os.Getenv("REMOTE_OBJECT_PATH")
	googleApplicationCredentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	opts := option.WithCredentialsFile(googleApplicationCredentials)
	ctx := context.Background()
	client, err := storage.NewClient(ctx, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "storage.NewClient: %v", err)
	}
	defer client.Close()

	// Open local file.
	f, err := os.Open(localObjectPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open: %v", err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(remoteObjectPath).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		fmt.Fprintf(os.Stderr, "io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "Writer.Close: %v", err)
	}
	fmt.Fprintf(os.Stdout, "Blob %v uploaded.\n", remoteObjectPath)
}
