package gwent

// CardHorn boosts given row cards power
type CardHorn struct {
	BasicCard
}

// PlayOnRow adds Horn effect on given row of given player
func (c *CardHorn) PlayOnRow(p *Player, row CardRange) {
	//Apply horn
	switch row {
	case RangeClose:
		p.HornClose = true
	case RangeRanged:
		p.HornRanged = true
	case RangeSiege:
		p.HornSiege = true
	}
}

// Type reports that this is Horn card
func (c *CardHorn) Type() CardType {
	return TypeHorn
}

// AppliedOnRow reports that we need to select row
func (c *CardHorn) AppliedOnRow() bool {
	return true
}
