package main

import (
	"fmt"
	"math/rand"
	"time"
)

type KeyCard struct {
	player1 PlayerKeyCard
	player2 PlayerKeyCard
}

type PlayerKeyCard struct{
	colors []Color
}

func (p PlayerKeyCard) text() string {
	text := ""
	for i := 0; i < len(p.colors); i++ {
		//text += fmt.Sprintf("%d: %s \n", i + 1, p.colors[i].toString())
		text += fmt.Sprintf("%s \n", p.colors[i].toString())
	}

	return text
}

type KeyCardColorParis struct {
	pairs [25]ColorPair
}

func (pairs KeyCardColorParis) shuffle() KeyCardColorParis {
	rand.Seed(time.Now().UnixNano())
	shuffled := KeyCardColorParis{pairs: pairs.pairs}
	n := len(shuffled.pairs)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		shuffled.pairs[i], shuffled.pairs[j] = shuffled.pairs[j], shuffled.pairs[i]
	}

	return shuffled
}

func (pairs KeyCardColorParis) makeKeyCard() KeyCard {
	var player1Colors []Color
	for i := 0; i < len(pairs.pairs); i++ {
		player1Colors = append(player1Colors, pairs.pairs[i].player1)
	}

	var player2Colors []Color
	for i := 0; i < len(pairs.pairs); i++ {
		player2Colors = append(player2Colors, pairs.pairs[i].player2)
	}

	return KeyCard {
		player1: PlayerKeyCard{colors: player1Colors},
		player2: PlayerKeyCard{colors: player2Colors},
	}
}

type ColorPair struct {
	player1 Color
	player2 Color
}

// @todo: 表現できる値をEnumみたいに制限したい
type Color string

func (c Color) toString() string {
	return string(c)
}

func main() {
	printKeyCard()
}

func printKeyCard() {
	keyCard := keyCardColorPairs().makeKeyCard()
	fmt.Println("---------------player1KeyCard---------------------")
	fmt.Println(keyCard.player1.text())
	fmt.Println("---------------player1KeyCard---------------------")
	//fmt.Println("-\n\n\n--\n\n\n---\n\n\n----\n\n\n ")
	fmt.Println("---------------player2KeyCard---------------------")
	fmt.Println(keyCard.player2.text())
	fmt.Println("---------------player2KeyCard---------------------")
}

func keyCardColorPairs() KeyCardColorParis {
	green := Color("green")
	black := Color("black")
	white := Color("white")

	pairs := KeyCardColorParis{pairs: [25]ColorPair {	// 25個ない時はエラーにしたい
		{black, white},
		{black, black},
		{black, green},
		{green, black},
		{green, green},
		{green, green},
		{green, green},
		{green, white},
		{green, white},
		{green, white},
		{green, white},
		{green, white},
		{white, black},
		{white, green},
		{white, green},
		{white, green},
		{white, green},
		{white, green},
		{white, white},
		{white, white},
		{white, white},
		{white, white},
		{white, white},
		{white, white},
		{white, white},
	}}

	return pairs.shuffle()
}