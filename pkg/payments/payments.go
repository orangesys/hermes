package payments

import (
	"fmt"
	"time"

	"github.com/stripe/stripe-go"

	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/sub"
	"github.com/stripe/stripe-go/usagerecord"

	"github.com/stripe/stripe-go/paymentsource"
	"github.com/stripe/stripe-go/token"
)

// deleteCustomer if card is invalid
func deleteCustomer(stripeCustomerID string) error {
	stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"
	params := &stripe.CustomerParams{}
	_, err := customer.Del(stripeCustomerID, params)
	if err != nil {
		return err
	}
	return nil
}

// CreateCustomer with email , unique email
func CreateCustomer(companyname, email string) (cust *stripe.Customer, err error) {
	stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"

	if err := customerIsExist(email); err != nil {
		return nil, err
	}
	newCustomerParams := &stripe.CustomerParams{
		Name:  stripe.String(companyname),
		Email: stripe.String(email),
	}

	if cust, err = customer.New(newCustomerParams); err != nil {
		return nil, err
	}
	return cust, nil
}

// customerIsExist
func customerIsExist(email string) error {
	stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"

	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("email", "", email)
	i := customer.List(params)

	if i.Next() {
		return fmt.Errorf("%s is exsiter.", email)
	}
	return nil
}

//AddSource is add card to customer
func AddSource(cardNumber, expMonth, expYear, cvc, stripeCustomerID string) (*stripe.PaymentSource, error) {
	stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"

	tokenParams := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number:   stripe.String(cardNumber),
			ExpMonth: stripe.String(expMonth),
			ExpYear:  stripe.String(expYear),
			CVC:      stripe.String(cvc),
		},
	}
	t, err := token.New(tokenParams)
	if err != nil {
		if err := deleteCustomer(stripeCustomerID); err != nil {
			return nil, fmt.Errorf("can not delete invalid card user, please check stripe")
		}
		return nil, err
	}

	customerSourceParams := &stripe.CustomerSourceParams{
		Customer: stripe.String(stripeCustomerID),
		Source: &stripe.SourceParams{
			Token: stripe.String(t.ID),
		},
	}

	return paymentsource.New(customerSourceParams)
}

//Addsubscription add subscription with customer by monthly
func Addsubscription(planID, stripeCustomerID string) (string, error) {
	stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"
	subParams := &stripe.SubscriptionParams{
		Customer: stripe.String(stripeCustomerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan: stripe.String(planID),
			},
		},
	}
	s, err := sub.New(subParams)
	if err != nil {
		return "", err
	}
	return s.Items.Data[0].ID, nil
}

//AddUsageRecord create usage record daily by cusmtomer
// PromQL:	count(node_boot_time_seconds)
// 			count(node_boot_time_seconds)[24h:1h]
func AddUsageRecord(subItemID, stripeCustomerID string, quantity int64) error {
	stripe.Key = "sk_test_ljCYC27PV9LBxE1XYAA813jq"
	params := &stripe.UsageRecordParams{
		Quantity:         stripe.Int64(quantity),
		SubscriptionItem: stripe.String(subItemID),
		Timestamp:        stripe.Int64(time.Now().Unix()),
	}

	if _, err := usagerecord.New(params); err != nil {
		return err
	}
	return nil
}
