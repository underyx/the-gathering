package main

import (
	"fmt"
	"os"

	"github.com/adlio/trello"
	"github.com/underyx/the-gathering/giantbomb"
)

type GameMatch struct {
	Card *trello.Card
	Game *giantbomb.GameType
}

func main() {
	trelloKey := os.Getenv("THEG_TRELLO_KEY")
	trelloToken := os.Getenv("THEG_TRELLO_TOKEN")
	trelloBoardID := os.Getenv("THEG_TRELLO_BOARD_ID")
	giantBombKey := os.Getenv("THEG_GIANTBOMB_KEY")

	trelloClient := trello.NewClient(trelloKey, trelloToken)
	board, err := trelloClient.GetBoard(trelloBoardID, trello.Defaults())
	if err != nil {
		panic(err)
	}

	cards, err := board.GetCards(map[string]string{"fields": "name"})
	if err != nil {
		panic(err)
	}

	giantbombClient := giantbomb.NewClient(giantBombKey)

	var result []GameMatch
	for _, card := range cards {
		game, err := giantbombClient.Search(card.Name)
		if err != nil {
			panic(err)
		}
		result = append(result, GameMatch{card, game})
		fmt.Printf("Matched %s to %s (%v)\n", card.Name, game.Name, game.Platforms)
	}
}
