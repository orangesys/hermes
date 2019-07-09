package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"

	"google.golang.org/api/option"

	"github.com/orangesys/hermes/pkg/billing"
	"github.com/orangesys/hermes/pkg/payments"	
)
func registerBatch() {
	jsonPath := os.Getenv("FIREBASE_JSON_PATH")
	// userID := "YZ3KuBygNIOhVvSjvxjl"
	
	opt := option.WithCredentialsFile(jsonPath)
	ctx := context.Background()
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		// return err
		log.Fatalln(err)
	}
	
	client, err := app.Firestore(ctx)
	if err != nil {
		// return err
		log.Fatalln(err)
	}
	defer client.Close()
	
	// snapIter := client.Collection("users").Where("state", "==", true).Snapshots(ctx)
	// defer snapIter.Stop()
	// for {
	// 	snap, err := snapIter.Next()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	docs, err := snap.Documents.GetAll()
	// 	fmt.Printf("data size: %d\n", snap.Size)
	// 	for i, data := range docs {
	// 		fmt.Printf("data %d, content: %+v\n", i, data.Data())
	// 	}
	// 	fmt.Println()
	// }
	server := "http://127.0.0.1:9090"
	sumNodes := billing.CountNodesFromQuerier(server)
	fmt.Println(sumNodes)
	
	iter := client.Collection("users").Where("state", "==", true).Documents(ctx)
	// batchlist := make([]interface{}, 0)
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
		// for k := range d {
		// 	fmt.Println(k, d[k])
		// 	fmt.Println("---")
		// }
	}
}

