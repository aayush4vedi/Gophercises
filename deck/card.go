//go:generate stringer -type=Suit,Rank  
//cmd: $ go generate

package deck

import "fmt"

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // this is a special case
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New() []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {      // new way of traversal.Just fucking look at it
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	return cards
}