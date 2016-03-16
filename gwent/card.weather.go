package gwent

type CardWeather struct {
	Target CardRange

	BasicCard
}

func (c *CardWeather) Play(p *Player, target Card) {
	c.PutOnTable(p)
}

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

func (c *CardWeather) GetType() CardType {
	return TypeWeather
}

func (c *CardWeather) GetRange() CardRange {
	return c.Target
}
