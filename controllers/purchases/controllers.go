package purchases

import (
	"checkinfix.com/handlers/purchases"
	"checkinfix.com/requests"
	"checkinfix.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesGroup(rg *gin.RouterGroup) {
	router := rg.Group("/purchases")
	{
		router.POST("", createPurchases)
		router.GET("", getListPurchases)
		//router.PATCH("/:product_id", updateProduct)
		//router.GET("/:product_id", getProductDetail)
		//router.DELETE("/:product_id", deleteProduct)
	}
}

func createPurchases(c *gin.Context) {
	var createPurchasePayload requests.CreatePurchaseRequest
	if err := c.ShouldBindJSON(&createPurchasePayload); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	newPurchase, err := purchases.CreatePurchase(createPurchasePayload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": newPurchase,
	})
}

func getListPurchases(c *gin.Context) {
	subscriberID := c.Query("subscriber_id")
	if subscriberID == "" {
		_ = c.Error(utils.ErrorBadRequest.New("subscriber_id query missing"))
		return
	}

	purchases, err := purchases.GetListPurchases(subscriberID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": purchases,
	})
}
