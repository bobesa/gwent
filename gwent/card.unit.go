package gwent

const (
	ABILITY_NONE = iota
)

type CardUnit struct {
	Name, Description string
	Type, Power, Range, Ability, Faction int
	Hero bool
	Guid GUID
}

func (c *CardUnit) Play(p *Player, target Card) {
	c.PutOnTable(p)
}

func (c *CardUnit) PlayOnRow(p *Player, row int) {	
}

func (c *CardUnit) PutOnTable(p *Player) {
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

func (c *CardUnit) SetGUID(g GUID) {
	c.Guid = g
}
	
func (c *CardUnit) GetGUID() GUID {
	return c.Guid
}

func (c *CardUnit) GetName() string {
	return c.Name
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
		(c.GetRange() == RANGE_CLOSE && p.Game.WeatherClose) ||
		(c.GetRange() == RANGE_RANGED && p.Game.WeatherRanged) ||
		(c.GetRange() == RANGE_SIEGE && p.Game.WeatherSiege) ) {
		pwr = 1
	}
	
	//Apply horn if available
	if (c.GetRange() == RANGE_CLOSE && p.HornClose) ||
		(c.GetRange() == RANGE_RANGED && p.HornRanged) ||
		(c.GetRange() == RANGE_SIEGE && p.HornSiege) {
		pwr *= 2
	}
	
	return pwr
}

func (c *CardUnit) IsHero() bool {
	return c.Hero	
}

func (c *CardUnit) IsAppliedOnRow() bool {
	return false
}

func (c *CardUnit) IsTargetable() bool {
	//TODO: Treat by unit ability
	return false
}