package main

import (
	"encoding/json"
	"fmt"
	"milito-golang/config"
	"milito-golang/game"
	"milito-golang/shared"
	"os"
)

func main() {
	tmp := config.LoadDeckConfig(shared.AncientBritish)
	b, _ := json.Marshal(tmp)
	_, _ = os.Stdout.Write(b)
	return

	state := game.InitialSetup()

	fmt.Println("STATE:")
	fmt.Println(state)
	fmt.Println("-------")

	state = game.PlaceAUnit(state, shared.PlaceUnitEvent{
		DiscardedCards: []int{7, 21},
		SelectedCard:   10,
		SelectedColumn: 3,
		SelectedRow:    1,
	})

	state = game.PlaceAUnit(state, shared.PlaceUnitEvent{
		DiscardedCards: []int{},
		SelectedCard:   29,
		SelectedColumn: 4,
		SelectedRow:    1,
	})

}
