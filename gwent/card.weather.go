package gwent

type CardWeather struct {
	Name, Description string
	Target int
	Guid GUID
}

func (c *CardWeather) Play(p *Player, target Card) {
	c.PutOnTable(p)
}

func (c *CardWeather) PlayOnRow(p *Player, row int) {	
}

func (c *CardWeather) PutOnTable(p *Player) {
	//Set weather card to proper row
	switch(c.Target) {
		case RANGE_CLOSE:
			p.Game.WeatherClose = true
		case RANGE_RANGED:
			p.Game.WeatherRanged = true
		case RANGE_SIEGE:
			p.Game.WeatherSiege = true
		case RANGE_NONE:
			p.Game.ClearWeather()
	}
}

func (c *CardWeather) SetGUID(g GUID) {
	c.Guid = g
}
	
func (c *CardWeather) GetGUID() GUID {
	return c.Guid
}

func (c *CardWeather) GetName() string {
	return c.Name
}

func (c *CardWeather) GetType() int {
	return TYPE_WEATHER
}

func (c *CardWeather) GetFaction() int {
	return FACTION_NEUTRAL
}

func (c *CardWeather) GetRange() int {
	return c.Target
}

func (c *CardWeather) GetPower(*Player) int {
	return 0
}

func (c *CardWeather) IsHero() bool {
	return false	
}

func (c *CardWeather) IsAppliedOnRow() bool {
	return false
}

func (c *CardWeather) IsTargetable() bool {
	return false
}