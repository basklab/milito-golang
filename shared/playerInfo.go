package shared

type HandDTO struct {
	Cards []CardDTO
}

type CardDTO struct {
	Id       int
	UnitType string
}

type PlayerInfoDTO struct {
	Hand     HandDTO
	PlayerId int
	Faction  FactionsEnum
	Row1     []CardDTO
	Row2     []CardDTO
}

type GameTableDTO struct {
	Neutral       [5]int
	Phase         PhasesEnum
	CurrentPlayer PlayerInfoDTO
	AnotherPlayer PlayerInfoDTO
}
