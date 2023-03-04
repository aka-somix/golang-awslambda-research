package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type AnyEvent struct {}

func HandleRequest(ctx context.Context, event AnyEvent) (string, error) {
        return "Hello World", nil
}

func main() {
        lambda.Start(HandleRequest)
}
