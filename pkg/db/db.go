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

type FirestoreClientImpl struct {
	context.Context
	*firestore.Client
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

// checkPaymentsHistory
// func PaymentsHistoryIsExist(ctx context.Context, client *firestore.Client, payref string, sumnodes int64) bool {
func (f *FirestoreClientImpl) PaymentsHistoryIsExist(payref string, sumnodes int64) bool {
	p := strings.Split(payref, "/")
	year, month, day := time.Now().AddDate(0, 0, -1).Date()
	payhistorydate := fmt.Sprintf("%d%02d%02d", year, month, day)
	paymentshistory, _ := f.Collection(defaultCollection).Doc(p[0]).Collection(defaultSubCollection).Doc(p[1]).Get(f.Context)
	return paymentshistory.Data()["paymentshistory"].(map[string]interface{})[payhistorydate] == sumnodes
}

// AddPaymentsHistory(ctx context.Context, client *firestore.Client, payref string, data map[string]interface{}) error {
func (f *FirestoreClientImpl) AddPaymentsHistory(payref string, sumnodes int64) error {
	p := strings.Split(payref, "/")
	year, month, day := time.Now().AddDate(0, 0, -1).Date()
	payhistorydate := fmt.Sprintf("%d%02d%02d", year, month, day)
	data := map[string]interface{}{
		"paymentshistory": map[string]interface{}{
			payhistorydate: sumnodes,
		},
	}
	if _, err := f.Collection(defaultCollection).Doc(p[0]).Collection(defaultSubCollection).Doc(p[1]).Set(f.Context, data, firestore.MergeAll); err != nil {
		return err
	}
	return nil
}

func (f *FirestoreClientImpl) GetBatchPaymentsList() (map[string]interface{}, error) {
	iter := f.Collection(defaultCollection).Where("state", "==", true).Documents(f.Context)
	batchlist := make(map[string]interface{})
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		colPath := fmt.Sprintf("%s/%s/%s", defaultCollection, doc.Ref.ID, defaultSubCollection)
		payiter := f.Collection(colPath).Where("state", "==", true).Documents(f.Context)
		for {
			paydoc, err := payiter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			paycolPath := fmt.Sprintf("%s/%s", doc.Ref.ID, paydoc.Ref.ID)
			batchlist[paycolPath] = paydoc.Data()
		}
	}
	return batchlist, nil
}

func (f *FirestoreClientImpl) AddPaymentsCollection(email string, payments *Payments) error {
	userID, err := f.GetUserRefIdWithEmail(email)
	if err != nil {
		return err
	}
	_, _, err = f.Collection(defaultCollection).Doc(userID).Collection(defaultSubCollection).Add(f.Context, payments)
	if err != nil {
		return err
	}
	return nil
}

func (f *FirestoreClientImpl) UpdateUserState(email string, data map[string]interface{}, payments *Payments) error {
	userID, err := f.GetUserRefIdWithEmail(email)
	if err != nil {
		return err
	}
	_, err = f.Collection(defaultCollection).Doc(userID).Set(f.Context, data, firestore.MergeAll)
	if err != nil {
		return err
	}
	if payments != nil {
		if err := f.AddPaymentsCollection(email, payments); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (f *FirestoreClientImpl) GetUserRefIdWithEmail(email string) (string, error) {
	iter := f.Collection(defaultCollection).Where("email", "==", email).Documents(f.Context)
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
