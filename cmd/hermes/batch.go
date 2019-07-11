package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/orangesys/hermes/pkg/billing"
	"github.com/orangesys/hermes/pkg/db"
	"github.com/orangesys/hermes/pkg/payments"
	// "github.com/orangesys/hermes/pkg/payments"
)

func registerBatch() {
	firebaseApp, err := db.InitApp()
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}
	firestoreClient, err := db.InitFirestoreClient(firebaseApp)
	if err != nil {
		log.Fatalf("error initializing firestore client: %v\n", err)
	}

	ctx := context.Background()
	// app, err := firebase.NewApp(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("error initializing app: %v\n", err)
	// }

	// client, err := app.Firestore(ctx)
	// if err != nil {
	// 	// return err
	// 	log.Fatalln(err)
	// }
	// defer client.Close()

	server := "http://127.0.0.1:9090"
	sumNodes := billing.CountNodesFromQuerier(server)
	fmt.Println(sumNodes)
	batchlist, err := db.GetBatchPaymentsList(ctx, firestoreClient)
	if err != nil {
		fmt.Printf("can not cat batch payments list: %v\n", err)
	}
	// iter := client.Collection("users").Where("state", "==", true).Documents(ctx)
	// var batchlist []interface{}
	// for {
	// 	doc, err := iter.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	d := doc.Data()["payments"]

	// 	if d != nil {
	// 		// fmt.Println(doc.Ref.ID)
	// 		// fmt.Println(doc.Data())
	// 		batchlist = append(batchlist, d)
	// 	}
	// }
	for _, data := range batchlist {
		fmt.Println(data)
		// d := data.(map[string]interface{})

		fmt.Println(data["customerID"], data["subscriptionID"])
		q := int64(sumNodes)
		customerid := data["customerID"].(string)
		subscriptionid := data["subscriptionID"].(string)

		if err := payments.AddUsageRecord(subscriptionid, customerid, q); err != nil {
			fmt.Printf("cat not create %d usage record with %s", q, customerid)
		} else {
			fmt.Printf("create %d unit with %s", q, customerid)
		}
	}
}
