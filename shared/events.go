package shared

type PlayCardEvent struct {
	SelectedCard   int
	SelectedColumn int
	SelectedRow    int
}

type DiscardCardEvent struct {
	DiscardedCard int
}

type InitiateAttackEvent struct {
	BonusCardId    int
	SelectedColumn int
	SelectedRow    int
}
