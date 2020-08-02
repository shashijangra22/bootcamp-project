package main

import (
	"MyApp/pkg/populate"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("us-east-1"),
	}))
	db := dynamodb.New(sess)

	args := os.Args[1:]
	if len(args) == 1 {
		path, _ := os.Getwd()
		populate.Customers(db, path+"/"+args[0]+"/customers.csv")
		populate.Orders(db, path+"/"+args[0]+"/orders.csv")
		populate.Restaurants(db, path+"/"+args[0]+"/restaurants.csv")
		fmt.Println("Done populating all the tables :)")
	} else {
		fmt.Println("Please give directory path of DB schema as an argument!")
	}
}
