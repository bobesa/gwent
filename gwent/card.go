package gwent

import (
	"math/rand"
)

type Cards []Card

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

func (deck Cards) Without(c Card) Cards {
	for i, card := range deck {
		if card == c {
			return append(deck[:i], deck[i+1:]...)
		}
	}
	return deck
}

func (deck Cards) WithoutType(cardType CardType, cardRange CardRange) (Cards, Card) {
	for i, card := range deck {
		if card.Type() == cardType && card.Range() == cardRange {
			return append(deck[:i], deck[i+1:]...), card
		}
	}
	return deck, nil
}

func (deck Cards) WithoutRandom() (Cards, Card) {
	if len(deck) > 0 {
		i := rand.Intn(len(deck))
		card := deck[i]
		return deck.Without(card), card
	}
	return deck, nil
}

func (deck Cards) Has(card Card) bool {
	for _, c := range deck {
		if c == card {
			return true
		}
	}
	return false
}

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
