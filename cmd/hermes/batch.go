package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"

	"github.com/orangesys/hermes/pkg/billing"
	"github.com/orangesys/hermes/pkg/payments"
)

func registerBatch() {

	// export GOOGLE_APPLICATION_CREDENTIALS=<service-account.json>
	ctx := context.Background()
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		// return err
		log.Fatalln(err)
	}
	defer client.Close()

	server := "http://127.0.0.1:9090"
	sumNodes := billing.CountNodesFromQuerier(server)
	fmt.Println(sumNodes)

	iter := client.Collection("users").Where("state", "==", true).Documents(ctx)
	var batchlist []interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		d := doc.Data()["payments"]

		if d != nil {
			// fmt.Println(doc.Ref.ID)
			// fmt.Println(doc.Data())
			batchlist = append(batchlist, d)
		}
	}
	for _, data := range batchlist {
		d := data.(map[string]interface{})
		fmt.Println(d["customerID"], d["subscriptionID"])
		q := int64(sumNodes)
		customerid := d["customerID"].(string)
		subscriptionid := d["subscriptionID"].(string)

		if err := payments.AddUsageRecord(subscriptionid, customerid, q); err != nil {
			fmt.Printf("cat not create %d usage record with %s", q, customerid)
		} else {
			fmt.Printf("create %d unit with %s", q, customerid)
		}
	}
}
