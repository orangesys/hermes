package db

import (
	"time"

	"golang.org/x/net/context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/iterator"
)

type Database struct {
	client *firestore.Client
	ctx    context.Context
}
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

func AddPaymentHistory(ctx context.Context, client *firestore.Client) error {
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

func GetBatchPaymentsList(ctx context.Context, client *firestore.Client) ([]map[string]interface{}, error) {
	iter := client.Collection(defaultCollection).Where("state", "==", true).Documents(ctx)
	// var batchlist []Payments
	batchlist := []map[string]interface{}{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			return batchlist, nil
		}
		if err != nil {
			return nil, err
		}

		payiter := client.Collection(defaultCollection).Doc(doc.Ref.ID).Collection(defaultSubCollection).Where("state", "==", true).Documents(ctx)
		for {
			doc, err := payiter.Next()
			if err == iterator.Done {
				return batchlist, nil
			}
			if err != nil {
				return nil, err
			}
			// fmt.Println(doc.Data())
			batchlist = append(batchlist, doc.Data())
		}
	}
}

func AddPaymentsCollection(ctx context.Context, client *firestore.Client, email string, payments *Payments) error {
	// defer client.Close()
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

func UpdateUserState(ctx context.Context, client *firestore.Client, email string, data map[string]interface{}) error {
	userID, err := getUserRefIdWithEmail(ctx, client, email)
	if err != nil {
		return err
	}
	_, err = client.Collection("users").Doc(userID).Set(ctx, data, firestore.MergeAll)
	if err != nil {
		return err
	}
	return nil
}

func getUserRefIdWithEmail(ctx context.Context, client *firestore.Client, email string) (string, error) {
	// defer client.Close()
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
