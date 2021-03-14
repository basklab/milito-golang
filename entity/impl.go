package entity

func (h Hand) TakeCardById(id int) Card {
	//var result Card
	//cards := make([]Card, 0, len(h.Cards) - 1)
	for i, v := range h.Cards {
		if v.Id() == id {
			h.Cards = append(h.Cards[:i], h.Cards[i+1:]...)
			return v
		}
	}
	panic("TakeCardById not found")
}

func (d *Deck) Push(card Card) {
	d.Cards = append(d.Cards, card)
}

func (d *Deck) Pop() Card {
	idx := len(d.Cards) - 1
	result := d.Cards[idx]
	d.Cards = d.Cards[:idx]
	return result
}
