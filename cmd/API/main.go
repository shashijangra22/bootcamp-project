package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shashijangra22/bootcamp-project/pkg/auth"
	"github.com/shashijangra22/bootcamp-project/pkg/customer"
	"github.com/shashijangra22/bootcamp-project/pkg/order"
	"github.com/shashijangra22/bootcamp-project/pkg/restaurant"
	"google.golang.org/grpc"
)

// homepage of API
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Team 2": "Hello from Aadithya, Abhishek, Priya, Shashi!",
	})
}

func main() {

	fmt.Println("Hello from the ginAPI :)")
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v", err)
		os.Exit(1)
	}

	defer conn.Close()

	customer.CSC = customer.NewCustomerServiceClient(conn)
	order.OSC = order.NewOrderServiceClient(conn)
	restaurant.RSC = restaurant.NewRestaurantServiceClient(conn)

	router := gin.Default()

	loginRouter := router.Group("/login")
	loginRouter.POST("/", auth.Login)

	apiRouter := router.Group("/api")
	apiRouter.Use(auth.VerifyUser)

	apiRouter.GET("/", Index)

	apiRouter.GET("/orders", order.GetAll)
	apiRouter.GET("/order/:id", order.GetOne)
	apiRouter.POST("/order", order.Add)

	apiRouter.GET("/customers", customer.GetAll)
	apiRouter.GET("/customer/:id", customer.GetOne)
	apiRouter.POST("/customer", customer.Add)

	apiRouter.GET("/restaurants", restaurant.GetAll)
	apiRouter.GET("/restaurant/:id", restaurant.GetOne)
	// apiRouter.POST("/restaurant", restaurant.Add)

	router.Run("localhost:9001")
}
