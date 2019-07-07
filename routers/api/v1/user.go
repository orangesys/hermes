package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User is firebase database
type User struct {
	Email       string `json:"email" binding:"required"`
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
	// c.JSON(http.StatusOK, gin.H{"status": "signup successed"})
	c.JSON(http.StatusOK, gin.H{
		"email":       u.Email,
		"companyName": u.CompanyName,
		"planID":      u.PlanID,
		"cardNumber":  u.CardNumber,
		"expMonth":    u.ExpMonth,
		"expYear":     u.ExpYear,
		"cvc":         u.CVC,
	})
}
