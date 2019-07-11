package main

import (
	// "github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// app    = kingpin.New("hermes", "hermes command for orangesys")
	serve = kingpin.Command("serve", "hermes serve for orangesys")

	batch = kingpin.Command("batch", "batch is create usage record to stripe")
)

func main() {
	switch kingpin.Parse() {
	case serve.FullCommand():
		registerServe()
	case batch.FullCommand():
		registerBatch()
	}

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
