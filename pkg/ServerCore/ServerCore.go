package ServerCore

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	CustomerServices "github.com/shashijangra22/bootcamp-project/pkg/customer/services"
	OrderServices "github.com/shashijangra22/bootcamp-project/pkg/order/services"
	RestaurantServices "github.com/shashijangra22/bootcamp-project/pkg/restaurant/services"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/shashijangra22/bootcamp-project/pkg/Models"
	"github.com/shashijangra22/bootcamp-project/pkg/customer"
	"github.com/shashijangra22/bootcamp-project/pkg/order"
	"github.com/shashijangra22/bootcamp-project/pkg/restaurant"
)

type Server struct{}

var DB *dynamodb.DynamoDB

type AWS_STRUCT struct {
	AWS_KEY_ID     string
	AWS_SECRET_KEY string
	REGION         string
}

var secret AWS_STRUCT

func createDBSession(filename string) *dynamodb.DynamoDB {
	secretsFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening secrets.json!")
		os.Exit(1)
	}
	defer secretsFile.Close()
	byteValue, _ := ioutil.ReadAll(secretsFile)
	json.Unmarshal(byteValue, &secret)
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(secret.REGION),
		// Endpoint: aws.String("http://localhost:8000"), // uncomment this to use outside the container
		Endpoint:    aws.String("http://192.168.0.2:8000"), // [use http] comment this to connect to cloud dynamodDB
		Credentials: credentials.NewStaticCredentials(secret.AWS_KEY_ID, secret.AWS_SECRET_KEY, ""),
	}))
	db := dynamodb.New(sess)
	return db
}

func init() {
	path, _ := os.Getwd()
	DB = createDBSession(path + "/secrets.json")
}

func (*Server) AddCustomer(ctx context.Context, req *customer.Customer) (*customer.Customer, error) {
	fmt.Println("AddCustomer Function called... ")
	id := req.GetID()
	name := req.GetName()
	address := req.GetAddress()
	phone := req.GetPhone()
	cst := Models.Customer{ID: id, Name: name, Address: address, Phone: phone}
	CustomerServices.AddOne(DB, cst)
	return req, nil
}

func (*Server) GetCustomer(ctx context.Context, req *customer.IDRequest) (*customer.Customer, error) {
	fmt.Println("GetCustomer is called... ")
	id := req.GetID()
	res := CustomerServices.GetOne(DB, id)
	return res, nil
}

// fetch customers from db and give it as response to client
func (*Server) GetCustomers(ctx context.Context, req *customer.NoParamRequest) (*customer.Customers, error) {
	fmt.Println("GetCustomers is called... ")
	allCustomers := CustomerServices.GetAll(DB)
	res := &customer.Customers{Customers: allCustomers}
	return res, nil
}

func (*Server) AddOrder(ctx context.Context, req *order.Order) (*order.Order, error) {
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

func (*Server) GetOrder(ctx context.Context, req *order.IDRequest) (*order.Order, error) {
	fmt.Println("GetOrder Function called... ")
	id := req.GetID()
	res := OrderServices.GetOne(DB, id)
	return res, nil
}

func (*Server) GetOrders(ctx context.Context, req *order.NoParamRequest) (*order.Orders, error) {
	fmt.Println("GetOrders is called...")
	allOrders := OrderServices.GetAll(DB)
	res := &order.Orders{Orders: allOrders}
	return res, nil
}

func (*Server) GetRestaurant(ctx context.Context, req *restaurant.IDRequest) (*restaurant.Restaurant, error) {
	fmt.Println("GetRestaurant Function called... ")
	id := req.GetID()
	res := RestaurantServices.GetOne(DB, id)
	return res, nil
}

func (*Server) GetRestaurants(ctx context.Context, req *restaurant.NoParamRequest) (*restaurant.Restaurants, error) {
	fmt.Println("GetRestaurants is called...")
	allRestaurants := RestaurantServices.GetAll(DB)
	res := &restaurant.Restaurants{Restaurants: allRestaurants}
	return res, nil
}
