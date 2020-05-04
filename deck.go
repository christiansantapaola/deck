package deck

import (
	"math/rand"
	"time"
)

type color string

const (
	bastoni = "Bastoni"
	coppe   = "Coppe"
	denara  = "Denara"
	spade   = "Spade"
)

// Card is a pair value, color
type Card struct {
	Value int
	color color
}

// Deck is a collection of 40 Card
type Deck struct {
	cards     [40]Card
	pickOrder []int
	top       int
	rand      *rand.Rand
}

// NewDeck will create a new deck structure
func NewDeck() *Deck {
	d := Deck{}
	for index := range d.cards {
		d.cards[index].Value = 1 + index%10
	}
	source := rand.NewSource(time.Now().UnixNano())
	d.rand = rand.New(source)
	return &d
}

// Shuffle will shuffle the deck
func (d *Deck) Shuffle() {
	//d.rand.Seed(time.Now().UnixNano())
	d.pickOrder = d.rand.Perm(40)
	//fmt.Println(d.pickOrder)
}

// Draw will draw a card from the top of the deck
func (d *Deck) Draw() *Card {
	if d.top+1 > 40 {
		return nil
	}
	card := d.cards[d.pickOrder[d.top]]
	d.top++
	return &card
}

// Reset will reset the state of the draw
func (d *Deck) Reset() {
	d.top = 0
}
