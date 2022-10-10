package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"fmt"
	"encoding/json"
)

type errorResponse struct {
	Error errorResponseDetails `json:"error"`
}

type errorResponseDetails struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
}

const invalidJSONError = "could not create JSON response"

func JSONResponse(code int, body interface{}) (events.APIGatewayProxyResponse, error) {
	bb, err := json.Marshal(body)
	if err != nil {
		return JSONErrResponse(http.StatusInternalServerError, invalidJSONError)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: code, 
		Body: string(bb), 
		Headers: map[string]string{"Content-Type": "application/json"},
		}, nil
}

func JSONErrResponse(code int, message string) (events.APIGatewayProxyResponse, error) {
	body := errorResponse{Error: errorResponseDetails{code, message, http.StatusText(code)}}
	bb, err := json.Marshal(body)
	if err != nil {
		code := http.StatusInternalServerError
		message := invalidJSONError
		status := http.StatusText(code)
		bodyText := fmt.Sprintf(
			`{"error":{"code":%d,"message":"%s", "status":"%s"}}`, 
			code,
			message, 
			status,
		)

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError, 
			Body: bodyText, 
			Headers: map[string]string{"Content-Type": "application/json"},
		}, nil

	}

	return events.APIGatewayProxyResponse{
		StatusCode: code, 
		Body: string(bb), 
		Headers: map[string]string{"Content-Type": "application/json"},
	}, nil
}

func handler(e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return JSONResponse(200, "Hello World")
}

func main() {
	lambda.Start(handler)
}