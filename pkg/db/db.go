package db

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/iterator"
)

type PaymentsHistory struct {
	Date int64 `firestore:"date,omitempty"`
}

type Payments struct {
	PlanID         string    `firestore:"planID,omitempty"`
	CustomerID     string    `firestore:"customerID,omitempty"`
	SubscriptionID string    `firestore:"subscriptionID,omitempty"`
	StartDate      time.Time `firestore:"startDate,omitempty"`
	State          bool      `firestore:"state,omitempty"`
}

type UserData struct {
	CompanyName      string `firestore:"companyName,omitempty"`
	Email            string `firestore:"email,omitempty"`
	PrometheusLables string `firestore:"prometheusLables,omitempty"`
	SubDomain        string `firestore:"subDomain,omitempty"`
	TelegrafToken    string `firestore:"telegrafToken,omitempty"`
	State            bool   `firestore:"state,omitempty"`
}

var defaultCollection = "users"
var defaultSubCollection = "payments"

// func AddPaymentsHistory(ctx context.Context, client *firestore.Client, payref string, data map[string]interface{}) error {
func AddPaymentsHistory(ctx context.Context, client *firestore.Client, payref string, sumnodes int64) error {
	p := strings.Split(payref, "-")
	year, month, day := time.Now().Date()
	payhistorydate := fmt.Sprintf("%d%d%d", year, month, day)
	data := map[string]interface{}{
		"paymentshistory": map[string]interface{}{
			payhistorydate: sumnodes,
		},
	}
	if _, err := client.Collection(defaultCollection).Doc(p[0]).Collection(defaultSubCollection).Doc(p[1]).Set(ctx, data, firestore.MergeAll); err != nil {
		return err
	}
	return nil
}

func GetBatchPaymentsList(ctx context.Context, client *firestore.Client) (map[string]interface{}, error) {
	iter := client.Collection(defaultCollection).Where("state", "==", true).Documents(ctx)
	// var batchlist []Payments
	// batchlist := []map[string]interface{}{}
	batchlist := make(map[string]interface{})
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			return batchlist, nil
		}
		if err != nil {
			return nil, err
		}
		colPath := fmt.Sprintf("%s/%s/%s", defaultCollection, doc.Ref.ID, defaultSubCollection)
		payiter := client.Collection(colPath).Where("state", "==", true).Documents(ctx)
		// payiter := client.Collection(defaultCollection).Doc(doc.Ref.ID).Collection(defaultSubCollection).Where("state", "==", true).Documents(ctx)
		for {
			paydoc, err := payiter.Next()
			if err == iterator.Done {
				return batchlist, nil
			}
			if err != nil {
				return nil, err
			}
			paycolPath := fmt.Sprintf("%s-%s", doc.Ref.ID, paydoc.Ref.ID)
			batchlist[paycolPath] = paydoc.Data()
		}
	}
}

func AddPaymentsCollection(ctx context.Context, client *firestore.Client, email string, payments *Payments) error {
	userID, err := getUserRefIdWithEmail(ctx, client, email)
	if err != nil {
		return err
	}
	_, _, err = client.Collection(defaultCollection).Doc(userID).Collection(defaultSubCollection).Add(ctx, payments)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserState(ctx context.Context, client *firestore.Client, email string, data map[string]interface{}, payments *Payments) error {
	userID, err := getUserRefIdWithEmail(ctx, client, email)
	if err != nil {
		return err
	}
	_, err = client.Collection("users").Doc(userID).Set(ctx, data, firestore.MergeAll)
	if err != nil {
		return err
	}
	if payments != nil {
		if err := AddPaymentsCollection(ctx, client, email, payments); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func getUserRefIdWithEmail(ctx context.Context, client *firestore.Client, email string) (string, error) {
	iter := client.Collection(defaultCollection).Where("email", "==", email).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			return "", err
		}
		if err != nil {
			return "", err
		}
		return doc.Ref.ID, nil
	}
}

func InitApp() (*firebase.App, error) {
	// Get a Firebase app
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func InitFirestoreClient(app *firebase.App) (*firestore.Client, error) {
	// Get a Firestore client.
	c, err := app.Firestore(context.Background())
	if err != nil {
		return nil, err
	}
	return c, nil
}
