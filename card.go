package cards

import "fmt"

//Card ...
type Card struct {
	color Color
	face  Face
}

func (c Card) String() string {
	return fmt.Sprintf("%v/%v", c.face, c.color)
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
