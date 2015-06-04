package gwent

const (
	TYPE_DRAW = "draw"
)

type Event struct {	
	Id int64 `json:"id"`
	Message string `json:"msg"`
	Type string `json:"type"`
	Source GUID `json:"source"` //GUID of source card
	Target string `json:"target"` //grave, close, ranged, siege, hand, popup
	Player int	`json:"player"`	//1 = player 1 etc.
	
	game Game
}

func MakeEvent(g Game, source Card, msg, t, target string, p *Player) Event {
	e := Event{
		Id: g.GetNextEventId(),
		Message: msg,
		Type: t,
		Source: source.GetGUID(),
		Target: target,		
	}
	
	if g.Player1 == p {
		e.Player = 1
	} else {
		e.Player = 2
	}
	
	return e
}