package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// usersテーブルからタスクの取得
	res := "gogogo"
	// userごとに指定の数のテストデータを作成

	// tweetsテーブルにインサート

	// usersテーブルのアップデート

	return events.APIGatewayProxyResponse{
		Body:       string(res),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
