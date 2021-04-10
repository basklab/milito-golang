package config

import (
	"milito-golang/entity"
)

func NewAlexandrianMacedonian() entity.Deck {
	units := makeUnitsAM()
	leaders := makeLeadersAM()
	return initializeDeck(units, leaders)
}

func makeUnitsAM() []entity.UnitCard {
	var units []entity.UnitCard
	units = append(units, multiply(7, pikes)...)
	units = append(units, multiply(4, spears)...)
	units = append(units, multiply(3, lightInfantry)...)
	units = append(units, multiply(2, slingers)...)
	units = append(units, multiply(2, companions)...)
	units = append(units, multiply(2, heavyCavalry)...)
	units = append(units, multiply(2, lightCavalry)...)
	return units
}

func makeLeadersAM() []entity.LeaderCard {
	return []entity.LeaderCard{
		{
			CombatValue:      3,
			PlaceUnitAbility: 3,
		},
		{
			CombatValue:      2,
			PlaceUnitAbility: 3,
		},
		{
			CombatValue:      2,
			PlaceUnitAbility: 2,
		},
		{
			CombatValue:      2,
			PlaceUnitAbility: 2,
			SpecialEffect:    "if combat is tied...",
		},
		{
			CombatValue:      2,
			PlaceUnitAbility: 2,
			SpecialEffect:    "+1 if attacking",
		},
		{
			CombatValue:      1,
			PlaceUnitAbility: 2,
			SpecialEffect:    "+2 if played with Pikes",
		},
		{
			CombatValue:      1,
			PlaceUnitAbility: 2,
			SpecialEffect:    "+2 if played with Compamnions",
		},
		{
			CombatValue:      -1,
			PlaceUnitAbility: 1,
		},
	}
}

var pikes = entity.UnitCard{
	UnitType:        "Pikes",
	Speed:           0,
	AttackStrength:  5,
	DefenceStrength: 5,
	DeployPenalty:   2,
	FlankPenalty:    1,
}

var spears = entity.UnitCard{
	UnitType:        "Spears",
	Speed:           0,
	AttackStrength:  4,
	DefenceStrength: 5,
	DeployPenalty:   1,
	FlankPenalty:    1,
}

var lightInfantry = entity.UnitCard{
	UnitType:        "Light_Infantry",
	Speed:           3,
	AttackStrength:  2,
	DefenceStrength: 2,
	CombineAbility:  true,
}

var companions = entity.UnitCard{
	UnitType:        "Companions",
	Speed:           3,
	AttackStrength:  4,
	DefenceStrength: 2,
}

var heavyCavalry = entity.UnitCard{
	UnitType:        "Heavy_Cavalry",
	Speed:           3,
	AttackStrength:  4,
	DefenceStrength: 3,
	DeployPenalty:   1,
	FlankPenalty:    1,
}
