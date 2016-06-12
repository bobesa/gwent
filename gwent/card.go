package gwent

import (
	"math/rand"
)

// Cards is a collection of Card objects
type Cards []Card

// Scorch kills certain cards of given player based on given power (pwr)
// Reports back actual cards and destroyed cards
func (deck Cards) Scorch(owner *Player, pwr int) (Cards, Cards) {
	cards, destroyed := Cards{}, Cards{}
	for _, card := range deck {
		if !card.Hero() && card.Power(owner) == pwr {
			destroyed = append(destroyed, card)
		} else {
			cards = append(cards, card)
		}
	}
	return cards, destroyed
}

// ByID reports back Card with given ID or nil
func (deck Cards) ByID(cardID GUID) Card {
	for _, card := range deck {
		if card.GUID() == cardID {
			return card
		}
	}
	return nil
}

// Without reports back Cards without a single given card
func (deck Cards) Without(c Card) Cards {
	for i, card := range deck {
		if card == c {
			return append(deck[:i], deck[i+1:]...)
		}
	}
	return deck
}

// WithoutType reports back Cards without cards of given type and range
func (deck Cards) WithoutType(cardType CardType, cardRange CardRange) (Cards, Card) {
	for i, card := range deck {
		if card.Type() == cardType && card.Range() == cardRange {
			return append(deck[:i], deck[i+1:]...), card
		}
	}
	return deck, nil
}

// WithoutRandom reports back Cards without one random card
func (deck Cards) WithoutRandom() (Cards, Card) {
	if len(deck) > 0 {
		i := rand.Intn(len(deck))
		card := deck[i]
		return deck.Without(card), card
	}
	return deck, nil
}

// Has reports if given card is in the collection
func (deck Cards) Has(card Card) bool {
	for _, c := range deck {
		if c == card {
			return true
		}
	}
	return false
}

// Card describes what methods must be on all card types
type Card interface {
	Play(*Player, Card)
	PlayOnRow(*Player, CardRange)
	PutOnTable(*Player)

	Name() string
	Faction() CardFaction
	Type() CardType
	Range() CardRange
	Power(*Player) int
	Hero() bool
	AppliedOnRow() bool
	Targettable() bool

	//Guids
	SetGUID(GUID)
	GUID() GUID
}
