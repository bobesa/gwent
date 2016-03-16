package gwent

// CardScorch destroys cards with most power
type CardScorch struct {
	BasicCard
}

// Play put Scorch card to grave and plays Scorch effect
func (c *CardScorch) Play(p *Player, target Card) {
	p.Grave = append(p.Grave, c)
	p.Scorch()
}

// Type reports that this is a scorch card
func (c *CardScorch) Type() CardType {
	return TypeScorch
}
