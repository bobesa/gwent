package gwent

const (
	EVENT_TO_HAND  = "to_hand"  //"Card" should go to hand
	EVENT_TO_DECK  = "to_deck"  //"Card" should go to deck
	EVENT_TO_GRAVE = "to_grave" //"Card" should go to grave
	EVENT_TO_TABLE = "to_table" //"Card" should go to table

	EVENT_PICK   = "pick"   //From server: Sent with Cards for player to select
	EVENT_PICKED = "picked" //From player: Sent with selected Card
)

type Event struct {
	Id     int64  `json:"id"`
	Type   string `json:"type",omitempty`
	Cards  Cards  `json:"cards",omitempty`
	Card   GUID   `json:"card",omitempty`   //GUID of selected card
	Target GUID   `json:"target",omitempty` //GUID of target card
	Player int    `json:"player"`           //1 = player 1 etc.

	game *Game
}

func MakeEvent(g *Game, card, target Card, eventType string, p *Player) Event {
	e := Event{
		Id:     g.GetNextEventId(),
		Card:   card.GUID(),
		Target: card.GUID(),
		Type:   eventType,

		game: g,
	}

	if g.Player1 == p {
		e.Player = 1
	} else {
		e.Player = 2
	}

	return e
}
