package game

import (
	"milito-golang/entity"
	"milito-golang/shared"
	"reflect"
)

func PlayACard(state entity.GameState, action shared.PlayCardEvent) entity.GameState {
	println("ACTION: ", reflect.TypeOf(action).Name())
	if state.Step != entity.PLAY_CARD || state.Phase != shared.PHASE_4_PLAYER_ACTIONS {
		panic("bad step or phase")
	}
	player := state.ActivePlayer
	newCard := player.Hand.TakeCardById(action.SelectedCard).(*entity.UnitCard)
	state.DeployPenalty = newCard.DeployPenalty
	if action.SelectedColumn == 0 || action.SelectedColumn == 4 {
		state.DeployPenalty += newCard.FlankPenalty
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
	if state.DeployPenalty == 0 {
		state.Step = entity.PLAY_CARD
	} else {
		state.Step = entity.DISCARD_DEPLOY_COST
	}
	return state
}

func DiscardCardFromHand(state entity.GameState, action shared.DiscardCardEvent) entity.GameState {
	println("ACTION: ", reflect.TypeOf(action).Name())
	if state.Step != entity.DISCARD_DEPLOY_COST || state.Phase != shared.PHASE_4_PLAYER_ACTIONS {
		panic("bad step or phase")
	}
	player := state.ActivePlayer
	newCard := player.Hand.TakeCardById(action.DiscardedCard)
	switch leaderCard := newCard.(type) {
	case *entity.UnitCard:
		state.DeployPenalty -= 1
	case *entity.LeaderCard:
		state.DeployPenalty -= leaderCard.PlaceUnitAbility
	}
	if state.DeployPenalty <= 0 {
		state.DeployPenalty = 0
		state.Step = entity.PLAY_CARD
	} else {
		state.Step = entity.DISCARD_DEPLOY_COST
	}
	return state
}

func InitiateAttack(state entity.GameState, action shared.InitiateAttackEvent) entity.GameState {
	println("ACTION: ", reflect.TypeOf(action).Name())
	if state.Phase != shared.PHASE_4_PLAYER_ACTIONS {
		panic("bad step or phase")
	}
	var mainCard entity.UnitCard
	var supportCard entity.UnitCard
	if action.SelectedRow == 1 {
		mainCard = state.ActivePlayer.Row1[action.SelectedColumn]
		supportCard = state.ActivePlayer.Row2[action.SelectedColumn]
	} else {
		mainCard = state.ActivePlayer.Row2[action.SelectedColumn]
		supportCard = state.ActivePlayer.Row1[action.SelectedColumn]
	}
	if mainCard.Id() == 0 {
		panic("card not found")
	}
	var bonusCard entity.Card
	if action.BonusCardId == 0 {
		bonusCard = state.ActivePlayer.Deck.Pop()
	} else {
		bonusCard = state.ActivePlayer.Hand.TakeCardById(action.BonusCardId)
	}
	state.BattleState = entity.AttackState{
		AttackColumn:  action.SelectedColumn,
		DefenceColumn: action.SelectedColumn,
		SquadFormation: entity.SquadFormation{
			BonusCard:   bonusCard,
			MainUnit:    mainCard,
			SupportUnit: supportCard,
		},
	}
	state.Phase = shared.PHASE_4_DEFENDER_ACTIONS
	return state
}
