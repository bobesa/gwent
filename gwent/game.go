package gwent

const (
	FACTION_NEUTRAL = iota
	FACTION_NILFGAARD //Wins any round ending in draw
	FACTION_NORTHERN_REALMS //Grants an extra card upon winning round
	FACTION_MONSTERS //Keeps random Unit Card out after each round
	FACTION_SCOIATAEL //Decides who takes first turn
)

func MakeGame(p1 *Player, p2 *Player) (*Game) {		
	//Setup game
	g := &Game{
		Closed: false,
		Player1: p1,
		Player2: p2,
		eventIdChannel: make(chan int64, 10),
	}
	
	//Generate event ID channel
	go func(){
		eId := int64(0)
		for !g.Closed {
			eId++
			g.eventIdChannel <- eId
		}
	}()
	
	//Reset game
	g.Reset()
	
	return g
}

type Game struct {
	//States
	Closed bool
	Turn int
	WeatherClose, WeatherRanged, WeatherSiege bool
	
	//Players
	LastRoundWinner *Player
	Player1 *Player
	Player2 *Player
	
	//Events
	History []Event
	eventIdChannel chan int64
}

func (g *Game) GetNextEventId() int64 {
	return <- g.eventIdChannel
}

func (g *Game) MakeEvent(card, target Card, eventType string, p *Player) {
	g.PostEvent(MakeEvent(g, card, target, eventType, p))
}

func (g *Game) PostEvent(e Event) {
	//g.Player1.NewEvents <- e
	//g.Player2.NewEvents <- e
	g.History = append(g.History, e)
}

func (g *Game) GetCurrentPlayer() (*Player) {
	if g.Turn % 2 == 0 {
		return g.Player1
	} else {
		return g.Player2
	}
}

func (g *Game) Reset() {
	g.Player1.Game = g
	g.Player2.Game = g
	g.Player1.Reset()
	g.Player2.Reset()
	g.ClearWeather()
}

func (g *Game) ClearWeather() {
	g.WeatherClose, g.WeatherRanged, g.WeatherSiege = false, false, false	
}

func (g *Game) NextRound() {
	//Check winner
	if g.Player1.Lifes == 0 || g.Player2.Lifes == 0 {
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
			if g.Player1.Faction == FACTION_NILFGAARD && g.Player2.Faction != FACTION_NILFGAARD {
				//Player 1 is NILFGAARD
				g.Player2.Lost()				
			} else if g.Player1.Faction != FACTION_NILFGAARD && g.Player2.Faction == FACTION_NILFGAARD {
				//Player 2 is NILFGAARD
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