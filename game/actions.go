package game

import (
	"milito-golang/entity"
	"milito-golang/shared"
	"reflect"
)

func PlaceAUnit(state entity.GameState, action shared.PlaceUnitEvent) entity.GameState {
	println("ACTION: ", reflect.TypeOf(action).Name())
	player := state.CurrentPlayer
	deployValue := 0
	for _, discardId := range action.DiscardedCards {
		discardCard := player.Hand.TakeCardById(discardId)
		switch c := discardCard.(type) {
		case *entity.LeaderCard:
			deployValue += c.PlaceUnitAbility
		case *entity.UnitCard:
			deployValue += 1
		}
		player.DiscardPile.Push(discardCard)
	}
	// play a new card
	newCard := player.Hand.TakeCardById(action.SelectedCard).(*entity.UnitCard)
	deployPenalty := 0
	if action.SelectedColumn == 0 || action.SelectedColumn == 4 {
		deployPenalty += newCard.DeployPenalty + newCard.FlankPenalty
	}
	if deployPenalty > deployValue {
		panic("deployPenalty > deployValue")
	}
	// todo check another unit must be the same type or must have combine ability
	var currentRow [5]entity.UnitCard
	if action.SelectedRow == 1 {
		currentRow = player.Row1
	} else {
		currentRow = player.Row2
	}
	tmp := currentRow[action.SelectedColumn]
	if tmp.UnitType != "" {
		player.DiscardPile.Push(&tmp)
	}
	currentRow[action.SelectedColumn] = *newCard
	return state
}
