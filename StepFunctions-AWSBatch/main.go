package main

import (
	"database/sql"
	"flag"
	"math/rand"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Tweet struct {
	Text string `json:"text"`
}

type Tweets []Tweet

func main() {
	// コマンドライン引数
	userID, _ := strconv.Atoi(*flag.String("userID", "1", "user.id"))
	userNum, _ := strconv.Atoi(*flag.String("userNum", "50", "user.num"))
	flag.Parse()

	// 接続情報（一時的なDBのため公開状態を許容する）
	db, err := sql.Open("mysql", "admin:Yk080211@tcp(parallel-distributed-processing.cld5vrk9jap3.ap-northeast-1.rds.amazonaws.com:3306)/prod")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 指定の数のテストデータを作成
	tweets, err := makeTestData(userNum)
	if err != nil {
		panic(err.Error())
	}

	// tweetsテーブルにインサート
	err = insertTweets(db, userID, tweets)
	if err != nil {
		panic(err.Error())
	}
}

func makeTestData(num int) (Tweets, error) {
	var tweets Tweets
	for i := 0; i < num; i++ {
		var tweet Tweet

		// ランダム文字列の生成
		var letter = []rune("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん")

		b := make([]rune, num)
		for i := range b {
			b[i] = letter[rand.Intn(len(letter))]
		}
		tweet.Text = string(b)

		tweets = append(tweets, tweet)
	}
	return tweets, nil
}

func insertTweets(db *sql.DB, userID int, tweets Tweets) error {

	for _, tweet := range tweets {
		// プリペアードステートメント
		stmt, err := db.Prepare("INSERT INTO tweets(tw_user_id, text) VALUES(?,?);")
		if err != nil {
			return err
		}
		defer stmt.Close()

		// クエリ実行
		_, err = stmt.Exec(userID, tweet.Text)
		if err != nil {
			return err
		}
	}
	return nil
}
