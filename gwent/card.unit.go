package gwent

const (
	AbilityNone = iota
)

type CardUnit struct {
	Type, Power, Range, Ability, Faction int
	Hero bool

	BasicCard
}

func (c *CardUnit) Play(p *Player, target Card) {
	c.PutOnTable(p)
}

func (c *CardUnit) PutOnTable(p *Player) {
	//Add card to proper row
	switch(c.GetRange()) {
		case RangeClose:
			p.RowClose = append(p.RowClose, c)
		case RangeRanged:
			p.RowRanged = append(p.RowRanged, c)
		case RangeSiege:
			p.RowSiege = append(p.RowSiege, c)
	}
}

func (c *CardUnit) GetType() int {
	return c.Type
}

func (c *CardUnit) GetFaction() int {
	return c.Faction
}

func (c *CardUnit) GetRange() int {
	return c.Range
}

func (c *CardUnit) GetPower(p *Player) int {
	pwr := c.Power	
		
	//Apply weather if not a hero card
	if !c.IsHero() && (
		(c.GetRange() == RangeClose && p.Game.WeatherClose) ||
		(c.GetRange() == RangeRanged && p.Game.WeatherRanged) ||
		(c.GetRange() == RangeSiege && p.Game.WeatherSiege) ) {
		pwr = 1
	}
	
	//Apply horn if available
	if (c.GetRange() == RangeClose && p.HornClose) ||
		(c.GetRange() == RangeRanged && p.HornRanged) ||
		(c.GetRange() == RangeSiege && p.HornSiege) {
		pwr *= 2
	}
	
	return pwr
}

func (c *CardUnit) IsHero() bool {
	return c.Hero	
}

func (c *CardUnit) IsTargetable() bool {
	//TODO: Treat by unit ability
	return false
}