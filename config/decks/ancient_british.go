package decks

import (
	"milito-golang/entity"
)

func NewAncientBritish() entity.Deck {
	units := makeUnitsAB()
	leaders := makeLeadersAB()
	return initializeDeck(units, leaders)
}

func makeUnitsAB() []entity.UnitCard {
	var units []entity.UnitCard
	units = append(units, multiply(6, warbandMediumInfantry)...)
	units = append(units, multiply(6, slingers)...)
	units = append(units, multiply(6, chariots)...)
	units = append(units, multiply(4, lightInfantry)...)
	return units
}

func makeLeadersAB() []entity.LeaderCard {
	return []entity.LeaderCard{
		{
			CombatValue:      3,
			PlaceUnitAbility: 3,
		},
		{
			CombatValue:      2,
			PlaceUnitAbility: 2,
		},
		{
			CombatValue:      2,
			PlaceUnitAbility: 1,
			SpecialEffect:    "if win combat",
		},
		{
			CombatValue:      1,
			PlaceUnitAbility: 2,
			SpecialEffect:    "+2 if played in Wood or Rough column.",
		},
		{
			CombatValue:      1,
			PlaceUnitAbility: 1,
			SpecialEffect:    "+2 if played with 2x Chariots",
		},
		{
			CombatValue:      1,
			PlaceUnitAbility: 1,
			SpecialEffect:    "+2 if played with 2x Slingers",
		},
		{
			CombatValue:      1,
			PlaceUnitAbility: 1,
			SpecialEffect:    "+2 if played with 2x Warband Medium Infantry",
		},
		{
			CombatValue:      -1,
			PlaceUnitAbility: 1,
		},
	}
}

var warbandMediumInfantry = entity.UnitCard{
	UnitType:        "Warband_Med_Inf",
	Speed:           2,
	AttackStrength:  4,
	DefenceStrength: 2,
}

var chariots = entity.UnitCard{
	UnitType:        "Chariots",
	Speed:           3,
	AttackStrength:  4,
	DefenceStrength: 2,
}
