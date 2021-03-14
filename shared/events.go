package shared

type PlaceUnitEvent struct {
	DiscardedCards []int
	SelectedCard   int
	SelectedColumn int
	SelectedRow    int
}
