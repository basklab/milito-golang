package config

import "milito-golang/entities"

func InitialSetup() entities.GameState {

	currentPlayer := entities.PlayerState{
		Deck:        entities.Deck{},
		DeadPile:    entities.Deck{},
		DiscardPile: entities.Deck{},
		Faction:     "",
		Hand:        entities.Hand{},
		PlayerId:    0,
		Row1:        [5]entities.UnitCard{},
		Row2:        [5]entities.UnitCard{},
		State:       "",
	}

	anotherPlayer := entities.PlayerState{
		Deck:        entities.Deck{},
		DeadPile:    entities.Deck{},
		DiscardPile: entities.Deck{},
		Faction:     "",
		Hand:        entities.Hand{},
		PlayerId:    0,
		Row1:        [5]entities.UnitCard{},
		Row2:        [5]entities.UnitCard{},
		State:       "",
	}

	return entities.GameState{
		Neutral:       [5]int{0, 0, 0, 0, 0},
		Phase:         entities.PhasesEnum("OLOLO"),
		CurrentPlayer: currentPlayer,
		AnotherPlayer: anotherPlayer,
		BattleState:   entities.AttackState{},
	}
}
