package cards

import "math/rand"

//Face represents the Face of a Card, e.g: A, K, Q, V, 10, 9, ...
type Face int

//Value represents the number of points of the cards (for the blackjack)
type Value int

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

//IsValid checks the validity of a Face
func (f Face) IsValid() bool {
	return (f >= Two && f <= Ace)
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

//RandFace : Returns a random face
func RandFace() Face {
	return Face(rand.Int31n(int32(Ace) + 1))
}
