package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type card struct {
	number string
	suit   string
}

//SUITS Card suits
var SUITS = [4]string{"clubs", "hearts", "diamonds", "spades"}

func main() {
	var deck = createDeck()
	fmt.Println(deck)
}

func createDeck() []card {
	var deck []card
	for i := 0; i < len(SUITS); i++ {
		for j := 1; j <= 13; j++ {
			var cardNum = ""
			if j == 1 {
				cardNum = "A"
			} else if j == 11 {
				cardNum = "Jack"
			} else if j == 12 {
				cardNum = "Queen"
			} else if j == 13 {
				cardNum = "King"
			} else {
				cardNum = strconv.Itoa(j)
			}
			deck = append(deck, card{cardNum, SUITS[i]})
		}
	}
	Shuffle(deck)
	return deck
}

//Shuffle Takes in a slice and shuffles it
func Shuffle(vals []card) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}
