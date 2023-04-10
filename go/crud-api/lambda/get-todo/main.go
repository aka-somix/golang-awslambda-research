package main

import (
	"fmt"

	"github.com/aka-somix/crud-api/_shared/Todo"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type GetAllTodosResponseBody struct {
	Todos []Todo.Todo
}

type GetTodoResponseBody struct {
	Todos Todo.Todo
}

func GetTodo(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Event Received: %v \n", req)
	
	fmt.Printf("Path elements: %v \n", req.QueryStringParameters)

	if id, found := req.QueryStringParameters["id"]; found {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body: "Requested ID:" + id,
		}, nil
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body: "Requested All",
		}, nil
	}
}

func main() {
	lambda.Start(GetTodo)
}
