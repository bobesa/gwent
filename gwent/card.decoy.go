package gwent

type CardDecoy struct {
	Name, Description string
	Row int
	Guid GUID
}

func (c *CardDecoy) Play(p *Player, target Card) {
	if target != nil {
		p.RowClose = p.RowClose.Without(target)
		p.RowRanged = p.RowRanged.Without(target)
		p.RowSiege = p.RowSiege.Without(target)
		p.GiveCard(target)
		c.Row = target.GetRange()
		c.PutOnTable(p)		
	}
}

func (c *CardDecoy) PlayOnRow(p *Player, row int) {	
}

func (c *CardDecoy) PutOnTable(p *Player) {
	//Add card to proper row
	switch(c.GetRange()) {
		case RANGE_CLOSE:
			p.RowClose = append(p.RowClose, c)
		case RANGE_RANGED:
			p.RowRanged = append(p.RowRanged, c)
		case RANGE_SIEGE:
			p.RowSiege = append(p.RowSiege, c)
	}
}

func (c *CardDecoy) SetGUID(g GUID) {
	c.Guid = g
}
	
func (c *CardDecoy) GetGUID() GUID {
	return c.Guid
}

func (c *CardDecoy) GetName() string {
	return c.Name
}

func (c *CardDecoy) GetFaction() int {
	return FACTION_NEUTRAL
}

func (c *CardDecoy) GetType() int {
	return TYPE_HORN
}

func (c *CardDecoy) GetRange() int {
	return c.Row
}

func (c *CardDecoy) GetPower(*Player) int {
	return 0
}

func (c *CardDecoy) IsHero() bool {
	return false	
}

func (c *CardDecoy) IsAppliedOnRow() bool {
	return false
}

func (c *CardDecoy) IsTargetable() bool {
	return true
}