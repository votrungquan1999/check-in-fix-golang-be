package main

import (
	"checkinfix.com/controllers"
	"checkinfix.com/middlewares"
	"checkinfix.com/setup"
	"github.com/gin-gonic/gin"
)

func main() {
	setup.LoadConfig()
	r := gin.Default()

	r.Use(
		//cors.New(cors.Config{
		//	//AllowOrigins:     []string{"*"},
		//	AllowAllOrigins: true,
		//	//AllowOriginFunc: func(origin string) bool {
		//	//	return true
		//	//},
		//	AllowMethods: []string{"PUT", "POST", "GET", "PATCH", "OPTIONS"},
		//	AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host",
		//		"Token", "Authorization"},
		//	AllowCredentials: true,
		//	//ExposeHeaders:    []string{"access_token", "content-type"},
		//	MaxAge: 12 * time.Hour,
		//	//AllowWildcard:    true,
		//}),
		middlewares.HandleCORS(),
		middlewares.WithCommonError(),
	)

	controllers.SetupRoutes(r)

	setup.StartFirebase()

	_ = r.Run()
}
