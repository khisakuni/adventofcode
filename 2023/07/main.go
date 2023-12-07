package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Name     string
	Cards    string
	Strength int64
	Bid      int
}

var cardToScore = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
}

var categories = make([][]Hand, 7)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		hand := parts[0]
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		cards := make([]int, 15)
		for _, card := range hand {
			cards[cardScore(card)]++
		}

		var name string
		var hasPair bool
		var has3OAK bool
		category := -1
		for _, count := range cards {
			switch {
			// 5 of a kind
			case count == 5:
				category = 6
				name = "5 of a kind"
				break
			// 4 of a kind
			case count == 4:
				category = 5
				name = "4 of a kind"
				break
			case count == 3:
				// Full house
				if hasPair {
					category = 4
					name = "full house"
					break
				}
				has3OAK = true
			case count == 2:
				// 2 pair
				if hasPair {
					category = 2
					name = "2 pair"
					break
				}
				if has3OAK {
					category = 4
					name = "full house"
					break
				}
				hasPair = true
			}
		}

		if category < 0 {
			if has3OAK {
				name = "3 of a kind"
				category = 3
			} else if hasPair {
				name = "1 pair"
				category = 1
			} else {
				category = 0
			}
		}

		var strength int64
		for i, c := range hand {
			s := cardScore(c)
			strength += int64(math.Pow10(10-(i*2))) * s
		}

		h := Hand{Cards: hand, Bid: bid, Strength: strength, Name: name}
		categories[category] = append(categories[category], h)
	}

	var winnings int
	i := 1
	for _, hands := range categories {
		sort.Slice(hands, func(i, j int) bool {
			return hands[i].Strength < hands[j].Strength
		})
		for _, hand := range hands {
			fmt.Printf("%v #%v (%v) bid: %v\n", hand.Cards, i, hand.Name, hand.Bid)
			winnings += (i * hand.Bid)
			i++
		}
	}

	fmt.Printf("winnings: %v\n", winnings)
}

func cardScore(card rune) int64 {
	s, ok := cardToScore[card]
	if !ok {
		var err error
		s, err = strconv.Atoi(string(card))
		if err != nil {
			panic(err)
		}
	}

	return int64(s)
}
