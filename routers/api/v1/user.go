package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orangesys/hermes/pkg/db"
	"github.com/orangesys/hermes/pkg/payments"
)

// User is firebase database
type User struct {
	Email       string `json:"email" binding:"required,email"`
	UserID      string `json:"userid" binding:"required"`
	PlanID      string `json:"planid" binding:"required"`
	CompanyName string `json:"companyname" binding:"required"`
	CardNumber  string `json:"cardnumber" binding:"required"`
	ExpMonth    string `json:"expmonth" binding:"required"`
	ExpYear     string `json:"expyear" binding:"required"`
	CVC         string `json:"cvc" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// create customer with email, email is unique
	cus, err := payments.CreateCustomer(u.CompanyName, u.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	}
	fmt.Println(cus.ID)
	// add card source to customer
	// cardNumber, expMonth, expYear, cvc, stripeCustomerID string
	if _, err := payments.AddSource(u.CardNumber, u.ExpMonth, u.ExpYear, u.CVC, cus.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	}

	//add subscription to customer with planID
	subItemID, err := payments.Addsubscription(u.PlanID, cus.ID)
	fmt.Println(subItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	}

	paymentData := map[string]interface{}{
		"payments": map[string]interface{}{
			"customerID":     cus.ID,
			"subscriptionID": subItemID,
		},
	}

	if err := db.UpdateDB(u.UserID, paymentData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "signup successed"})
	}
}
