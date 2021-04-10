package entity

type Card interface {
	SetId(idx int)
	Id() int
}

type UnitCard struct {
	id int
	UnitType         string `yaml:"unit_type"`
	Speed            int `yaml:"speed"`
	AttackStrength   int `yaml:"attack_strength"`
	DefenceStrength  int `yaml:"defence_strength"`
	DeployPenalty    int `yaml:"deploy_penalty"`
	FlankPenalty     int `yaml:"flank_penalty"`
	CombineAbility   bool `yaml:"combine_ability"`
	AttackModifiers  string `yaml:"attack_modifiers"`
	DefenceModifiers string `yaml:"defence_modifiers"`
}

func (c *UnitCard) SetId(idx int) {
	c.id = idx
}

func (c *UnitCard) Id() int {
	return c.id
}

type LeaderCard struct {
	id int
	CombatValue      int `yaml:"combat_value"`
	PlaceUnitAbility int `yaml:"place_unit_ability"`
	SpecialEffect    string `yaml:"special_effect"`
}

func (c *LeaderCard) SetId(idx int) {
	c.id = idx
}

func (c *LeaderCard) Id() int {
	return c.id
}
