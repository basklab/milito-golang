package config

import (
	"milito-golang/config/decks"
	"milito-golang/entity"
	"milito-golang/shared"
)

func InitialSetup() entity.GameState {

	currentPlayer := entity.PlayerState{
		Deck:        decks.NewAlexandrianMacedonian(),
		DeadPile:    entity.Deck{},
		DiscardPile: entity.Deck{},
		Faction:     shared.AlexandrianMacedonian,
		Hand:        entity.Hand{},
		PlayerId:    0,
		Row1:        [5]entity.UnitCard{},
		Row2:        [5]entity.UnitCard{},
		State:       "",
	}

	anotherPlayer := entity.PlayerState{
		Deck:        decks.NewAncientBritish(),
		DeadPile:    entity.Deck{},
		DiscardPile: entity.Deck{},
		Faction:     shared.AncientBritish,
		Hand:        entity.Hand{},
		PlayerId:    1,
		Row1:        [5]entity.UnitCard{},
		Row2:        [5]entity.UnitCard{},
		State:       "",
	}

	return entity.GameState{
		Neutral:       [5]int{0, 0, 0, 0, 0},
		Phase:         entity.PhasesEnum("OLOLO"),
		CurrentPlayer: currentPlayer,
		AnotherPlayer: anotherPlayer,
		BattleState:   entity.AttackState{},
	}
}
