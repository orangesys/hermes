package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/orangesys/hermes/pkg/billing"
	"github.com/orangesys/hermes/pkg/db"
	"github.com/orangesys/hermes/pkg/payments"
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

	server := "http://127.0.0.1:9090"
	sumNodes := billing.CountNodesFromQuerier(server)
	// fmt.Println(sumNodes)
	batchlist, err := db.GetBatchPaymentsList(ctx, firestoreClient)
	if err != nil {
		fmt.Printf("can not cat batch payments list: %v\n", err)
	}

	for payref, data := range batchlist {
		// fmt.Println(payref)
		// fmt.Println(data)
		d := data.(map[string]interface{})

		// fmt.Println(d["customerID"], d["subscriptionID"])
		q := int64(sumNodes)
		customerid := d["customerID"].(string)
		subscriptionid := d["subscriptionID"].(string)
		if err := payments.AddUsageRecord(subscriptionid, customerid, q); err != nil {
			fmt.Println(err)
			// fmt.Printf("cat not create %d usage record with %s", q, customerid)
		} else {
			if err := db.AddPaymentsHistory(ctx, firestoreClient, payref, q); err != nil {
				fmt.Printf("cat not add payments history to firestore: %v\n", err)
			}
			fmt.Printf("create %d unit with %s", q, customerid)
		}
	}
}
