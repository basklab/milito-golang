package entities

type Card struct {
	Id int
}

type UnitCard struct {
	Card
	UnitType         string
	Speed            int
	AttackStrength   int
	DefenceStrength  int
	DeployPenalty    int
	FlankPenalty     int
	CombineAbility   bool
	AttackModifiers  string
	DefenceModifiers string
}

type LeaderCard struct {
	Card
	CombatValue      int
	PlaceUnitAbility int
	SpecialEffect    int
}
