package main

import (
	"MyApp/pkg/populate"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		populate.Customers(args[0] + "/customers.csv")
		populate.Orders(args[0] + "/orders.csv")
		populate.Restaurants(args[0] + "/restaurants.csv")
		fmt.Println("Done populating all the tables :)")
	} else {
		fmt.Println("Please give directory path of DB schema as an argument!")
	}
}
