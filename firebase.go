package main

import (
	"cloud.google.com/go/storage"
	"context"

	firebase "firebase.google.com/go"
	_ "firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func InitBucket() (*storage.BucketHandle, error) {
	config := &firebase.Config{
		StorageBucket: BucketName,
	}
	opt := option.WithCredentialsFile(Credentials)
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return nil, err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
