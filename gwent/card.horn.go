package gwent

type CardHorn struct {
	BasicCard
}

func (c *CardHorn) PlayOnRow(p *Player, row int) {	
	//Apply horn
	switch(row){
	case RangeClose:
		p.HornClose = true
	case RangeRanged:
		p.HornRanged = true
	case RangeSiege:
		p.HornSiege = true
	}
}

func (c *CardHorn) GetType() int {
	return TypeHorn
}

func (c *CardHorn) IsAppliedOnRow() bool {
	return true
}