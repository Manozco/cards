package cards

import "math/rand"

//Color is an int
type Color uint8

//Enum for our colors
const (
	Spade   Color = iota
	Club    Color = iota
	Heart   Color = iota
	Diamond Color = iota
)

//IsValid returns true if a color is valid
func (c Color) IsValid() bool {
	return (c >= Spade && c <= Diamond)
}

//RandColor : Returns a random color
func RandColor() Color {
	return Color(rand.Int31n(int32(Diamond + 1)))
}
