package entity

import "milito-golang/shared"

type GameState struct {
	Neutral       [5]int
	Phase         shared.PhasesEnum
	ActivePlayer  PlayerState
	AnotherPlayer PlayerState
	BattleState   AttackState
	Step          StepsEnum
	DeployPenalty int
}

type AttackState struct {
	AttackColumn int
	DefenceColumn int
	SquadFormation SquadFormation
}

type SquadFormation struct {
	BonusCard   Card
	MainUnit    UnitCard
	SupportUnit UnitCard
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
	State       shared.PhasesEnum
}
