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

type Users []User

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// 接続情報（一時的なDBのため公開状態を許容する）
	db, err := sql.Open("mysql", "admin:Yk080211@tcp(parallel-distributed-processing.cld5vrk9jap3.ap-northeast-1.rds.amazonaws.com:3306)/prod")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// usersテーブルからタスクの取得
	users, err := fetchUsers(db)
	if err != nil {
		panic(err.Error())
	}

	// userごとに指定の数のテストデータを作成

	// tweetsテーブルにインサート

	// usersテーブルのアップデート

	return events.APIGatewayProxyResponse{
		Body:       string(users[0].Name),
		StatusCode: 200,
	}, nil
}

func fetchUsers(db *sql.DB) (Users, error) {
	// ユーザー一覧
	var users Users

	// ユーザー取得
	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	// ユーザーをマップ
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Num)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}
	return users, nil
}

func main() {
	lambda.Start(Handler)
}
