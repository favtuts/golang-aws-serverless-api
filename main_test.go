package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	request events.APIGatewayProxyRequest
	expect  string
	err     error
}

func TestHandlerIfPersonNotProvided(t *testing.T) {
	testCase := TestCase{
		// Test that the handler response ErrPersonNotProvided
		// when no person json data is provided in the HTTP body
		request: events.APIGatewayProxyRequest{Body: ""},
		expect:  "",
		err:     ErrPersonNotProvided,
	}

	response, err := handler(testCase.request)
	assert.IsType(t, testCase.err, err)
	assert.Equal(t, testCase.expect, response.Body)
}

func TestHandlerIfPersonDataProvided(t *testing.T) {
	testCase := TestCase{
		// Test that the handler response with the correct response
		// when no person json data is provided in the HTTP body
		request: events.APIGatewayProxyRequest{Body: "{\"firstName\":\"Shane\",\"lastName\":\"Van Boening\"}"},
		expect:  "{\"message\":\"Hello Shane Van Boening\"}",
		err:     nil,
	}

	response, err := handler(testCase.request)
	assert.IsType(t, testCase.err, err)
	assert.Equal(t, testCase.expect, response.Body)
}
