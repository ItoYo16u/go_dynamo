package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Record struct {
	Id   int    `dynamodbav:"id"`
	Name string `dynamodbav:"name"`
}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("http://localhost:8000"),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	}))

	tableName := "test"
	dynamoDB := dynamodb.New(sess)
	fmt.Println("connectionEstablished")
	/*
			input := &dynamodb.CreateTableInput{
				AttributeDefinitions: []*dynamodb.AttributeDefinition{
					{
						AttributeName: aws.String("id"),
						AttributeType: aws.String("N"),
					},
					{
						AttributeName: aws.String("name"),
						AttributeType: aws.String("S"),
					},
				},
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("id"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("name"),
						KeyType:       aws.String("RANGE"),
					},
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
				TableName: aws.String(tableName),
			}

			_, err2 := dynamoDB.CreateTable(input)

			if err2 != nil {
				fmt.Println("Got error calling CreateTable:")
				fmt.Println(err2.Error())
				os.Exit(1)
			}
		    fmt.Println("create table!")
	*/

	record := Record{
		Id:   128,
		Name: "test",
	}
	//	av, err3 := dynamodbattribute.MarshalMap(record)
	//	if err3 != nil {
	//		fmt.Println("Got error marshalling new record:")
	//		fmt.Println(err3.Error())
	//		os.Exit(1)
	//	}

	//putItemInput := &dynamodb.PutItemInput{
	//		TableName: aws.String(tableName),
	//		Item:      av,
	//	}

	param := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName), // テーブル名を指定

		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String("3"),
			},
			"name": {
				S: aws.String(record.Name),
			},
		},

		//あとは返してくる情報の種類を指定する
		ReturnConsumedCapacity:      aws.String("NONE"),
		ReturnItemCollectionMetrics: aws.String("NONE"),
		ReturnValues:                aws.String("NONE"),
	}

	_, err4 := dynamoDB.UpdateItem(param)
	if err4 != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err4.Error())
		os.Exit(1)
	}

	var records []Record = []Record{}
	result, err5 := dynamoDB.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	//.Query(queryParams)
	if err5 != nil {
		panic(err5)
	}

	fmt.Printf("%s", result)

	for _, rec := range result.Items {
		fmt.Printf("%v\n", rec)
		record := Record{}
		errx := dynamodbattribute.UnmarshalMap(rec, &record)
		if errx != nil {
			fmt.Printf("failed to unmarshal rec")
		}
		records = append(records, record)
	}
	fmt.Printf("%v", records)

	description, err6 := dynamoDB.DescribeTable(
		&dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		})
	if err6 != nil {

	}
	fmt.Printf("%s", description)
}
