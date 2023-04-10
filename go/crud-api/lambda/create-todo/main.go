package main

import (
	"encoding/json"
	"fmt"

	"github.com/aka-somix/crud-api/_shared/Todo"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/valyala/fastjson"
)

type CreateTodoRequestBody struct {
	Name string
	Description string
}

type CreateTodoResponseBody struct {
	Id string
}

func saveTodo(newTodo Todo.Todo) error {
	return nil;
}


func CreateTodo(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Event Received: %v \n", req)

	response := events.APIGatewayProxyResponse{}

	// validates json and returns error if not working
	if err := fastjson.Validate(req.Body); err != nil {
		response.Body = "Error: Invalid JSON payload ||| " + fmt.Sprint(err) + " Body Obtained" + "||||" + req.Body
		response.StatusCode = 400

		return response, nil
	}

	// Parse the request body
	var body CreateTodoRequestBody;

	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
			return events.APIGatewayProxyResponse{}, err
	}

	// Validate input
	if body.Name == "" {
		response.Body = fmt.Sprintf("Error. Invalid Body: %q. Missing Name", body)
		response.StatusCode = 400

		return response, nil
	}

	// Create and Store the Todo
	newTodo := Todo.New(body.Name, body.Description)

	// TODO aka-somix: could be deleted?
	fmt.Printf("Todo Created is: %v \n", newTodo)

	if err := saveTodo(newTodo); err != nil {
			return events.APIGatewayProxyResponse{}, err
	}

	// Build the response body
	jbytes, _ := json.Marshal(CreateTodoResponseBody{Id: newTodo.Id})
	jstr := string(jbytes)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: jstr,
	}, nil
}

func main() {
        lambda.Start(CreateTodo)
}
