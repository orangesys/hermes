package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/orangesys/hermes/pkg/db"
	"github.com/orangesys/hermes/pkg/payments"
)

// User is firebase database
// type User struct {
// 	Email       string `json:"email" binding:"required,email"`
// 	PlanID      string `json:"planid" binding:"required"`
// 	CompanyName string `json:"companyname" binding:"required"`
// 	CardNumber  string `json:"cardnumber" binding:"required"`
// 	ExpMonth    string `json:"expmonth" binding:"required"`
// 	ExpYear     string `json:"expyear" binding:"required"`
// 	CVC         string `json:"cvc" binding:"required"`
// }

func CreateUser(c *gin.Context) {
	var u payments.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	initUser, err := payments.InitPayUser(&u)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	}
	// fmt.Println(initUser)
	// create customer with email, email is unique
	// cus, err := payments.CreateCustomer(u.CompanyName, u.Email)
	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"messages": err.Error(),
	// 	})
	// 	return
	// }
	// fmt.Println(cus.ID)
	// add card source to customer
	// cardNumber, expMonth, expYear, cvc, stripeCustomerID string
	// if _, err := payments.AddSource(u.CardNumber, u.ExpMonth, u.ExpYear, u.CVC, cus.ID); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"messages": err.Error(),
	// 	})
	// 	return
	// }

	//add subscription to customer with planID
	// subItemID, err := payments.Addsubscription(u.PlanID, cus.ID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"messages": err.Error(),
	// 	})
	// 	return
	// }
	var state bool = true
	payments := &db.Payments{
		PlanID:         u.PlanID,
		CustomerID:     initUser["cusID"],
		SubscriptionID: initUser["subItemID"],
		StartDate:      time.Now(),
		State:          state,
	}

	userdata := map[string]interface{}{
		"state": state,
	}
	firebaseApp, err := db.InitApp()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	}
	firestoreClient, err := db.InitFirestoreClient(firebaseApp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	}
	ctx := context.Background()
	if err := db.UpdateUserState(ctx, firestoreClient, u.Email, userdata); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	} else {
		if err := db.AddPaymentsCollection(ctx, firestoreClient, u.Email, payments); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"messages": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "signup successed"})
		}
	}
}
