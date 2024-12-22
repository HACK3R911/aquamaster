package main

import (
	"log"
	"os"

	"aquamaster/controllers"
	"aquamaster/database"
	"aquamaster/middleware"
	"aquamaster/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	r := gin.New()

	r.LoadHTMLGlob("ui/html/pages/*.html")
	r.Static("/static", "ui/static")

	r.Use(gin.Logger())
	routes.UserRoutes(r)
	r.Use(middleware.Authentication())

	r.GET("/addtocart", app.AddToCart())
	r.GET("/removeitem", app.RemoveItem())
	r.GET("/listcart", controllers.GetItemFromCart())
	r.POST("/addaddress", controllers.AddAddress())
	r.PUT("/edithomeaddress", controllers.EditHomeAddress())
	r.PUT("/editworkaddress", controllers.EditWorkAddress())
	r.GET("/deleteaddresses", controllers.DeleteAddress())
	r.GET("/cartcheckout", app.BuyFromCart())
	r.GET("/instantbuy", app.InstantBuy())

	log.Fatal(r.Run(":" + port))
}
