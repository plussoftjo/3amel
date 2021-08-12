// Package routes (Setup Routes Group)
package routes

import (
	"server/config"
	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Setup >>>
func Setup() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "authorization", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))
	// gin.SetMode(gin.ReleaseMode)
	r.Use(static.Serve("/public", static.LocalFile(config.ServerInfo.PublicPath+"public", true)))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success",
		})
	})

	// -------- Auth Groups ----------//

	// ~~~ Auth Group ~~~ //
	auth := r.Group("/auth")
	auth.POST("/login", controllers.LoginController)
	auth.POST("/register", controllers.RegisterController)
	auth.POST("/app/register", controllers.AppRegisterController)
	auth.POST("/app/login", controllers.AppLoginController)
	auth.GET("/app/auth", controllers.AuthAppUser)
	auth.POST("/app/changePassword", controllers.ChangePassword)
	auth.GET("/auth", controllers.Auth)
	auth.GET("/users/index", controllers.UsersListIndex)
	auth.GET("/users/delete/:id", controllers.DeleteUser)
	auth.POST("/update", controllers.UpdateUser)
	auth.POST("/app/update", controllers.AppUpdateUser)
	auth.POST("/checkHasPhone", controllers.CheckIfHasPhone)
	auth.POST("/resetPassword", controllers.ResetPassword)

	// --------- Basics ------- //
	basics := r.Group("/basics")

	// UploadImage => For All
	basics.POST("/upload_image/:imageType", controllers.UpdateImage)

	// --------- User Controller ----------------- //
	user := r.Group("/users")
	// ~~~ User Roles ~~~ //
	user.POST("/roles/store", controllers.StoreUserRoles)
	user.POST("/roles/update", controllers.UpdateUserRole)
	user.GET("/roles/index", controllers.IndexUserRoles)
	user.GET("/roles/delete/:id", controllers.DeleteUserRole)
	// --------------- Employ Controller ----------- //
	user.POST("/employee/store", controllers.StoreEmployee)
	user.GET("/employee/index", controllers.IndexEmployee)
	user.GET("/employee/delete/:id", controllers.DeleteEmployee)
	user.POST("/employee/update", controllers.UpdateEmployee)

	// -------- Services ---------- //
	services := r.Group("/services")
	services.POST("/store", controllers.StoreService)
	services.GET("/index", controllers.IndexServices)
	services.GET("/destroy/:id", controllers.DestroyService)
	services.POST("/update", controllers.UpdateService)

	// Sub Services
	subServices := r.Group("/subServices")
	subServices.POST("/store", controllers.StoreSubService)
	subServices.GET("/index", controllers.IndexSubServices)
	subServices.GET("/destroy/:id", controllers.DestroySubService)
	subServices.GET("/indexWithServiceID/:id", controllers.IndexSubServicesWithServiceID)
	subServices.POST("/update", controllers.UpdateSubService)

	// Services Options
	servicesOptions := r.Group("/servicesOptions")
	servicesOptions.POST("/store", controllers.StoreServicesOptions)
	servicesOptions.GET("/index", controllers.IndexServicesOptions)
	servicesOptions.GET("/destroy/:id", controllers.DestroyServiceOptions)
	servicesOptions.POST("/update", controllers.UpdateServicesOptions)

	// Application
	application := r.Group("/app")
	application.GET("/main/index", controllers.IndexMain)

	// Orders
	orders := r.Group("orders")
	orders.POST("/store", controllers.StoreOrder)
	orders.GET("/finishOrderFromUser/:id", controllers.FinishOrderFromUser)
	orders.GET("/show", controllers.ShowOrder)

	r.Run(":8082")
}