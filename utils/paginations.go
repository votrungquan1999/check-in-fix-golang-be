package utils

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func WithPagination() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := setLimit(context)
		if err != nil {
			handlePaginationError(err, context)
			return
		}

		err = setCurrentPage(context)
		if err != nil {
			handlePaginationError(err, context)
			return
		}

		context.Next()
	}
}

func setLimit(c *gin.Context) error {
	limit := c.Query("limit")
	var intLimit int64

	if limit == "" {
		intLimit = 10
		c.Set("limit", int(intLimit))
		return nil
	}

	intLimit, err := strconv.ParseInt(limit, 10, 0)
	if err != nil {
		return err
	}

	c.Set("limit", int(intLimit))
	return nil
}

func setCurrentPage(c *gin.Context) error {
	currentPage := c.Query("current_page")
	var intCurrentPage int64

	if currentPage == "" {
		intCurrentPage = 1
		c.Set("current_page", int(intCurrentPage))
		return nil
	}

	intCurrentPage, err := strconv.ParseInt(currentPage, 10, 0)
	if err != nil {
		return err
	}

	c.Set("current_page", int(intCurrentPage))
	return nil
}

func handlePaginationError(err error, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}

func PaginateQuery(query firestore.Query, limit int, currentPage int, key string) (firestore.Query, error) {
	ctx := context.Background()

	if currentPage == 1 {
		return query, nil
	}

	prevData := query.Limit((currentPage - 1) * limit).Documents(ctx)
	docs, err := prevData.GetAll()
	if err != nil {
		return query, ErrorInternal.New(err.Error())
	}

	lastDoc := docs[len(docs)-1]

	paginatedQuery := query.StartAfter(lastDoc.Data()[key])

	return paginatedQuery, nil
}
