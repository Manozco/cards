package cards

import "math/rand"

type game [8]deck

//NewGame : Initializes a new game
func NewGame() game {
	var ret game
	for i := 0; i < len(ret); i++ {
		ret[i] = NewDeck()
	}
	return ret
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
