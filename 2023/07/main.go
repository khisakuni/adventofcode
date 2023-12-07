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
	'J': 1,
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
		var jokers int
		for _, card := range hand {
			if card == 'J' {
				jokers++
			} else {
				cards[cardScore(card)]++
			}

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
				break
			// 4 of a kind
			case count == 4:
				category = 5
				break
			case count == 3:
				// Full house
				if hasPair {
					category = 4
					break
				}
				has3OAK = true
			case count == 2:
				// 2 pair
				if hasPair {
					category = 2
					break
				}
				if has3OAK {
					category = 4
					break
				}
				hasPair = true
			}
		}

		if category < 0 {
			if has3OAK {
				category = 3
			} else if hasPair {
				category = 1
			} else {
				category = 0
			}
		}

		if jokers > 0 {
			switch category {
			case 6:
			case 5:
				if jokers > 0 {
					category++
				}
			case 4:
				if jokers > 0 {
					category += jokers
				}
			case 3:
				if jokers > 0 {
					// 4 of a kind or higher.
					category = 5 + jokers - 1
				}
			case 2:
				if jokers > 0 {
					// to full house
					category = 4
				}
			case 1:
				if jokers == 1 {
					category = 3
				}
				if jokers > 1 {
					category = 5 + jokers - 2
				}
			case 0:
				if jokers == 1 {
					category = 1
				}
				if jokers == 2 {
					category = 3
				}
				if jokers > 2 {
					category = 5 + jokers - 3
				}
			}
		}

		switch category {
		case 6:
			name = "5 of a kind"
		case 5:
			name = "4 of a kind"
		case 4:
			name = "full house"
		case 3:
			name = "3 of a kind"
		case 2:
			name = "2 pair"
		case 1:
			name = "1 pair"
		case 0:
			name = "high card"
		}

		if category > 6 {
			category = 6
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
