package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	todoId string
}

func CreateTodo(event any) (Response, error) {
	fmt.Printf("Event Received: %s \n", event)
	return Response{todoId: "1"}, nil
}

func main() {
        lambda.Start(CreateTodo)
}
