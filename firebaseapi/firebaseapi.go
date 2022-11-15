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

func (f *FirebaseApp) ReadMatchTeamWithQueryOption(opt *ReadMatchQueryOption) []FireStoreDataSchema {
	return f.db.readMatchTeamWithQueryOption(f.ctx, opt)
}

func (f *FirebaseApp) ReadSeaonTeamWithQueryOption(opt *ReadTeamWithSeasonQueryOption) []FireStoreDataSchema {
	return f.db.readTeamWithSeason(f.ctx, opt)
}

func (f *FirebaseApp) ReadSeasonList() FireStoreDataSchema {
	return f.db.readAllLCKSeason(f.ctx)
}

func (f *FirebaseApp) ReadTeamList() FireStoreDataSchema {
	return f.db.readAllLCKTeam(f.ctx)
}
