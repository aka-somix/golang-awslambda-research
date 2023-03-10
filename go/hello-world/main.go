package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() (string, error) {
	fmt.Println("TESTING WITH A BREAKPOINT")
	return "Hello World", nil
}

func main() {
        lambda.Start(Handler)
}
