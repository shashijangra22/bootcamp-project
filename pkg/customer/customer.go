package customer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var CSC CustomerServiceClient

// client call to get all the customers
func GetAll(c *gin.Context) {
	req := &NoParamRequest{}
	res, err := CSC.GetCustomers(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"_": res.Customers})
}

// client call to get a particular customer
func GetOne(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	req := &IDRequest{ID: id}
	res, err := CSC.GetCustomer(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"_": res})
}

// To add new customer
// func Add(c *gin.Context) {

// 	req := &Customer{}
// 	res, err := CSC.AddCustomer(c, req)

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
