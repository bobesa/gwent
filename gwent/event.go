package gwent

const (
	// EventToHand means that card should go to hand
	EventToHand = "to_hand"

	// EventToDeck means that card should go to deck
	EventToDeck = "to_deck"

	// EventToGrave means that card should go to grave
	EventToGrave = "to_grave"

	// EventToTable means that card should go to table
	EventToTable = "to_table"

	// EventPick gives player list of cards to select from.
	// This should be followed by EventPicked event coming from player
	EventPick = "pick"

	// EventPicked player picked card from selection
	EventPicked = "picked"
)

// EventID defines what should be used for type
type EventID uint64

// Event is json structure describing game events
type Event struct {
	ID     EventID `json:"id"`
	Type   string  `json:"type,omitempty"`
	Cards  Cards   `json:"cards,omitempty"`
	Card   GUID    `json:"card,omitempty"`   //GUID of selected card
	Target GUID    `json:"target,omitempty"` //GUID of target card
	Player int     `json:"player"`           //1 = player 1 etc.

	game *Game
}

// MakeEvent creates new event for given game, card, target etc.
func MakeEvent(g *Game, card, target Card, eventType string, p *Player) Event {
	e := Event{
		ID:     g.GetNextEventID(),
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
