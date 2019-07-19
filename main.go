package main

import (
	// "fmt"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
)

type PaymentsHistory struct {
	Date int64 `firestore:"date,omitempty"`
}

type Payments struct {
	PlanID          string    `firestore:"planID,omitempty"`
	CustomerID      string    `firestore:"customerID,omitempty"`
	SubscriptionID  string    `firestore:"subscriptionID,omitempty"`
	StartDate       time.Time `firestore:"startDate,omitempty"`
	State           bool      `firestore:"state,omitempty"`
	paymentshistory PaymentsHistory
}

func main() {
	ctx := context.Background()
	app, err := firebase.NewApp(context.Background(), nil)
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
	// server := "http://127.0.0.1:9090"
	// sumNodes := billing.CountNodesFromQuerier(server)
	// fmt.Println(sumNodes)
	// var sumNodes int64 = 146
	// email := "hogehoge3@example.com"
	// customerID := "cus_FPmKc8HFnpQM5j"
	// q := int64(189)
	// paydata := map[string]interface{}{
	// 	"paymentshistory": map[string]interface{}{
	// 		"20190713": q,
	// 	},
	// }
	// payref := "users/6Qn2ZFo4jnyY8l2JK5rC/payments/PbEV8pIJ7qb8M4iLBOlu"
	// a := strings.Split(payref, "/")
	// fmt.Println(a)
	var defaultCollection = "users"
	// var defaultSubCollection = "payments"
	// // batchlist := []map[string]interface{}{}
	// batchlist := make(map[string]interface{})
	// _, err = client.Collection("users").Doc("6Qn2ZFo4jnyY8l2JK5rC").Collection("payments").Doc("PbEV8pIJ7qb8M4iLBOlu").Set(ctx, paydata, firestore.MergeAll)
	dsnap, _ := client.Collection(defaultCollection).Doc("WJ1oSC5PxHdF9rHZHvXL").Collection("payments").Doc("RKnvdEe4DOOX38AwAYhy").Get(ctx)
	// query := users.Where("paymentshistory", "array-contains", "2019718").Documents(ctx)
	// fmt.Println(query)
	m := dsnap.Data()["paymentshistory"].(map[string]interface{})["2019718"]
	// fmt.Printf("Document data: %#v\n", m["paymentshistory"])
	// n := m["paymentshistory"].(map[string]interface{})
	fmt.Println(m)

	// for {
	// 	doc, err := iters.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(doc.Data())
	// 	// colPath := fmt.Sprintf("%s/%s/%s", defaultCollection, doc.Ref.ID, defaultSubCollection)
	// 	// // payiter := client.Collection(colPath).Where("state", "==", true).Documents(ctx)
	// 	// payiter := client.Collection(colPath).Where("2019718").Documents(ctx)
	// 	// for {
	// 	// 	paydoc, err := payiter.Next()
	// 	// 	if err == iterator.Done {
	// 	// 		break
	// 	// 	}
	// 	// 	if err != nil {
	// 	// 		fmt.Println(err)
	// 	// 	}
	// 	// 	// batchlist = append(batchlist, paydoc.Data())
	// 	// 	paycolPath := fmt.Sprintf("%s-%s", doc.Ref.ID, paydoc.Ref.ID)
	// 	// 	batchlist[paycolPath] = paydoc.Data()
	// 	// }
	// }

	// for i, b := range batchlist {
	// 	fmt.Println(i, b)
	// }

	// iter := client.Collection("users").Doc("6Qn2ZFo4jnyY8l2JK5rC").Collection("payments").Where("customerID", "==", customerID).Documents(ctx)
	// batchlist := make([]interface{}, 0)
	// var batchlist []interface{}
	// for {
	// 	doc, err := iter.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(doc.Data())
	// 	// d := doc.Data()["payments"]
	// 	// fmt.Println(doc.Ref.ID)
	// 	// if d != nil {
	// 	// 	// fmt.Println(doc.Ref.ID)
	// 	// 	// fmt.Println(doc.Data())
	// 	// 	batchlist = append(batchlist, d)
	// 	// }
	// }
	// for _, data := range batchlist {
	// 	d := data.(map[string]interface{})
	// 	fmt.Println(d["customerID"], d["subscriptionID"])
	// 	q := int64(sumNodes)
	// 	customerid := d["customerID"].(string)
	// 	subscriptionid := d["subscriptionID"].(string)

	// 	if err := payments.AddUsageRecord(subscriptionid, customerid, q); err != nil {
	// 		fmt.Printf("cat not create %d usage record with %s", q, customerid)
	// 	} else {
	// 		fmt.Printf("create %d unit with %s", q, customerid)
	// 	}
	// 	// for k := range d {
	// 	// 	fmt.Println(k, d[k])
	// 	// 	fmt.Println("---")
	// 	// }
	// }
	// for _, data := range batchlist {
	// 	for k, v := range data.(map[string]interface{}) {
	// 		fmt.Println(k, v)
	// 	}
	// }

	// reference
	// iter := client.Collection("users").Documents(ctx)
	// for {
	// 	_, err := iter.Next()
	// 	if err == iterator.Done {
	// 		// return nil
	// 		break
	// 	}
	// 	if err != nil {
	// 		// return err
	// 		log.Fatalf("Failed to iterate: %v", err)
	// 	}
	// 	// fmt.Println(doc.Data())
	// }
	// go func() {
	// 	// service Conections
	// 	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// quit := make(chan os.Signal)
	// // kill (no param) default send syscall.SIGTERM
	// // kill -2 is syscall.SIGINT
	// // kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("ShutDown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := s.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// select {
	// case <-ctx.Done():
	// 	log.Println("timeout of 5 seconds.")
	// }
	// log.Println("Server exiting")

	//AddUsageRecord
	// var q int64 = 100
	// q := int64(nodes)
	// customerid := "cus_FM1aNamxCy9S2S"
	// subscriptionid := "si_FM6vfuQW7M6R7u"

	// if err := payments.AddUsageRecord(subscriptionid, customerid, q); err != nil {
	// 	fmt.Printf("cat not create %d usage record with %s", q, customerid)
	// } else {
	// 	fmt.Printf("create %d unit with %s", q, customerid)
	// }

	//ListUsageREcord
	// stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"

	// params := &stripe.UsageRecordSummaryListParams{
	// 	SubscriptionItem: stripe.String(subscriptionid),
	// }
	// // params.Filters.AddFilter("limit", "", "3")
	// // params.Filters.AddFilter("ending_before", "", "1562284800")
	// i := usagerecordsummary.List(params)
	// for i.Next() {
	// 	u := i.UsageRecordSummary()
	// 	fmt.Println(u)
	// 	fmt.Println(u.Period)
	// }

	// Create prometheus plan
	// params := &stripe.PlanParams{
	// 	Amount:   stripe.Int64(10),
	// 	Interval: stripe.String("month"),
	// 	Product: &stripe.PlanProductParams{
	// 		Name: stripe.String("prometheus unit"),
	// 	},
	// 	ID:        stripe.String("promeunit"),
	// 	Currency:  stripe.String(string(stripe.CurrencyJPY)),
	// 	UsageType: stripe.String("metered"),
	// }
	// p, err := plan.New(params)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(p.ID)
	// }

	// list all plans
	// params := &stripe.PlanListParams{}
	// params.Filters.AddFilter("limit", "", "3")
	// i := plan.List(params)
	// for i.Next() {
	// 	p := i.Plan()

	// 	fmt.Println(p.ID)
	// }

}
