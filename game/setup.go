package game

import (
	"milito-golang/config"
	"milito-golang/entity"
	"milito-golang/shared"
)

const NCardsAtGameStart = 9

func TakeCard(player entity.PlayerState) entity.PlayerState {
	card := player.Deck.Pop()
	player.Hand.Push(card)
	return player
}

func InitialSetup() entity.GameState {

	currentPlayer := entity.PlayerState{
		Deck:        config.LoadDeckConfig(shared.AlexandrianMacedonian),
		DeadPile:    entity.Deck{},
		DiscardPile: entity.Deck{},
		Faction:     shared.AlexandrianMacedonian,
		Hand:        entity.Hand{Deck: entity.Deck{Cards: []entity.Card{}}},
		PlayerId:    0,
		Row1:        [5]entity.UnitCard{},
		Row2:        [5]entity.UnitCard{},
		State:       "",
	}

	anotherPlayer := entity.PlayerState{
		Deck:        config.LoadDeckConfig(shared.AncientBritish),
		DeadPile:    entity.Deck{Cards: []entity.Card{}},
		DiscardPile: entity.Deck{Cards: []entity.Card{}},
		Faction:     shared.AncientBritish,
		Hand:        entity.Hand{Deck: entity.Deck{Cards: []entity.Card{}}},
		PlayerId:    1,
		Row1:        [5]entity.UnitCard{},
		Row2:        [5]entity.UnitCard{},
		State:       "",
	}

	for i := 0; i < NCardsAtGameStart; i++ {
		currentPlayer = TakeCard(currentPlayer)
		anotherPlayer = TakeCard(anotherPlayer)
	}

	return entity.GameState{
		Neutral:       [5]int{0, 0, 0, 0, 0},
		Phase:         shared.PHASE_4_PLAYER_ACTIONS,
		Step:          entity.PLAY_CARD,
		ActivePlayer:  currentPlayer,
		AnotherPlayer: anotherPlayer,
		BattleState:   entity.AttackState{},
	}

}
