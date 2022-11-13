package firebaseapi

import (
	"context"
	"fmt"
	"log"
	"strconv"

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

func (f *fireStoreClient) where(ref *firestore.Query, field, condition, match string) *firestore.Query {
	query := ref.Where(field, condition, match)
	return &query
}

func (f *fireStoreClient) orderBy(ref *firestore.Query, field string, dir firestore.Direction) *firestore.Query {
	query := ref.OrderBy(field, dir)
	return &query
}

func (f *fireStoreClient) startAt(ref *firestore.Query, start int64) *firestore.Query {
	query := ref.StartAfter(start)
	return &query
}

func (f *fireStoreClient) limitQuery(ref *firestore.Query, limit int64) *firestore.Query {
	query := ref.Limit(int(limit))
	return &query
}

func (f *fireStoreClient) readMatchTeamWithQueryOption(ctx context.Context, opt *ReadMatchQueryOption) []FireStoreDataSchema {
	var ret []FireStoreDataSchema
	query := &f.client.Collection(LCK_MATCH_COLLECTION).Query
	if opt.Season != "" {
		query = f.where(query, "Season", "==", opt.Season)
	}

	if opt.Team != "" {
		if opt.Result == "W" {
			query = f.where(query, "WinTeam", "==", opt.Team)
		} else if opt.Result == "L" {
			query = f.where(query, "LoseTeam", "==", opt.Team)
		}
	}

	if opt.SortOpt != "desc" {
		query = f.orderBy(query, "PublishedAt", firestore.Desc)
	} else {
		query = f.orderBy(query, "PublishedAt", firestore.Asc)
	}

	if opt.PublishedAt != "" {
		publishedAt, err := strconv.ParseInt(opt.PublishedAt, 10, 64)
		if err != nil {
			publishedAt = 20
			fmt.Printf("Error: limit is not valid:%s!\n", err)
		} else {
			query = f.startAt(query, publishedAt)
		}
	}

	if opt.Limit != "" {
		limit, err := strconv.ParseInt(opt.Limit, 10, 64)
		if err != nil {
			limit = 20
			fmt.Printf("Error: limit is not valid:%s!\n", err)
		}
		query = f.limitQuery(query, limit)
	}
	iter := query.Documents(ctx)

	// if opt.Limit
	// iter := f.client.Collection(LCK_MATCH_COLLECTION).
	// 	Where("Season", "==", opt.Season).
	// 	Where("WinTeam", "==").
	// 	OrderBy("PublishedAt", firestore.Asc).
	// 	StartAt("").
	// 	Limit(10).
	// 	Documents(ctx)
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
