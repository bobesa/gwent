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
	ID       EventID `json:"id"`
	Type     string  `json:"type,omitempty"`
	Cards    Cards   `json:"cards,omitempty"`
	CardID   GUID    `json:"card,omitempty"`   //GUID of selected card
	TargetID GUID    `json:"target,omitempty"` //GUID of target card
	PlayerID int     `json:"player"`           //1 = player 1 etc.

	game *Game
}

// MakeEvent creates new event for given game, card, target etc.
func MakeEvent(g *Game, card, target Card, eventType string, p *Player) Event {
	e := Event{
		ID:   g.GetNextEventID(),
		Type: eventType,

		game: g,
	}

	if card != nil {
		e.CardID = card.GUID()
	}
	if target != nil {
		e.TargetID = target.GUID()
	}

	if g.Player1 == p {
		e.PlayerID = 1
	} else {
		e.PlayerID = 2
	}

	return e
}

// Player reports event player instance
func (e Event) Player() *Player {
	if e.PlayerID == 1 {
		return e.game.Player1
	}
	return e.game.Player2
}
