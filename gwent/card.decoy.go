package gwent

// CardDecoy replaces any oponnents card
type CardDecoy struct {
	Row CardRange `json:"row"`

	BasicCard
}

// Play plays decoy card on given target
func (c *CardDecoy) Play(p *Player, target Card) {
	if target != nil {
		p.RowClose = p.RowClose.Without(target)
		p.RowRanged = p.RowRanged.Without(target)
		p.RowSiege = p.RowSiege.Without(target)
		p.GiveCard(target)
		c.Row = target.Range()
		c.PutOnTable(p)
	}
}

// PutOnTable puts decoy card onto the table
func (c *CardDecoy) PutOnTable(p *Player) {
	//Add card to proper row
	switch c.Range() {
	case RangeClose:
		p.RowClose = append(p.RowClose, c)
	case RangeRanged:
		p.RowRanged = append(p.RowRanged, c)
	case RangeSiege:
		p.RowSiege = append(p.RowSiege, c)
	}
}

// Range reports what row has been chosen
func (c *CardDecoy) Range() CardRange {
	return c.Row
}

// Targettable reports that decoy card has to be targetted on some other unit card
func (c *CardDecoy) Targettable() bool {
	return true
}
