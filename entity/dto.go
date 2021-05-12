package entity

import "milito-golang/shared"

func (s PlayerState) ToDto() shared.PlayerInfoDTO {
	return shared.PlayerInfoDTO{
		Hand:     s.Hand.ToDto(),
		PlayerId: 0,
		Faction:  "",
		Row1:     nil,
		Row2:     nil,
	}
}

func (h Hand) ToDto() shared.HandDTO {
	return shared.HandDTO{
		Cards: nil,
	}
}

func (gs GameState) ToDto() shared.GameTableDTO {
	return shared.GameTableDTO{
		Neutral:       gs.Neutral,
		Phase:         "",
		CurrentPlayer: gs.ActivePlayer.ToDto(),
		AnotherPlayer: gs.AnotherPlayer.ToDto(),
	}
}
