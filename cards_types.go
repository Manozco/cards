package cards

import "fmt"

//Face represents the Face of a Card, e.g: A, K, Q, V, 10, 9, ...
type Face int

//Color is an int
type Color int

//Value represents the number of points of the cards (for the blackjack)
type Value int

//Card ...
type Card struct {
	color Color
	face  Face
}

// No joker for the moment
type deck [52]Card

type game [8]deck

// All The faces possible
const (
	Two   Face = iota
	Three Face = iota
	Four  Face = iota
	Five  Face = iota
	Six   Face = iota
	Seven Face = iota
	Eight Face = iota
	Nine  Face = iota
	Ten   Face = iota
	Jack  Face = iota
	Queen Face = iota
	King  Face = iota
	Ace   Face = iota
)

//Enum for our colors
const (
	Spade   Color = iota
	Club    Color = iota
	Heart   Color = iota
	Diamond Color = iota
)

var values = [...]Value{
	11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
}

//FaceValue: Mapping between faces and values
var FaceValue = map[Face][]Value{
	Ace:   {1, 11},
	King:  {10},
	Queen: {10},
	Jack:  {10},
	Ten:   {10},
	Nine:  {9},
	Eight: {8},
	Seven: {7},
	Six:   {6},
	Five:  {5},
	Four:  {4},

	Three: {3},
	Two:   {2},
}

func (c Card) String() string {
	return fmt.Sprintf("%v/%v", c.face, c.color)
}
