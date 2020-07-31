package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var CSC CustomerServiceClient

// To get all the customers
func GetAll(c *gin.Context) {
	req := &NoParamRequest{}
	res, err := CSC.GetCustomers(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": res.DummyRes})
}

// To add new customer
func Add(c *gin.Context) {

	req := &Customer{}
	res, err := CSC.AddCustomer(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": res.DummyRes,
	})
}
