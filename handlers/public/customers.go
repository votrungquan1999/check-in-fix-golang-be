package public

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"net/http"
)

func GetCustomers(c *gin.Context) {
	phoneNumber := c.Param("phone_number")
	firestoreClient := setup.FirestoreClient

	ctx := context.Background()
	iter := firestoreClient.Collection(constants.FirestoreUserDoc).Where("phone_number", "==", phoneNumber).Documents(ctx)

	users := make([]models.Customers, 0)

	for {
		var user models.Customers
		_, err := utils.GetNextDoc(iter, &user)
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		users = append(users, user)
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"data": users,
	})
}


