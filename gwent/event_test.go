package gwent

import "testing"

// TestEventToHand tests "to hand" event
func TestEventToHand(t *testing.T) {
	// Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeClose, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 30))

	// Create & reset game
	g := MakeGame(p1, p2)

	// Create card & toHand event
	_, handedCard := p1.Deck.WithoutRandom() //Pick random card from Deck
	e := MakeEvent(g, handedCard, nil, EventToHand, p1)
	t.Error(e.CardID, handedCard.GUID())

	// Read current amount of cards in hand
	handLen := len(p1.Hand)

	// Post event to game
	g.PostEvent(e)

	// Check that we have more cards after registering the event
	if currentHandLen := len(p1.Hand); handLen+1 != currentHandLen {
		t.Errorf("Amount of cards in hand should be %d, but got %d instead", handLen+1, currentHandLen)
	}

	// End game
	g.Next()
}
