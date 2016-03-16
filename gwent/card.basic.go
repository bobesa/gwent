package gwent

type BasicCard struct {
	Name string `json:"name"`
	Description string `json:"desc"`
	ID GUID `json:"guid"`
}

func (c BasicCard) Play(p *Player, target Card) {
}

func (c BasicCard) PlayOnRow(p *Player, row CardRange) {
}

func (c BasicCard) PutOnTable(p *Player) {
}

func (c BasicCard) SetGUID(id GUID) {
	c.ID = id
}

func (c BasicCard) GUID() GUID {
	return c.ID
}

func (c BasicCard) GetName() string {
	return c.Name
}

func (c BasicCard) GetFaction() CardFaction {
	return FactionNeutral
}

func (c BasicCard) GetType() CardType {
	return TypeBasic
}

func (c BasicCard) GetRange() CardRange {
	return RangeNone
}

func (c BasicCard) GetPower(*Player) int {
	return 0
}

func (c BasicCard) IsHero() bool {
	return false
}

func (c BasicCard) IsAppliedOnRow() bool {
	return false
}

func (c BasicCard) IsTargettable() bool {
	return false
}
