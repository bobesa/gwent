package gwent

// CardLeader is a ultimate card that each player has before the game starts
type CardLeader struct {
	CannotUse     bool
	LeaderEffect  int
	LeaderFaction CardFaction

	BasicCard
}

// PlayWeatherFromDeck plays weather card from players deck
func (c *CardLeader) PlayWeatherFromDeck(where CardRange, p *Player) bool {
	var card Card
	p.Deck, card = p.Deck.WithoutType(TypeWeather, where)
	if card != nil {
		card.Play(p, nil) //Cast card
		return true
	}
	return false
}

// Play plays leader effect on given card
func (c *CardLeader) Play(p *Player, target Card) {
	if !c.AppliedOnRow() {
		c.Cancel()

		//Process various effects (by Type)
		switch c.LeaderEffect {
		//Other
		case LeaderFxCancel:
			p.OtherPlayer().Leader.Cancel()
		case LeaderFxDrawExtraCard:
			p.DrawCard()
		case LeaderFxDrawGrave:
		case LeaderFxDrawOpponentGrave:
		case LeaderFxDiscard2Draw1:
		case LeaderFxLook3OpponentCards:
			//Todo: Actually solve this :)

		//Destroy
		case LeaderFxDestroyClose10Plus:
			if p.OtherPlayer().ComputePowerOfRow(RangeClose) >= 10 {
				maxPwr := 0
				for _, card := range p.OtherPlayer().RowClose {
					cardPower := card.Power(p.OtherPlayer())
					if cardPower > maxPwr {
						maxPwr = cardPower
					}
				}
				p.OtherPlayer().RowClose, _ = p.OtherPlayer().RowClose.Scorch(p.OtherPlayer(), maxPwr)
			}

		case LeaderFxDestroySiege10Plus:
			if p.OtherPlayer().ComputePowerOfRow(RangeSiege) >= 10 {
				maxPwr := 0
				for _, card := range p.OtherPlayer().RowSiege {
					cardPower := card.Power(p.OtherPlayer())
					if cardPower > maxPwr {
						maxPwr = cardPower
					}
				}
				p.OtherPlayer().RowSiege, _ = p.OtherPlayer().RowSiege.Scorch(p.OtherPlayer(), maxPwr)
			}

		//Weather
		case LeaderFxPlayWeatherClose:
			if !c.PlayWeatherFromDeck(RangeClose, p) {
				c.Ready()
			}
		case LeaderFxPlayWeatherRanged:
			if !c.PlayWeatherFromDeck(RangeRanged, p) {
				c.Ready()
			}
		case LeaderFxPlayWeatherSiege:
			if !c.PlayWeatherFromDeck(RangeSiege, p) {
				c.Ready()
			}
		case LeaderFxWeatherClear:
			p.Game.ClearWeather()

		//Horn
		case LeaderFxHornClose:
			p.HornClose = true
		case LeaderFxHornRanged:
			p.HornRanged = true
		case LeaderFxHornSiege:
			p.HornSiege = true
		}
	}
}

// PlayOnRow plays leader effect on given row
func (c *CardLeader) PlayOnRow(p *Player, row CardRange) {
	if c.AppliedOnRow() {
		c.Cancel()

		//Process various effects (by Type)
		switch c.LeaderEffect {
		case LeaderFxPlayWeatherAny:
			if !c.PlayWeatherFromDeck(row, p) {
				c.Ready()
			}
		}
	}
}

// Cancel marks leader fx as used
func (c *CardLeader) Cancel() {
	c.CannotUse = true
}

// Ready marks leader fx as not used
func (c *CardLeader) Ready() {
	c.CannotUse = false
}

// Faction reports faction of the card
func (c *CardLeader) Faction() CardFaction {
	return c.LeaderFaction
}

// Type reports that this is a leader card
func (c *CardLeader) Type() CardType {
	return TypeLeader
}

// Targettable reports if leader fx is targettable or not
func (c *CardLeader) Targettable() bool {
	//Todo: Treat by effect type
	return false
}
