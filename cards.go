package cards

import "fmt"
import "math/rand"

//IsValid checks the validity of a Face
func (f Face) IsValid() bool {
	return (f >= Two && f <= Ace)
}

//IsValid returns true if a color is valid
func (c Color) IsValid() bool {
	return (c >= Spade && c <= Diamond)
}

//IsValid checks the validity of a Value
func (v Value) IsValid() bool {
	ret := false
	for _, val := range values {
		if val == v {
			ret = true
			break
		}
	}
	return ret
}

//Value : Returns the value associated to a card, if the card is invalid, err is != nil
func (c Card) Value() (ret []Value, err error) {
	ret, err = []Value{0}, nil
	if !c.face.IsValid() {
		return ret, fmt.Errorf("%v is not a valid card", c.face)
	}
	ret, ok := FaceValue[c.face]
	if !ok {
		return ret, fmt.Errorf("Unable to find the associated value to: %v", c.face)
	}
	return ret, err

}

//RandFace : Returns a random face
func RandFace() Face {
	return Face(rand.Int31n(int32(Ace) + 1))
}

//RandColor : Returns a random color
func RandColor() Color {
	return Color(rand.Int31n(int32(Diamond + 1)))
}

//NewCard : Returns an object Card
func NewCard(f Face, c Color) (ret Card, err error) {
	if !f.IsValid() {
		return ret, fmt.Errorf("Face %v is not valid", f)
	}
	if !c.IsValid() {
		return ret, fmt.Errorf("Color %v is not valid", c)
	}
	ret = Card{
		face:  f,
		color: c,
	}
	return ret, nil
}

//NewCardRandFace : Returns a new card with a random face
func NewCardRandFace(c Color) (Card, error) {
	f := RandFace()
	return NewCard(f, c)
}

//NewCardRandColor : Returns a new card with a random color
func NewCardRandColor(f Face) (Card, error) {
	c := RandColor()
	return NewCard(f, c)
}

//NewCardRandom : Returns a new random card
func NewCardRandom() Card {
	f, c := RandFace(), RandColor()
	ret, _ := NewCard(f, c)
	return ret
}

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

//NewGame : Initializes a new game
func NewGame() game {
	var ret game
	for i := 0; i < len(ret); i++ {
		ret[i] = NewDeck()
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

//Shuffle : Shuffle a game
func (g *game) Shuffle() {
	for i := 0; i < len(g); i++ {
		j := rand.Intn(i + 1)
		g[i].Shuffle()
		g[j].Shuffle()
		g[i], g[j] = g[j], g[i]
	}
}
