package main

import (
	"MyApp/pkg/Models"
	"MyApp/pkg/customer"
	CustomerServices "MyApp/pkg/customer/services"
	"MyApp/pkg/order"
	OrderServices "MyApp/pkg/order/services"
	"MyApp/pkg/restaurant"
	RestaurantServices "MyApp/pkg/restaurant/services"
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

func (*server) AddCustomer(ctx context.Context, req *customer.Customer) (*customer.Customer, error) {
	fmt.Println("AddCustomer Function called... ")
	id := req.GetID()
	name := req.GetName()
	address := req.GetAddress()
	phone := req.GetPhone()
	cst := Models.Customer{ID: id, Name: name, Address: address, Phone: phone}
	CustomerServices.AddOne(DB, cst)
	return req, nil
}

func (*server) GetCustomer(ctx context.Context, req *customer.IDRequest) (*customer.Customer, error) {
	fmt.Println("GetCustomer is called... ")
	id := req.GetID()
	res := CustomerServices.GetOne(DB, id)
	return res, nil
}

// fetch customers from db and give it as response to client
func (*server) GetCustomers(ctx context.Context, req *customer.NoParamRequest) (*customer.Customers, error) {
	fmt.Println("GetCustomers is called... ")
	allCustomers := CustomerServices.GetAll(DB)
	res := &customer.Customers{Customers: allCustomers}
	return res, nil
}

func (*server) AddOrder(ctx context.Context, req *order.Order) (*order.Order, error) {
	fmt.Println("AddOrder Function called... ")
	itemlist := req.GetItemLine()
	var items []Models.Item
	for i := range itemlist {
		items = append(items, Models.Item{
			Name:  itemlist[i].GetName(),
			Price: itemlist[i].GetPrice(),
		})
	}
	orderDetails := Models.Order{ID: req.GetID(), C_ID: req.GetC_ID(), ItemLine: items, Price: req.GetPrice(), Discount: req.GetDiscount()}
	OrderServices.Add(DB, orderDetails)
	return req, nil
}

func (*server) GetOrder(ctx context.Context, req *order.IDRequest) (*order.Order, error) {
	fmt.Println("GetOrder Function called... ")
	id := req.GetID()
	res := OrderServices.GetOne(DB, id)
	return res, nil
}

func (*server) GetOrders(ctx context.Context, req *order.NoParamRequest) (*order.Orders, error) {
	fmt.Println("GetOrders is called...")
	allOrders := OrderServices.GetAll(DB)
	res := &order.Orders{Orders: allOrders}
	return res, nil
}

func (*server) GetRestaurant(ctx context.Context, req *restaurant.IDRequest) (*restaurant.Restaurant, error) {
	fmt.Println("GetRestaurant Function called... ")
	id := req.GetID()
	res := RestaurantServices.GetOne(DB, id)
	return res, nil
}

func (*server) GetRestaurants(ctx context.Context, req *restaurant.NoParamRequest) (*restaurant.Restaurants, error) {
	fmt.Println("GetRestaurants is called...")
	allRestaurants := RestaurantServices.GetAll(DB)
	res := &restaurant.Restaurants{Restaurants: allRestaurants}
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
	restaurant.RegisterRestaurantServiceServer(s, &server{})

	if s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
