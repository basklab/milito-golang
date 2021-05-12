package config

import (
	_ "embed"
	"gopkg.in/yaml.v3"
	"math/rand"
	"milito-golang/entity"
	"milito-golang/shared"
)

type YamlConfig struct {
	Leaders []entity.LeaderCard `yaml:"leaders"`
	Units   []entity.UnitCard   `yaml:"units"`
	Counts  map[string]int      `yaml:"counts"`
}

//go:embed milito-config/decks/AlexandrianMacedonian.yaml
var alexandrianMacedonianDeckConfig []byte

//go:embed milito-config/decks/AncientBritish.yaml
var ancientBritishDeckConfig []byte

func LoadDeckConfig(deck shared.FactionsEnum) entity.Deck {
	var yamlConfig YamlConfig
	switch deck {
	case shared.AlexandrianMacedonian:
		_ = yaml.Unmarshal(alexandrianMacedonianDeckConfig, &yamlConfig)
	case shared.AncientBritish:
		_ = yaml.Unmarshal(ancientBritishDeckConfig, &yamlConfig)
	}
	var cards []entity.Card
	for _, leaderCard := range yamlConfig.Leaders {
		cards = append(cards, &leaderCard)
	}
	for _, unitCard := range yamlConfig.Units {
		cards = append(cards, multiply(yamlConfig.Counts[unitCard.UnitType], unitCard)...)
	}
	// set card uuids
	for i, _ := range cards {
		cards[i].SetId(i)
	}
	// shuffle deck in deterministic order
	rand.Seed(42)
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return entity.Deck{Cards: cards}
}

func multiply(nCards int, card entity.UnitCard) []entity.Card {
	result := make([]entity.Card, nCards)
	for i := range result {
		tmp := card
		result[i] = &tmp
	}
	return result
}
