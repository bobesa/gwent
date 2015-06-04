package gwent

type CardScorch struct {
	Name, Description string
	Type int
	Guid GUID
}

func (c *CardScorch) Play(p *Player, target Card) {
	p.Grave = append(p.Grave, c)	
	p.Scorch()
}

func (c *CardScorch) PlayOnRow(p *Player, row int) {	
}

func (c *CardScorch) PutOnTable(p *Player) {
}

func (c *CardScorch) SetGUID(g GUID) {
	c.Guid = g
}
	
func (c *CardScorch) GetGUID() GUID {
	return c.Guid
}

func (c *CardScorch) GetName() string {
	return c.Name
}

func (c *CardScorch) GetType() int {
	return TYPE_SCORCH
}

func (c *CardScorch) GetFaction() int {
	return FACTION_NEUTRAL
}

func (c *CardScorch) GetRange() int {
	return RANGE_NONE
}

func (c *CardScorch) GetPower(*Player) int {
	return 0
}

func (c *CardScorch) IsHero() bool {
	return false	
}

func (c *CardScorch) IsAppliedOnRow() bool {
	return false
}

func (c *CardScorch) IsTargetable() bool {
	return false
}