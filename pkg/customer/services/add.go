package services

import (
	"MyApp/pkg/Models"
	"MyApp/pkg/customer"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AddOne(db *dynamodb.DynamoDB, cst Models.Customer) *customer.Customer {
	cstItem, err := dynamodbattribute.MarshalMap(cst)
	if err != nil {
		panic("Cannot map the values given in Customer struct for post request...")
	}
	params := &dynamodb.PutItemInput{
		TableName: aws.String("customers"),
		Item:      cstItem,
	}
	output, err := db.PutItem(params)
	if err != nil {
		panic("Error in putting the customer item")
	}
	err = dynamodbattribute.UnmarshalMap(output.Attributes, &cst)
	res := &customer.Customer{ID: cst.ID, Name: cst.Name, Address: cst.Address, Phone: cst.Phone}
	return res
}
