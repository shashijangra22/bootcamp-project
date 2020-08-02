package main

import (
	"MyApp/pkg/customer"
	CustomerServices "MyApp/pkg/customer/services"
	"MyApp/pkg/order"
	OrderServices "MyApp/pkg/order/services"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"google.golang.org/grpc"
)

type server struct{}

var DB *dynamodb.DynamoDB

// fetch customers from db and give it as response to client
func (*server) GetCustomers(ctx context.Context, req *customer.NoParamRequest) (*customer.Customers, error) {
	fmt.Println("GetCustomers is called... ")
	allCustomers := CustomerServices.GetAll(DB)
	res := &customer.Customers{Customers: allCustomers}
	return res, nil
}

func (*server) GetOrders(ctx context.Context, req *order.NoParamRequest) (*order.Orders, error) {
	fmt.Println("GetOrders is called...")
	allOrders := OrderServices.GetAll(DB)
	res := &order.Orders{Orders: allOrders}
	return res, nil
}

func main() {
	// fire the gRPC Server
	fmt.Println("Hello from grpc server.")

	lis, err := net.Listen("tcp", "0.0.0.0:5051")

	if err != nil {
		log.Fatalf("Sorry failed to load server %v:", err)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("us-east-1"),
	}))
	DB = dynamodb.New(sess)

	s := grpc.NewServer()

	customer.RegisterCustomerServiceServer(s, &server{})
	order.RegisterOrderServiceServer(s, &server{})
	// restaurant.RegisterRestaurantServiceServer(s, &server{})

	if s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
