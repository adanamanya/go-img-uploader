package repository

import (
	"context"
	"gogogo/pkg/errors"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

func FirebaseStorage() *storage.Client {
	config := &firebase.Config{
		StorageBucket: os.Getenv("STORAGE_BUCKET"),
	}
	opt := option.WithCredentialsFile("key.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf(errors.ErrFailedInit.Error(), err)
	}

	storage, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalf(errors.ErrStorageinit.Error(), err)
	}

	return storage
}
