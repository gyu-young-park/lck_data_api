package firebaseapi

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

type fireStoreClient struct {
	client *firestore.Client
}

type FireStoreDataSchema map[string]interface{}

func newFireStoreClient(ctx context.Context, app *firebase.App) *fireStoreClient {
	fireStoreClient := &fireStoreClient{}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fireStoreClient.client = client
	return fireStoreClient
}

func (f *fireStoreClient) fireStoreClose() {
	defer f.fireStoreClose()
	fmt.Println("firestore instance closed")
}

func (f *fireStoreClient) readMatchTeamWithQueryOption(ctx context.Context, opt *ReadMatchQueryOption) []FireStoreDataSchema {
	var ret []FireStoreDataSchema
	iter := f.client.Collection(LCK_MATCH_VIDEO_COLLECTION).
		OrderBy("Date", firestore.Desc).
		Where("Season", "==", opt.Season).
		Limit(2).
		Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return nil
		}
		ret = append(ret, doc.Data())
	}
	fmt.Println(ret)
	return ret
}
