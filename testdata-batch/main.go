package main

import (
	"database/sql"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 接続情報（一時的なDBのため公開状態を許容する）
	db, err := sql.Open("mysql", "admin:Yk080211@tcp(parallel-distributed-processing.cld5vrk9jap3.ap-northeast-1.rds.amazonaws.com:3306)/prod")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for i := 0; i < 99; i++ {

		// ランダム文字列の生成
		var letter = []rune("abcdefghijklmnopqrstuvwxyz")

		num := rand.Intn(10-5) + 5
		b := make([]rune, num)
		for i := range b {
			b[i] = letter[rand.Intn(len(letter))]
		}
		name := string(b)

		// プリペアードステートメント
		stmt, err := db.Prepare("INSERT INTO users(name, num) VALUES(?,?);")
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()

		// クエリ実行
		_, err = stmt.Exec(name, 50)
		if err != nil {
			panic(err.Error())
		}
	}

}
