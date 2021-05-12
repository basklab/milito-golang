package main

import (
	"encoding/json"
	"fmt"
	"milito-golang/game"
	"milito-golang/shared"
	"os"
)

func main() {
	//tmp := config.LoadDeckConfig(shared.AncientBritish)

	state := game.InitialSetup()
	b, _ := json.Marshal(state)
	_, _ = os.Stdout.Write(b)

	state = game.PlayACard(state, shared.PlayCardEvent{
		SelectedCard:   9,
		SelectedColumn: 3,
		SelectedRow:    1,
	})

	state = game.DiscardCardFromHand(state, shared.DiscardCardEvent{
		DiscardedCard: 7,
	})

	state = game.DiscardCardFromHand(state, shared.DiscardCardEvent{
		DiscardedCard: 11,
	})

	fmt.Println("STATE:")
	fmt.Println(state)
	fmt.Println("-------")

}
