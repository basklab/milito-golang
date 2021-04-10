package config

import (
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v3"
	"math/rand"
	"milito-golang/entity"
)

type YamlConfig struct {
	Leaders []entity.LeaderCard `yaml:"leaders"`
	Units []entity.UnitCard `yaml:"units"`
	Counts map[string]int `yaml:"counts"`
}

//go:embed milito-config/decks/AncientBritish.yaml
var s []byte

func Ololo() {
	var yamlConfig YamlConfig
	_ = yaml.Unmarshal(s, &yamlConfig)
	fmt.Println(yamlConfig.Leaders)
	fmt.Println(yamlConfig.Units)
	fmt.Println(yamlConfig.Counts)
}


func multiply(nCards int, card entity.UnitCard) []entity.UnitCard {
	result := make([]entity.UnitCard, nCards)
	for i := range result {
		result[i] = card
	}
	return result
}

func initializeDeck(units []entity.UnitCard, leaders []entity.LeaderCard) entity.Deck {
	cards := make([]entity.Card, len(units)+len(leaders))
	for id, card := range leaders {
		cards[id] = &card
	}
	for id, card := range units {
		cards[len(leaders)+id] = &card
	}
	for i, _ := range cards {
		cards[i].SetId(i)
	}

	rand.Seed(1)
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return entity.Deck{Cards: cards}
}

var slingers = entity.UnitCard{
	UnitType:        "Slingers",
	Speed:           3,
	AttackStrength:  1,
	DefenceStrength: 1,
	CombineAbility:  true,
}

var lightCavalry = entity.UnitCard{
	UnitType:        "Light_Cavalry",
	Speed:           5,
	AttackStrength:  1,
	DefenceStrength: 1,
}
