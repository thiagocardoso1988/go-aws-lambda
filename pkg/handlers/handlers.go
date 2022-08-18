package handlers

import (
	"net/http"

	"github.com/thiagocardoso1988/go-aws-lambda/pkg/models/user"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaclient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]

	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaclient)

		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{
				aws.String(err.Error()),
			})
		}

		return apiResponse(http.StatusOK, result)
	}

	result, err := user.FetchUsers(tableName, dynaclient)

	if err != nil {
		return apiResponse(http.StatusBadGateway, ErrorBody{
			aws.String(err.Error()),
		})
	}

	return apiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaclient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(req, tableName, dynaclient)

	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}

	return apiResponse(http.StatusCreated, result)
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaclient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.UpdateUser(req, tableName, dynaclient)

	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaclient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	err := user.DeleteUser(req, tableName, dynaclient)

	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}

	return apiResponse(http.StatusOK, nil)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
