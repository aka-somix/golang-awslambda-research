package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func CreateTodo(event any) (string, error) {
	fmt.Printf("Event Received: %s \n", event)
	return "Hello World", nil
}

func main() {
        lambda.Start(CreateTodo)
}
