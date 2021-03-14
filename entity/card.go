package entity

type Card interface {
	SetId(idx int)
	Id() int
}

type UnitCard struct {
	id int
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

func (c *UnitCard) SetId(idx int) {
	c.id = idx
}

func (c *UnitCard) Id() int {
	return c.id
}

type LeaderCard struct {
	id int
	CombatValue      int
	PlaceUnitAbility int
	SpecialEffect    string
}

func (c *LeaderCard) SetId(idx int) {
	c.id = idx
}

func (c *LeaderCard) Id() int {
	return c.id
}
