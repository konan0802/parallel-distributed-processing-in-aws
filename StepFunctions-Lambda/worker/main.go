package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type User struct {
	ID   int
	Name string
	Num  int
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// 接続情報（一時的なDBのため公開状態を許容する）
	db, err := sql.Open("mysql", "admin:Yk080211@tcp(parallel-distributed-processing.cld5vrk9jap3.ap-northeast-1.rds.amazonaws.com:3306)/prod")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 指定の数のテストデータを作成

	// tweetsテーブルにインサート

	// usersテーブルのアップデート

	return events.APIGatewayProxyResponse{
		Body:       string(users[0].Name),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
