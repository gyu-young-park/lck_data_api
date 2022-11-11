package firebaseapi

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	ctx *context.Context
	db  *fireStoreClient
}

func NewFireBaseAPI(secretFirebaseServiceAccountKeyPath string) *FirebaseApp {
	firebaseApp := &FirebaseApp{}
	ctx := context.Background()
	conf := &firebase.Config{}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(secretFirebaseServiceAccountKeyPath)

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}
	firebaseApp.ctx = &ctx
	firebaseApp.db = newFireStoreClient(ctx, app)
	return firebaseApp
}

func (f *FirebaseApp) Close() {
	f.db.fireStoreClose()
}

type ReadMatchQueryOption struct {
	Season  string
	Team    string
	Result  string
	SortOpt string
	Start   string
	End     string
}

func NewReadMatchQueryOption(season, team, result, sortOpt, start, end string) *ReadMatchQueryOption {
	return &ReadMatchQueryOption{
		Season:  season,
		Team:    team,
		Result:  result,
		SortOpt: sortOpt,
		Start:   start,
		End:     end,
	}
}

func (f *FirebaseApp) ReadMatchWithQueryOption(collection string, opt *ReadMatchQueryOption) []FireStoreDataSchema {
	var ret []FireStoreDataSchema
	iter := f.db.client.Collection(collection).Documents(*f.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("Failed to iterate %v\n", err)
			break
		}
		ret = append(ret, doc.Data())
	}
	return ret
}
