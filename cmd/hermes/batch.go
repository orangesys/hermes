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

	// ctx := context.Background()

	server := "http://127.0.0.1:9090"
	sumNodes := billing.CountNodesFromQuerier(server)

	fs := db.FirestoreClientImpl{
		context.Background(),
		firestoreClient,
	}
	paymentsbatchlist, err := fs.GetBatchPaymentsList()
	if err != nil {
		fmt.Printf("can not cat batch payments list: %v\n", err)
	}

	for payref, data := range paymentsbatchlist {
		d := data.(map[string]interface{})

		// fmt.Println(d["customerID"], d["subscriptionID"])
		q := int64(sumNodes)

		// if db.PaymentsHistoryIsExist(ctx, firestoreClient, payref, q) {
		if fs.PaymentsHistoryIsExist(payref, q) {
			fmt.Println("Had payment record in payment history")
			continue
		}
		customerid := d["customerID"].(string)
		subscriptionid := d["subscriptionID"].(string)
		if err := payments.AddUsageRecord(subscriptionid, customerid, q); err != nil {
			fmt.Printf("can not add %d nodes usage record to %s customerID : %v\n", q, customerid, err)
			continue
		}
		if err := fs.AddPaymentsHistory(payref, q); err != nil {
			fmt.Printf("cat not add payments history to firestore: %v\n", err)
			continue
		}
		fmt.Printf("create %d unit with %s\n", q, customerid)
	}
}
