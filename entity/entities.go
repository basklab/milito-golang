package entity

import "milito-golang/shared"

type GameState struct {
	Neutral       [5]int
	Phase         PhasesEnum
	CurrentPlayer PlayerState
	AnotherPlayer PlayerState
	BattleState   AttackState
}

type AttackState struct {
}

type SquadFormation struct {
	MainUnit    UnitCard
	SupportUnit UnitCard
	BonusCard   Card
}

type Deck struct {
	Cards []Card
}

type Hand struct {
	Deck
}

type PlayerState struct {
	Deck        Deck
	DeadPile    Deck
	DiscardPile Deck
	Faction     shared.FactionsEnum
	Hand        Hand
	PlayerId    int
	Row1        [5]UnitCard
	Row2        [5]UnitCard
	State       PhasesEnum
}
