package gwent

const (
	AbilityNone = iota
)

type CardUnit struct {
	UnitType    CardType
	UnitRange   CardRange
	UnitFaction CardFaction

	UnitPower, UnitAbility int
	UnitHero           bool

	BasicCard
}

func (c *CardUnit) Play(p *Player, target Card) {
	c.PutOnTable(p)
}

func (c *CardUnit) PutOnTable(p *Player) {
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

func (c *CardUnit) Type() CardType {
	return c.UnitType
}

func (c *CardUnit) Faction() CardFaction {
	return c.UnitFaction
}

func (c *CardUnit) Range() CardRange {
	return c.UnitRange
}

func (c *CardUnit) Power(p *Player) int {
	pwr := c.UnitPower

	//Apply weather if not a hero card
	if !c.Hero() && ((c.Range() == RangeClose && p.Game.WeatherClose) ||
		(c.Range() == RangeRanged && p.Game.WeatherRanged) ||
		(c.Range() == RangeSiege && p.Game.WeatherSiege)) {
		pwr = 1
	}

	//Apply horn if available
	if (c.Range() == RangeClose && p.HornClose) ||
		(c.Range() == RangeRanged && p.HornRanged) ||
		(c.Range() == RangeSiege && p.HornSiege) {
		pwr *= 2
	}

	return pwr
}

func (c *CardUnit) Hero() bool {
	return c.UnitHero
}

func (c *CardUnit) Targetable() bool {
	//TODO: Treat by unit ability
	return false
}
