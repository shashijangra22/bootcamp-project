package services

import (
	"MyApp/pkg/Models"
	"MyApp/pkg/order"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func GetAll(db *dynamodb.DynamoDB) []*order.Order {
	var allOrders []*order.Order
	// Create the Expression to fill the input struct with.
	filt := expression.Name("ID").GreaterThan(expression.Value(0))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("C_ID"), expression.Name("R_ID"), expression.Name("ItemLine"), expression.Name("Price"), expression.Name("Discount"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression for getting all Orders")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("orders"),
	}
	// Make the DynamoDB Query API call
	result, err := db.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed for Order table fetched")
		fmt.Println((err.Error()))
		os.Exit(1)
	}
	for _, i := range result.Items {
		orderItem := Models.Order{}
		var itemLine []*order.Item
		err = dynamodbattribute.UnmarshalMap(i, &orderItem)
		if err != nil {
			fmt.Println("Got error unmarshalling order table")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, item := range orderItem.ItemLine {
			itemLine = append(itemLine, &order.Item{Name: item.Name, Price: item.Price})
		}
		allOrders = append(allOrders, &order.Order{ID: orderItem.ID, C_ID: orderItem.C_ID, R_ID: orderItem.R_ID, ItemLine: itemLine, Price: orderItem.Price, Discount: orderItem.Discount})
	}
	return allOrders
}
