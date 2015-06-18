package cards

import "math/rand"

// No joker for the moment
type deck [52]Card

//NewDeck : Initializes a deck
func NewDeck() deck {
	var ret deck
	i := 0
	for c := Spade; c <= Diamond; c++ {
		for f := Two; f <= Ace; f++ {
			card, err := NewCard(f, c)
			if err == nil {
				ret[i] = card
				i++
			}
		}
	}
	return ret
}

//Shuffle : Shuffle a deck
func (d *deck) Shuffle() {
	for i := 0; i < len(d); i++ {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}
