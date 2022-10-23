package cmd

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

func Run() {
	run(context.Background())
}

func run(ctx context.Context) {
	if os.Getenv("AWS_EXECUTION_ENV") != "dev" {
		lambda.Start(handler)
	} else {
		r := events.APIGatewayProxyRequest{}
		_, err := handler(r)
		if err != nil {
			log.Fatal(err)
		}

	}
}

type Response struct {
	RequestMethod  string `json:"RequestMethod"`
	RequestBody    string `json:"RequestBody"`
	PathParameter  string `json:"PathParameter"`
	QueryParameter string `json:"QueryParameter"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// httpリクエストの情報を取得
	method := request.HTTPMethod
	body := request.Body
	pathParam := request.PathParameters["pathparam"]
	queryParam := request.QueryStringParameters["queryparam"]

	// レスポンスとして返すjson文字列を作る
	res := Response{
		RequestMethod:  method,
		RequestBody:    body,
		PathParameter:  pathParam,
		QueryParameter: queryParam,
	}
	jsonBytes, _ := json.Marshal(res)

	// 返り値としてレスポンスを返す
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}
