package services

import (
	"MyApp/pkg/Models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func Add(db *dynamodb.DynamoDB, ord Models.Order) {
	orderDynAttr, err := dynamodbattribute.MarshalMap(ord)
	if err != nil {
		panic("Cannot map the values given in Order struct for post request...")
	}
	params := &dynamodb.PutItemInput{
		TableName: aws.String("orders"),
		Item:      orderDynAttr,
	}
	_, err = db.PutItem(params)
	if err != nil {
		panic("Error in putting the order item")
	}
}
