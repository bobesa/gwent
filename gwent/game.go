package gwent

// MakeGame creates new game instance
func MakeGame(p1 *Player, p2 *Player) *Game {
	//Setup game
	g := &Game{
		Closed:         false,
		Player1:        p1,
		Player2:        p2,
		eventIDChannel: make(chan int64, 10),
	}

	//Generate event ID channel
	go func() {
		eID := int64(0)
		for !g.Closed {
			eID++
			g.eventIDChannel <- eID
		}
	}()

	//Reset game
	g.Reset()

	return g
}

// Game describes contents of single game
type Game struct {
	//States
	Closed                                    bool
	Turn                                      int
	WeatherClose, WeatherRanged, WeatherSiege bool

	//Players
	LastRoundWinner *Player
	Player1         *Player
	Player2         *Player

	//Events
	History        []Event
	eventIDChannel chan int64
}

// GetNextEventID generates new event ID
func (g *Game) GetNextEventID() int64 {
	return <-g.eventIDChannel
}

// MakeEvent creates new event and posts it to list of events
func (g *Game) MakeEvent(card, target Card, eventType string, p *Player) {
	g.PostEvent(MakeEvent(g, card, target, eventType, p))
}

// PostEvent posts event to history and players
func (g *Game) PostEvent(e Event) {
	//TODO: Actually post it to player events
	//g.Player1.NewEvents <- e
	//g.Player2.NewEvents <- e
	g.History = append(g.History, e)
}

// GetCurrentPlayer reports player that is on move
func (g *Game) GetCurrentPlayer() *Player {
	if g.Turn%2 == 0 {
		return g.Player1
	}
	return g.Player2
}

// Reset resets the game & effects
func (g *Game) Reset() {
	g.Player1.Game = g
	g.Player2.Game = g
	g.Player1.Reset()
	g.Player2.Reset()
	g.ClearWeather()
}

// ClearWeather clears the weather effects
func (g *Game) ClearWeather() {
	g.WeatherClose, g.WeatherRanged, g.WeatherSiege = false, false, false
}

// NextRound clears table, effects and moves to next round
func (g *Game) NextRound() {
	//Check winner
	if g.Player1.Lives == 0 || g.Player2.Lives == 0 {
		//WINNER
	}

	//Reset Weather
	g.ClearWeather()

	//Reset Passed
	g.Player1.Passed = false
	g.Player2.Passed = false

	//Reset Rows
	g.Player1.ResetRows()
	g.Player2.ResetRows()
}

// Next moves to next turn
// If both players are done, powers are computed and winner is chosen
func (g *Game) Next() {
	g.Turn++

	if g.Player1.Passed && g.Player2.Passed {
		//Compute winner
		pwr1, pwr2 := g.Player1.ComputePower(), g.Player2.ComputePower()

		if pwr1 > pwr2 {
			//Player1 wins
			g.Player2.Lost()
		} else if pwr1 < pwr2 {
			//Player2 wins
			g.Player1.Lost()
		} else {
			//Tie
			if g.Player1.Faction == FactionNilfgaard && g.Player2.Faction != FactionNilfgaard {
				//Player 1 is Nilfgaard
				g.Player2.Lost()
			} else if g.Player1.Faction != FactionNilfgaard && g.Player2.Faction == FactionNilfgaard {
				//Player 2 is Nilfgaard
				g.Player1.Lost()
			} else {
				//Both players lose life
				g.Player1.Lost()
				g.Player2.Lost()
				g.LastRoundWinner = nil
			}
		}

		g.NextRound()
	} else if g.GetCurrentPlayer().Passed {
		g.Next()
	}
}
