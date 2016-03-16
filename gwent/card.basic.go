package gwent

// BasicCard defines basic "blank" methods and fields for all cards
type BasicCard struct {
	CardName        string `json:"name"`
	CardDescription string `json:"desc"`
	ID              GUID   `json:"guid"`
}

// Play plays card by given player on given "other" card
func (c BasicCard) Play(p *Player, target Card) {
}

// PlayOnRow plays card by given player on given row
func (c BasicCard) PlayOnRow(p *Player, row CardRange) {
}

// PutOnTable adds card to game table
func (c BasicCard) PutOnTable(p *Player) {
}

// SetGUID sets ID of the card
func (c BasicCard) SetGUID(id GUID) {
	c.ID = id
}

// GUID reports back ID of the card
func (c BasicCard) GUID() GUID {
	return c.ID
}

// Name reports name of the card
func (c BasicCard) Name() string {
	return c.CardName
}

// Description reports description of the card
func (c BasicCard) Description() string {
	return c.CardDescription
}

// Faction reports faction of the card
func (c BasicCard) Faction() CardFaction {
	return FactionNeutral
}

// Type reports type of the card
func (c BasicCard) Type() CardType {
	return TypeBasic
}

// Range reports range/row of the card
func (c BasicCard) Range() CardRange {
	return RangeNone
}

// Power reports power of the card for given player
func (c BasicCard) Power(*Player) int {
	return 0
}

// Hero reports if card is a hero card
func (c BasicCard) Hero() bool {
	return false
}

// AppliedOnRow reports if card can be applied to some row
func (c BasicCard) AppliedOnRow() bool {
	return false
}

// Targettable reports if card can be targetted to some card
func (c BasicCard) Targettable() bool {
	return false
}
