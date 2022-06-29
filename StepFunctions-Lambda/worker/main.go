package main

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Num  int    `json:"num"`
}

type Tweet struct {
	TwID int    `json:"twid"`
	Text string `json:"text"`
}

type Tweets []Tweet

func Handler(ctx context.Context, user User) (events.APIGatewayProxyResponse, error) {
	// 接続情報（一時的なDBのため公開状態を許容する）
	db, err := sql.Open("mysql", "admin:Yk080211@tcp(parallel-distributed-processing.cld5vrk9jap3.ap-northeast-1.rds.amazonaws.com:3306)/prod")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 指定の数のテストデータを作成
	tweets, err := makeTestData(user.Num)
	if err != nil {
		panic(err.Error())
	}

	// tweetsテーブルにインサート

	// usersテーブルのアップデート

	return events.APIGatewayProxyResponse{
		Body:       string(tweets[0].Text),
		StatusCode: 200,
	}, nil
}

func makeTestData(num int) (Tweets, error) {
	var tweets Tweets
	for i := 0; i < num; i++ {
		var tweet Tweet
		// ランダムIDの生成
		var twID int
		rand.Seed(time.Now().Unix())
		twID = rand.Intn(100)

		// ランダム文字列の生成
		var letter = []rune("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん")

		b := make([]rune, twID)
		for i := range b {
			b[i] = letter[rand.Intn(len(letter))]
		}
		tweet.TwID = twID
		tweet.Text = string(b)

		tweets = append(tweets, tweet)
	}
	return tweets, nil
}

func main() {
	lambda.Start(Handler)
}
