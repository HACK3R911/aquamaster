package routes

import (
	"aquamaster/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/", controllers.GetHomeForm())

	//body
	incomingRoutes.GET("/profile", controllers.GetProfileForm())
	incomingRoutes.GET("/cart", controllers.GetCartForm())
	incomingRoutes.GET("/favorites", controllers.GetFavoritesForm())
	incomingRoutes.GET("/plumbing", controllers.GetPlumbingForm())
	incomingRoutes.GET("/interior", controllers.GetInteriorForm())
	incomingRoutes.GET("/accessories", controllers.GetAccessoriesForm())

	//footer
	incomingRoutes.GET("/about", controllers.GetAboutForm())
	incomingRoutes.GET("/for_legal_entities", controllers.GetAboutForm())
	incomingRoutes.GET("/for_users", controllers.GetForUsersForm())
	incomingRoutes.GET("/questions", controllers.GetQuestionsForm())

	incomingRoutes.GET("/login", controllers.GetLoginForm())
	incomingRoutes.GET("/signup", controllers.GetSignUpForm())

	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())

	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
