package entity

type Card interface {
	SetId(id int)
}

type UnitCard struct {
	Card
	Id int
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

func (c UnitCard) SetId(id int) {
	c.Id = id
}

type LeaderCard struct {
	Card
	Id int
	CombatValue      int
	PlaceUnitAbility int
	SpecialEffect    string
}

func (c LeaderCard) SetId(id int) {
	c.Id = id
}
