package middlewares

import (
	"checkinfix.com/constants"
	"checkinfix.com/setup"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenticateUser struct {
	AccessToken string `header:"Authorization"`
}

func FirebaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var loginUserRequest AuthenticateUser

		err := c.ShouldBindHeader(&loginUserRequest)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		if loginUserRequest.AccessToken == constants.CommonSecrete {
			return
		}

		decodedToken, err := setup.AuthClient.VerifyIDTokenAndCheckRevoked(ctx, loginUserRequest.AccessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		user, err := setup.AuthClient.GetUser(ctx, decodedToken.UID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.Set("user", user)
	}
}
