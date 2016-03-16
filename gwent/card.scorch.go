package gwent

type CardScorch struct {
	Type int

	BasicCard
}

func (c *CardScorch) Play(p *Player, target Card) {
	p.Grave = append(p.Grave, c)	
	p.Scorch()
}

func (c *CardScorch) GetType() int {
	return TypeScorch
}