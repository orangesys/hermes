package db

import (
	"os"

	"golang.org/x/net/context"

	"cloud.google.com/go/firestore"

	firebase "firebase.google.com/go"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func UpdateDB(userID string, data map[string]interface{}) error {

	jsonPath := os.Getenv("FIREBASE_JSON_PATH")

	opt := option.WithCredentialsFile(jsonPath)
	ctx := context.Background()
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
		// log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return err
		// log.Fatalln(err)
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(userID).Set(ctx, data, firestore.MergeAll)
	if err != nil {
		return err
		// log.Fatalf("Failed adding aturing: %v", err)
	}

	// reference
	iter := client.Collection("users").Documents(ctx)
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			return nil
			// break
		}
		if err != nil {
			return err
			// log.Fatalf("Failed to iterate: %v", err)
		}
		// fmt.Println(doc.Data())
	}
}
