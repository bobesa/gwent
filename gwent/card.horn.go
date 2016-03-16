package gwent

type CardHorn struct {
	BasicCard
}

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

func (c *CardHorn) Type() CardType {
	return TypeHorn
}

func (c *CardHorn) AppliedOnRow() bool {
	return true
}
