package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler054
//Poker Hands
var suits = map[string]int{
	"D": 1,
	"H": 2,
	"S": 3,
	"C": 4,
}

var numbers = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

const (
	highCard      = iota
	onePair       = iota
	twoPairs      = iota
	threeKind     = iota
	straight      = iota
	flush         = iota
	fullHouse     = iota
	fourKind      = iota
	straightFlush = iota
	royalFlush    = iota
)

type card struct {
	value int
	suit  int
}

type cardList []*card

type hand struct {
	cards        cardList
	name         string
	rank         int
	countByValue map[int][]cardList //first element has the higher count of the cards with the same value
	straight     bool               //all the cards are consecutive
	flush        bool
}

func NewHand(cards cardList, name string) *hand {
	h := &hand{}
	h.straight = true
	h.name = name
	h.flush = true

	h.cards = make(cardList, len(cards))
	copy(h.cards, cards)

	//sort desc by value
	sort.Slice(h.cards, func(i, j int) bool {
		return h.cards[i].value > h.cards[j].value
	})

	//exception, five high straight
	if h.cards[0].value == numbers["A"] &&
		h.cards[1].value == numbers["5"] &&
		h.cards[2].value == numbers["4"] &&
		h.cards[3].value == numbers["3"] &&
		h.cards[4].value == numbers["2"] {

		//transform Ace to Low
		h.cards[0].value = numbers["1"]
		//resort because now ace must be last
		sort.Slice(h.cards, func(i, j int) bool {
			return h.cards[i].value > h.cards[j].value
		})
	}

	tmp := make(map[int]cardList) //histogram based on card value
	for i, c := range h.cards {
		tmp[c.value] = append(tmp[c.value], c)

		if i == 0 {
			continue
		}

		if c.value != h.cards[i-1].value-1 {
			h.straight = false
		}
		if c.suit != h.cards[i-1].suit {
			h.flush = false
		}
	}

	h.countByValue = make(map[int][]cardList, 1)
	for _, cards := range tmp {
		h.countByValue[len(cards)] = append(h.countByValue[len(cards)], cards)
	}
	return h
}

func main() {
	for _, table := range input() {
		fmt.Println(getResult(table[0], table[1]).name)
	}
}

func getResult(p1, p2 *hand) *hand {
	r1, valuesToComp1 := getHandRank(p1)
	r2, valuesToComp2 := getHandRank(p2)

	if r1 > r2 {
		return p1
	}

	if r1 == r2 {
		for i, val := range valuesToComp1 {
			if val > valuesToComp2[i] {
				return p1
			}

			if val < valuesToComp2[i] {
				return p2
			}
			//are equal skip, compare the next one
		}
		log.Panicf("equal ranks and values %v", p1.cards)
	}

	//if r1 < r2 {
	return p2
}

/*
	Function does 2 things:
1. establish the rank of the found pattern
2. returns a list of ints to be compared, in case
of a rank equality. First value is most important.
*/
func getHandRank(h *hand) (int, []int) {
	var valuesToComp []int
	//higher cards, in order of importance when comparing
	for _, c := range h.cards {
		valuesToComp = append(valuesToComp, c.value)
	}

	unshift := func(x int) {
		valuesToComp = append([]int{x}, valuesToComp...)
	}

	_, hasThreeOfKind := h.countByValue[3]
	_, hasTwoOfKind := h.countByValue[2]

	hasRoyal := h.flush && h.straight &&
		h.cards[0].value == numbers["A"]

	if hasRoyal {
		return royalFlush, valuesToComp
	}

	if h.flush {
		return straightFlush, valuesToComp
	}

	if fours, ok := h.countByValue[4]; ok {
		unshift(fours[0][0].value)
		return fourKind, valuesToComp
	}

	if hasTwoOfKind && hasThreeOfKind {
		unshift(h.countByValue[2][0][0].value)
		//first we compare the 3 of a kind value, we add last
		unshift(h.countByValue[3][0][0].value)
		return fullHouse, valuesToComp
	}

	if h.flush {
		return flush, valuesToComp
	}

	if h.straight {
		return straight, valuesToComp
	}

	//3 of a kind: score: 60-75
	if hasThreeOfKind {
		unshift(h.countByValue[3][0][0].value)
		return threeKind, valuesToComp
	}

	//two different pairs, score: 30-60
	if hasTwoOfKind {
		pairs := h.countByValue[2]
		if len(pairs) == 2 {
			//is complicated, biggest pair first
			valueA := h.countByValue[2][0][0].value
			valueB := h.countByValue[2][1][0].value
			if valueA > valueB {
				unshift(valueA)
				unshift(valueB)
			} else {
				unshift(valueB)
				unshift(valueA)
			}
			return twoPairs, valuesToComp
		}

		//1 pair score: 15-30
		unshift(pairs[0][0].value)
		return onePair, valuesToComp
	}

	return highCard, valuesToComp
}

func input() (hands [][]*hand) {
	var T int

	_, err := fmt.Scanln(&T)
	if err != nil {
		log.Panic(err)
	}

	hands = make([][]*hand, 0)
	in := bufio.NewReader(os.Stdin)

	for l := T; l > 0; l-- {
		line, _ := in.ReadString('\n')
		p1, p2 := strToHands(line)

		hands = append(hands, []*hand{
			p1, p2,
		})
	}

	return
}

//transform "AD 2C 3S 4H 5C 3D 6C 7D TS QD" to structs
func strToHands(in string) (p1, p2 *hand) {
	cards := make(cardList, 10)
	for i, cardAsStr := range strings.Split(in, " ") {
		cards[i] = &card{numbers[string(cardAsStr[0])], suits[string(cardAsStr[1])]}
	}
	p1 = NewHand(cards[:5], "Player 1")
	p2 = NewHand(cards[5:], "Player 2")
	return
}
