package middlewares

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type AuthenticateUser struct {
	AccessToken string `header:"Authorization"`
}

func FirebaseAuth(scopes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := validateUser(c)
		if err != nil {
			HandleError(c, err)
			return
		}
		if user == nil {
			return
		}

		if err = validateEmployeeScopes(user, scopes); err != nil {
			HandleError(c, err)
			return
		}

		c.Set("user", user)
	}
}

func validateUser(c *gin.Context) (*auth.UserRecord, error) {
	ctx := context.Background()
	var loginUserRequest AuthenticateUser

	err := c.ShouldBindHeader(&loginUserRequest)
	if err != nil {
		return nil, utils.ErrorBadRequest.New(err.Error())
	}

	secrete, exists := os.LookupEnv("API_SECRETE")
	if exists && loginUserRequest.AccessToken == secrete {
		return nil, nil
	}

	decodedToken, err := setup.AuthClient.VerifyIDTokenAndCheckRevoked(ctx, loginUserRequest.AccessToken)
	if err != nil {
		return nil, utils.ErrorUnauthorized.New(err.Error())
	}

	user, err := setup.AuthClient.GetUser(ctx, decodedToken.UID)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return user, nil
}

func validateEmployeeScopes(user *auth.UserRecord, scopes []string) error {
	if len(scopes) == 0 {
		return nil
	}

	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	employeeIter := firestoreClient.Collection(constants.FirestoreEmployeeDoc).Where("user_id", "==",
		user.UID).Documents(ctx)

	var employee models.Employees
	id, err := utils.GetNextDoc(employeeIter, &employee)
	if err != nil {
		return utils.ErrorInternal.New(err.Error())
	}
	if id == "" {
		return utils.ErrorInternal.New("there is no employee associated with this user")
	}

	if employee.Scopes == nil || len(employee.Scopes) == 0 {
		return utils.ErrorForbidden.New(fmt.Sprintf("user need one of these scopes: %v", scopes))
	}

	for _, scope := range scopes {
		if utils.SliceContain(utils.InterfaceSlice(employee.Scopes), scope) {
			return nil
		}
	}

	return utils.ErrorForbidden.New(fmt.Sprintf("user need one of these scopes: %v", scopes))
}
