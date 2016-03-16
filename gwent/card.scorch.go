package gwent

type CardScorch struct {
	BasicCard
}

func (c *CardScorch) Play(p *Player, target Card) {
	p.Grave = append(p.Grave, c)
	p.Scorch()
}

func (c *CardScorch) Type() CardType {
	return TypeScorch
}
