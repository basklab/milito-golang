package config

import "milito-golang/entity"

func InitialSetup() entity.GameState {

	currentPlayer := entity.PlayerState{
		Deck:        entity.Deck{},
		DeadPile:    entity.Deck{},
		DiscardPile: entity.Deck{},
		Faction:     "",
		Hand:        entity.Hand{},
		PlayerId:    0,
		Row1:        [5]entity.UnitCard{},
		Row2:        [5]entity.UnitCard{},
		State:       "",
	}

	anotherPlayer := entity.PlayerState{
		Deck:        entity.Deck{},
		DeadPile:    entity.Deck{},
		DiscardPile: entity.Deck{},
		Faction:     "",
		Hand:        entity.Hand{},
		PlayerId:    0,
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
