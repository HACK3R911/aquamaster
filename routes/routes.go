package routes

import (
	"aquamaster/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/", controllers.GetHomeForm())

	incomingRoutes.GET("/profile", controllers.GetProfileForm())
	incomingRoutes.GET("/cart", controllers.GetCartForm())
	incomingRoutes.GET("/favorites", controllers.GetFavoritesForm())
	incomingRoutes.GET("/plumbing", controllers.GetPlumbingForm())
	incomingRoutes.GET("/interior", controllers.GetInteriorForm())
	incomingRoutes.GET("/accessories", controllers.GetAccessoriesForm())

	incomingRoutes.GET("/signup", controllers.GetSignUpForm())

	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
