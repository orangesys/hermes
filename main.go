package main

import (
	"fmt"
	"net/http"

	"github.com/orangesys/hermes/routers"
)

// "github.com/orangesys/hermes/pkg/db"

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
	}
	s.ListenAndServe()

	// currentTime := time.Now()
	// fmt.Println(currentTime)
	// start, end := oneDaysAgoTimestamp(currentTime)

	// // count prometheus server with usage record
	// // count(kube_pod_start_time{pod=~"prometheus-k8s.*"})

	// query := "count(node_boot_time_seconds)"
	// server := "http://127.0.0.1:9090"

	// client, err := promql.NewClient(server)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// resp, err := client.QueryRange(query, start, end)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// type valueEntry struct {
	// 	Metric map[string]string `json:"metric"`
	// 	Value  float64           `json:"value`
	// }

	// type timeEntry struct {
	// 	Time   int64         `jsno:"time"`
	// 	Values []*valueEntry `json:"values"`
	// }

	// entryByTime := map[int64]*timeEntry{}
	// var nodes float64
	// // Save count node number to firestore
	// // Show billing with app.orangesys.io dashboard
	// for _, r := range resp.Data.Result {
	// 	for _, v := range r.Values {
	// 		t := v.Time()
	// 		u := t.Unix()
	// 		e, ok := entryByTime[u]

	// 		if !ok {
	// 			e = &timeEntry{
	// 				Time:   u,
	// 				Values: []*valueEntry{},
	// 			}
	// 			entryByTime[u] = e
	// 		}
	// 		val, err := v.Value()
	// 		nodes = nodes + val
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		e.Values = append(e.Values, &valueEntry{
	// 			Metric: r.Metric,
	// 			Value:  val,
	// 		})
	// 	}
	// }
	// fmt.Printf("count node by 24H is %v\n", nodes)

	// s := make([]*timeEntry, len(entryByTime))
	// i := 0
	// for _, e := range entryByTime {
	// 	s[i] = e
	// 	i++
	// }

	// b, err := json.Marshal(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(b))
	// fmt.Println(res.Body)

	// data := map[string]interface{}{
	// 	"email":       "gavin.zhou@gmail.com",
	// 	"companyName": "Orangesys Inc.",
	// 	"payments": map[string]interface{}{
	// 		"customerID":     "cus_FM1aNamxCy9S2S",
	// 		"subscriptionID": "sub_FM6vmkb99rb7K3",
	// 	},
	// 	"planID":          "promunit",
	// 	"prometheusLable": "orangesys",
	// 	"telegrafToken":   "orangesys-token",
	// 	"subDomain":       "demo",
	// }

	// if err := db.UpdateDB("YZ3KuBygNIOhVvSjvxjl", data); err != nil {
	// 	log.Fatalf("Failed adding aturing: %v", err)
	// }

	// create customer with email, email is unique
	// if cus, err := payments.CreateCustomer("hogehoge1@example.com"); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(cus.ID)
	// }

	// cardNumber, expMonth, expYear, cvc, stripeCustomerID string
	// if _, err := payments.AddSource("4242424242424242", "11", "23", "123", "cus_FM1aNamxCy9S2S"); err != nil {
	// 	fmt.Println(err)
	// }

	//addsubscription
	// if subItemID, err := payments.Addsubscription("promeunit", "cus_FM1aNamxCy9S2S"); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(subItemID)
	// }

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

	// TODO: Customer IDなどのカスタマー情報をDBに保存する

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
