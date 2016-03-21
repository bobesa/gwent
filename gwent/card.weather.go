package gwent

// CardWeather sets power of all non-hero unit cards to 1
type CardWeather struct {
	Target CardRange

	BasicCard
}

// Play puts card on the table
func (c *CardWeather) Play(p *Player, target Card) {
	c.PutOnTable(p)
}

// PutOnTable puts card on the table
func (c *CardWeather) PutOnTable(p *Player) {
	//Set weather card to proper row
	switch c.Target {
	case RangeClose:
		p.Game.WeatherClose = true
	case RangeRanged:
		p.Game.WeatherRanged = true
	case RangeSiege:
		p.Game.WeatherSiege = true
	case RangeNone:
		p.Game.ClearWeather()
	}
}

// Type reports that this is a weather card
func (c *CardWeather) Type() CardType {
	return TypeWeather
}

// Range reports what range weather card has
func (c *CardWeather) Range() CardRange {
	return c.Target
}
