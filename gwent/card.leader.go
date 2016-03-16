package gwent

const (
	LEADER_FX_NONE = iota

	//Northern Realms
	LEADER_FX_PLAY_WEATHER_RANGED   //Instantly play an Impenetrable Fog card from your deck.
	LEADER_FX_WEATHER_CLEAR         //Clear any weather effects (on both sides) currently in play.
	LEADER_FX_HORN_SIEGE            //Doubles the strength of all Siege units, unless a Commander's Horn is already in play on that row
	LEADER_FX_DESTROY_SIEGE_10_PLUS //Destroy opponent's strongest Siege units if the strength of that row is 10 or higher.

	//Nilfgaardian Empire
	LEADER_FX_PLAY_WEATHER_SIEGE    //Pick a Torrential Rain card from your deck and play it instantly.
	LEADER_FX_LOOK_3_OPPONENT_CARDS //Look at 3 random cards from your opponent's hand.
	LEADER_FX_CANCEL                //Cancel your opponent's Leader Ability.
	LEADER_FX_DRAW_OPPONENT_GRAVE   //Draw a card from your opponent's discard pile.

	//Scoia'tael
	LEADER_FX_PLAY_WEATHER_CLOSE    //Pick a Biting Frost card from your deck and play it instantly.
	LEADER_FX_DRAW_EXTRA_CARD       //Draw an extra card at the beginning of the battle.
	LEADER_FX_HORN_RANGED           //Doubles the strength of all your Ranged Combat units (unless a Commander's Horn is also present on that row).
	LEADER_FX_DESTROY_CLOSE_10_PLUS //Destroy your enemy's strongest Close Combat unit(s) if the combined strength of all his or her Close Combat units is 10 or more.

	//Monsters
	LEADER_FX_PLAY_WEATHER_ANY //Pick any weather card from your deck and play it instantly.
	LEADER_FX_HORN_CLOSE       //Double the strength of all your Close Combat units (unless a Commander's Horn is also present on that row).
	LEADER_FX_DISCARD_2_DRAW_1 //Discard 2 cards and draw 1 card of your choice from your deck.
	LEADER_FX_DRAW_GRAVE       //Restore a card from your discard pile to your hand.
)

type CardLeader struct {
	CannotUse bool
	LeaderEffect    int
	LeaderFaction   CardFaction

	BasicCard
}

func (c *CardLeader) PlayWeatherFromDeck(where CardRange, p *Player) bool {
	var card Card
	p.Deck, card = p.Deck.WithoutType(TypeWeather, where)
	if card != nil {
		card.Play(p, nil) //Cast card
		return true
	} else {
		return false
	}
}

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

func (c *CardLeader) Cancel() {
	c.CannotUse = true
}

func (c *CardLeader) Ready() {
	c.CannotUse = false
}

func (c *CardLeader) Faction() CardFaction {
	return c.LeaderFaction
}

func (c *CardLeader) Type() CardType {
	return TypeLeader
}

func (c *CardLeader) Targettable() bool {
	//TODO: Treat by effect type
	return false
}
