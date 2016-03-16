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
		case LEADER_FX_CANCEL:
			p.OtherPlayer().Leader.Cancel()
		case LEADER_FX_DRAW_EXTRA_CARD:
			p.DrawCard()
		case LEADER_FX_DRAW_GRAVE:
		case LEADER_FX_DRAW_OPPONENT_GRAVE:
		case LEADER_FX_DISCARD_2_DRAW_1:
		case LEADER_FX_LOOK_3_OPPONENT_CARDS:
			//TODO: Actually solve this :)

		//Destroy
		case LEADER_FX_DESTROY_CLOSE_10_PLUS:
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

		case LEADER_FX_DESTROY_SIEGE_10_PLUS:
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
		case LEADER_FX_PLAY_WEATHER_CLOSE:
			if !c.PlayWeatherFromDeck(RangeClose, p) {
				c.Ready()
			}
		case LEADER_FX_PLAY_WEATHER_RANGED:
			if !c.PlayWeatherFromDeck(RangeRanged, p) {
				c.Ready()
			}
		case LEADER_FX_PLAY_WEATHER_SIEGE:
			if !c.PlayWeatherFromDeck(RangeSiege, p) {
				c.Ready()
			}
		case LEADER_FX_WEATHER_CLEAR:
			p.Game.ClearWeather()

		//Horn
		case LEADER_FX_HORN_CLOSE:
			p.HornClose = true
		case LEADER_FX_HORN_RANGED:
			p.HornRanged = true
		case LEADER_FX_HORN_SIEGE:
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
		case LEADER_FX_PLAY_WEATHER_ANY:
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
	//TODO: Treat by effect type
	return false
}
