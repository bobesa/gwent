package gwent

type CardDecoy struct {
	Row CardRange `json:"row"`

	BasicCard
}

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

func (c *CardDecoy) Type() CardType {
	return TypeHorn
}

func (c *CardDecoy) Range() CardRange {
	return c.Row
}

func (c *CardDecoy) Targettable() bool {
	return true
}
