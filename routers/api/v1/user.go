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
	if err := db.UpdateUserState(ctx, firestoreClient, u.Email, userdata, payments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "signup successed"})
}
