package gwent

const (
	// AbilityNone means unit card has no ability at all
	AbilityNone = iota
)

// CardUnit is single unit used for combat
type CardUnit struct {
	UnitType    CardType
	UnitRange   CardRange
	UnitFaction CardFaction

	UnitPower, UnitAbility int
	UnitHero               bool

	BasicCard
}

// Play puts unit card to table
func (c *CardUnit) Play(p *Player, target Card) {
	c.PutOnTable(p)
}

// PutOnTable puts unit card to table
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

// Type reports type of this unit card
func (c *CardUnit) Type() CardType {
	return c.UnitType
}

// Faction reports faction of this unit card
func (c *CardUnit) Faction() CardFaction {
	return c.UnitFaction
}

// Range reports range of this unit card
func (c *CardUnit) Range() CardRange {
	return c.UnitRange
}

// Power reports power of this unit card
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

// Hero reports if this unit card is hero card
func (c *CardUnit) Hero() bool {
	return c.UnitHero
}

// Targettable reports if this card can be targetted or not
func (c *CardUnit) Targettable() bool {
	//TODO: Treat by unit ability
	return false
}
