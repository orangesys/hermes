package main

import (
	// "fmt"

	"fmt"

	"github.com/orangesys/hermes/pkg/payments"
)

// // CreateCustomer with email , unique email
// func CreateCustomer(email string) (cust *stripe.Customer, err error) {
// 	stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"
// 	fmt.Println(stripe.Key)

// 	if err := customerIsExist(email); err != nil {
// 		return nil, err
// 	}
// 	newCustomerParams := &stripe.CustomerParams{
// 		Email: stripe.String(email),
// 	}

// 	if cust, err = customer.New(newCustomerParams); err != nil {
// 		return nil, err
// 	}
// 	return cust, err
// }

// // customerIsExist
// func customerIsExist(email string) error {
// 	stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"

// 	params := &stripe.CustomerListParams{}
// 	params.Filters.AddFilter("email", "", email)
// 	i := customer.List(params)

// 	if i.Next() {
// 		return fmt.Errorf("%s is exsiter.", email)
// 	}
// 	return nil
// }

// //AddSource is add card to customer
// func AddSource(cardNumber, expMonth, expYear, cvc, stripeCustomerID string) (*stripe.PaymentSource, error) {
// 	stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"

// 	tokenParams := &stripe.TokenParams{
// 		Card: &stripe.CardParams{
// 			Number:   stripe.String(cardNumber),
// 			ExpMonth: stripe.String(expMonth),
// 			ExpYear:  stripe.String(expYear),
// 			CVC:      stripe.String(cvc),
// 		},
// 	}
// 	t, err := token.New(tokenParams)
// 	if err != nil {
// 		return nil, err
// 	}

// 	customerSourceParams := &stripe.CustomerSourceParams{
// 		Customer: stripe.String(stripeCustomerID),
// 		Source: &stripe.SourceParams{
// 			Token: stripe.String(t.ID),
// 		},
// 	}

// 	return paymentsource.New(customerSourceParams)
// }

// //Addsubscription add subscription with customer by monthly
// func Addsubscription(planID, stripeCustomerID string) (string, error) {
// 	subParams := &stripe.SubscriptionParams{
// 		Customer: stripe.String(stripeCustomerID),
// 		Items: []*stripe.SubscriptionItemsParams{
// 			{
// 				Plan: stripe.String(planID),
// 			},
// 		},
// 	}
// 	s, err := sub.New(subParams)
// 	if err != nil {
// 		return "", err
// 	}
// 	return s.Items.Data[0].ID, nil
// }

// //UsageRecord create usage record daily by cusmtomer
// func DailyUsageRecord(subItemID, stripeCustomerID string, quantity int64) error {
// 	params := &stripe.UsageRecordParams{
// 		Quantity:         stripe.Int64(quantity),
// 		SubscriptionItem: stripe.String(subItemID),
// 		Timestamp:        stripe.Int64(time.Now().Unix()),
// 	}

// 	if _, err := usagerecord.New(params); err != nil {
// 		return err
// 	}
// 	return nil
// }

func main() {
	// stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"

	// Create UsageRecord
	// params := &stripe.UsageRecordParams{
	// 	Quantity:  stripe.Int64(20),
	// 	Timestamp: stripe.Int64(1564458167), // timestamp 2019 7/30
	// 	// SubscriptionItem: stripe.String(s.Items.Data[0].ID),
	// 	SubscriptionItem: stripe.String("si_FLyG6ZyMl6CtH5"),
	// }

	// record, err := usagerecord.New(params)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(record)
	// }

	// create customer with email, email is unique
	if cus, err := payments.CreateCustomer("hogehoge1@example.com"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cus.ID)
	}

	// cardNumber, expMonth, expYear, cvc, stripeCustomerID string
	// if _, err := AddSource("4242424242424242", "11", "23", "123", "cus_FM1aNamxCy9S2S"); err != nil {
	// 	fmt.Println(err)
	// }

	//addsubscription
	// if subItemID, err := Addsubscription("promeunit", "cus_FM1aNamxCy9S2S"); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(subItemID)
	// }

	//DailyUsageRecord
	var q int64 = 100
	if err := payments.DailyUsageRecord("si_FM6vfuQW7M6R7u", "cus_FM1aNamxCy9S2S", q); err != nil {
		fmt.Printf("cat not create %d usage record with %s", q, "cus_FM1aNamxCy9S2S")
	} else {
		fmt.Printf("create %d unit with %s", q, "cus_FM1aNamxCy9S2S")
	}

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
