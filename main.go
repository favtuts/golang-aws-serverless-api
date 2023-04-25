package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrPersonNotProvided is thrown when a person json data string is not provided
	ErrPersonNotProvided = errors.New("no person json data was provided in the HTTP body")
)

func main() {
	lambda.Start(handler)
}

// handler is your lambda function handler
// it uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// however you could use other event sources (S3, Kinesis etcc), or JSON-decoded primitive types such as 'string'.
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no data is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrPersonNotProvided
	}

	// Decode Http request body into Person object
	var person Person
	err := json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// Build the hello message to the Person
	msg := fmt.Sprintf("Hello %v %v", *person.FirstName, *person.LastName)
	responseBody := ResponseBody{
		Message: &msg,
	}

	// Build the json string from the Person object
	jbytes, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// Return the response
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jbytes),
	}

	return response, nil
}

type Person struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type ResponseBody struct {
	Message *string `json:"message"`
}
