package main

import (
	// "github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// app    = kingpin.New("hermes", "hermes command for orangesys")
	server = kingpin.Command("server", "hermes server for orangesys")

	batch = kingpin.Command("batch", "batch is create usage record to stripe")
)

func main() {
	switch kingpin.Parse() {
	case server.FullCommand():
		registerServer()
	case batch.FullCommand():
		registerBatch()
	}
	// app := kingpin.New(filepath.Base(os.Args[0]), "Hermes for Orangesys")
	// app.Version(version.Print("hermes"))

	// registerServer(app, "server")
	// if _, err := app.Parse(os.Args[1:]); err != nil {
	// 	fmt.Fprintln(os.Stderr, errors.Wrapf(err, "Error parsing commandline arguments"))
	// 	app.Usage(os.Args[1:])
	// 	os.Exit(2)
	// }
	// router := routers.InitRouter()

	// s := &http.Server{
	// 	Addr:    fmt.Sprintf(":%d", 8080),
	// 	Handler: router,
	// }
	// if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 	log.Fatalf("listen: %s\n", err)
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
