package order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var OSC OrderServiceClient

// client call to get all the orders
func GetAll(c *gin.Context) {
	req := &NoParamRequest{}
	res, err := OSC.GetOrders(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"_": res.Orders})
}

// client call to get a particular order
func GetOne(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	req := &IDRequest{ID: id}
	res, err := OSC.GetOrder(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"_": res})

}

// To Add a new order
// func Add(c *gin.Context) {
// 	req := &Order{}

// 	res, err := OSC.AddOrder(c, req)

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
