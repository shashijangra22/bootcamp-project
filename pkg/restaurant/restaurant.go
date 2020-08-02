package restaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var RSC RestaurantServiceClient

// client call to get all the restaurants
func GetAll(c *gin.Context) {
	req := &NoParamRequest{}
	res, err := RSC.GetRestaurants(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"_": res.Restaurants})
}

// To Add a new restaurant
// func Add(c *gin.Context) {
// 	req := &Restaurant{}

// 	res, err := RSC.AddRestaurant(c, req)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"response": res.DummyRes,
// 	})
// }
