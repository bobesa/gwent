package gwent

type CardHorn struct {
	Name, Description string
	Guid GUID
}

func (c *CardHorn) Play(p *Player, target Card) {
}

func (c *CardHorn) PlayOnRow(p *Player, row int) {	
	//Apply horn
	switch(row){
	case RANGE_CLOSE:
		p.HornClose = true
	case RANGE_RANGED:
		p.HornRanged = true
	case RANGE_SIEGE:
		p.HornSiege = true
	}
}

func (c *CardHorn) PutOnTable(p *Player) {
}

func (c *CardHorn) SetGUID(g GUID) {
	c.Guid = g
}
	
func (c *CardHorn) GetGUID() GUID {
	return c.Guid
}

func (c *CardHorn) GetName() string {
	return c.Name
}

func (c *CardHorn) GetFaction() int {
	return FACTION_NEUTRAL
}

func (c *CardHorn) GetType() int {
	return TYPE_HORN
}

func (c *CardHorn) GetRange() int {
	return RANGE_NONE
}

func (c *CardHorn) GetPower(*Player) int {
	return 0
}

func (c *CardHorn) IsHero() bool {
	return false	
}

func (c *CardHorn) IsAppliedOnRow() bool {
	return true
}

func (c *CardHorn) IsTargetable() bool {
	return false
}