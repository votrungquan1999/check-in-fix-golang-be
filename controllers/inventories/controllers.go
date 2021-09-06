package inventories

import (
	"checkinfix.com/handlers/inventories"
	"checkinfix.com/requests"
	"checkinfix.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesGroup(rg *gin.RouterGroup) {
	router := rg.Group("/products")
	{
		router.POST("", createProduct)
		router.GET("", getListProduct)
		router.PATCH("/:product_id", updateProduct)
		router.GET("/:product_id", getProductDetail)
		router.DELETE("/:product_id", deleteProduct)
	}
}

func createProduct(c *gin.Context) {
	var reqBody requests.CreateProductRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	newProduct, err := inventories.CreateNewProduct(reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": newProduct,
	})
}

func updateProduct(c *gin.Context) {
	productID := c.Param("product_id")
	var reqBody requests.UpdateProductRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	updatedProduct, err := inventories.UpdateProduct(reqBody, productID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": updatedProduct,
	})
}

func getProductDetail(c *gin.Context) {
	productID := c.Param("product_id")

	product, err := inventories.GetProductDetail(productID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func getListProduct(c *gin.Context) {
	subscriberID := c.Query("subscriber_id")
	if subscriberID == "" {
		_ = c.Error(utils.ErrorBadRequest.New("subscriber_id query missing"))
		return
	}

	products, err := inventories.GetProductBySubscriber(subscriberID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func deleteProduct(c *gin.Context) {
	productID := c.Param("product_id")

	if err := inventories.DeleteProduct(productID); err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
