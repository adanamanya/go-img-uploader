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

/*initializing firebase for better future*/

func FirebaseStorage() *storage.Client {
	//define storage bucket, calling from env
	config := &firebase.Config{
		StorageBucket: os.Getenv("STORAGE_BUCKET"),
	}
	//we need cred to connect with firebase
	opt := option.WithCredentialsFile("key.json")
	//bond connection with firebase
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf(errors.ErrFailedInit.Error(), err)
	}
	// init firebase storage
	storage, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalf(errors.ErrStorageinit.Error(), err)
	}

	return storage
}
