package firebaseapi

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseAppClinet *FirebaseApp

type FirebaseApp struct {
	ctx context.Context
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
	firebaseApp.ctx = ctx
	firebaseApp.db = newFireStoreClient(ctx, app)
	return firebaseApp
}

func (f *FirebaseApp) Close() {
	f.db.fireStoreClose()
}

type ReadMatchQueryOption struct {
	Season      string
	Team        string
	Result      string
	SortOpt     string
	PublishedAt string
	Limit       string
}

func NewReadMatchQueryOption(season, team, result, sortOpt, publishedAt, limit string) *ReadMatchQueryOption {
	return &ReadMatchQueryOption{
		Season:      season,
		Team:        team,
		Result:      result,
		SortOpt:     sortOpt,
		PublishedAt: publishedAt,
		Limit:       limit,
	}
}

const LCK_MATCH_COLLECTION = "lck_match"

func (f *FirebaseApp) ReadMatchTeamWithQueryOption(opt *ReadMatchQueryOption) []FireStoreDataSchema {
	return f.db.readMatchTeamWithQueryOption(f.ctx, opt)
}
