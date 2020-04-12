package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"math/rand"
	"time"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func hello(event MyEvent) (MyResponse, error) {
	printWordCards()
	return MyResponse{Message: fmt.Sprintf("Hello test %s!!", event.Name)}, nil
}

func main() {
	lambda.Start(hello)
	//printWordCards()
}

func printWordCards() {
	words := shuffle(words())
	for i := 0; i < 25; i++ {
		fmt.Println(i+1, words[i])
	}
}

// @todo: 25個に絞らないと無駄
func shuffle(words []string) []string {
	rand.Seed(time.Now().UnixNano())
	n := len(words)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		words[i], words[j] = words[j], words[i]
	}

	return words
}

func words() []string {
	return []string {
		"ベンチ",
		"スケート",
		"ケーキ",
		"スフィンクス",
		"サクラ",
		"映画",
		"ゴルフ",
		"バスケット",
		"ポテトチップス",
		"バイキング",
		"バレンタイン",
		"ラム",
		"ライフル",
		"ヤギ",
		"クレオパトラ",
		"バッタ",
		"ダイコン",
		"クシ",
		"ノコギリ",
		"サドル",
		"火山",
		"平和",
		"メガネ",
		"ロケット",
		"花婿",
		"カニ",
		"太陽",
		"ハワイ",
		"ミイラ",
		"ボクサー",
		"ブタ",
		"バイオリン",
		"ビール",
		"ブラジャー",
		"大学",
		"妹",
		"歯医者",
		"虹",
		"剣",
		"シャンプー",
		"クリーニング",
		"テント",
		"風船",
		"風呂",
		"バーベキュー",
	}
}