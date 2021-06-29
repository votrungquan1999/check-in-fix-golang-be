package middlewares

import (
	"checkinfix.com/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonError struct {
	Message string `json:"message"`
}

func WithCommonError() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		err := context.Errors.Last()
		if err == nil {
			return
		}

		HandleError(context, err.Err)
	}
}

func HandleError(c *gin.Context, err error) {
	switch t := err.(type) {
	case *utils.CustomError:
		currentErr := err.(*utils.CustomError)

		newError := CommonError{
			Message: currentErr.Message,
		}
		c.AbortWithStatusJSON(currentErr.Code, gin.H{
			"error": newError,
		})
		return
	default:
		newError := CommonError{
			Message: fmt.Sprintf("error type is invalid, type %T", t),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": newError,
		})
		return
	}
}
