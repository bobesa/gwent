package gwent

import (
	"math/rand"
)

// MakePlayer creates new instance of player
func MakePlayer(name string, faction CardFaction, cards Cards) *Player {
	return &Player{
		Name:    name,
		Faction: faction,
		Cards:   cards,
		Passed:  false,

		Deck:      make(Cards, 0),
		Hand:      make(Cards, 0),
		Grave:     make(Cards, 0),
		RowClose:  make(Cards, 0),
		RowRanged: make(Cards, 0),
		RowSiege:  make(Cards, 0),

		NewEvents: make(chan Event),
	}
}

// Player describes what cards, stats etc. player has
type Player struct {
	Name                             string
	Lives                            int
	Faction                          CardFaction
	Leader                           *CardLeader
	Cards                            Cards
	Game                             *Game
	Passed                           bool
	HornClose, HornRanged, HornSiege bool

	Deck      Cards
	Hand      Cards
	Grave     Cards
	RowClose  Cards
	RowRanged Cards
	RowSiege  Cards

	NewEvents chan Event
}

// IsActive reports if this player is currently on move
func (p *Player) IsActive() bool {
	return p == p.Game.GetCurrentPlayer()
}

// Play plays given card on given target
func (p *Player) Play(card Card, target Card) {
	card.Play(p, target)
	p.Hand = p.Hand.Without(card)
}

// PlayOnRow plays given card on given row
func (p *Player) PlayOnRow(card Card, row CardRange) {
	card.PlayOnRow(p, row)
	p.Hand = p.Hand.Without(card)
}

// PlayLeader plays leader card effect on target
func (p *Player) PlayLeader(target Card) {
	if p.Leader != nil && !p.Leader.CannotUse {
		p.Leader.Play(p, target)
	}
}

// PlayLeaderOnRow plays leader card effect on row
func (p *Player) PlayLeaderOnRow(target CardRange) {
	if p.Leader != nil && !p.Leader.CannotUse {
		p.Leader.PlayOnRow(p, target)
	}
}

// Horn applies "Horn" effect on given row
func (p *Player) Horn(where CardRange) {
	switch where {
	case RangeClose:
		p.HornClose = true
	case RangeRanged:
		p.HornRanged = true
	case RangeSiege:
		p.HornSiege = true
	}
}

// Scorch applies "Scorch" effect on opponent's table
func (p *Player) Scorch() {
	//Go trough whole table and select highest power card
	maxPower := 0
	table := Cards{}
	table = append(table, p.RowClose...)
	table = append(table, p.RowRanged...)
	table = append(table, p.RowSiege...)
	for _, c := range table {
		pwr := c.Power(p)
		if pwr > maxPower {
			maxPower = pwr
		}
	}
	table = Cards{}
	table = append(table, p.OtherPlayer().RowClose...)
	table = append(table, p.OtherPlayer().RowRanged...)
	table = append(table, p.OtherPlayer().RowSiege...)
	for _, c := range table {
		pwr := c.Power(p.OtherPlayer())
		if pwr > maxPower {
			maxPower = pwr
		}
	}

	//Kill all non-hero cards with this power
	k1, k2, k3 := Cards{}, Cards{}, Cards{}

	p.RowClose, k1 = p.RowClose.Scorch(p, maxPower)
	p.RowRanged, k2 = p.RowRanged.Scorch(p, maxPower)
	p.RowSiege, k3 = p.RowSiege.Scorch(p, maxPower)
	p.Grave = append(p.Grave, k1...)
	p.Grave = append(p.Grave, k2...)
	p.Grave = append(p.Grave, k3...)

	p.OtherPlayer().RowClose, k1 = p.OtherPlayer().RowClose.Scorch(p.OtherPlayer(), maxPower)
	p.OtherPlayer().RowRanged, k2 = p.OtherPlayer().RowRanged.Scorch(p.OtherPlayer(), maxPower)
	p.OtherPlayer().RowSiege, k3 = p.OtherPlayer().RowSiege.Scorch(p.OtherPlayer(), maxPower)
	p.OtherPlayer().Grave = append(p.OtherPlayer().Grave, k1...)
	p.OtherPlayer().Grave = append(p.OtherPlayer().Grave, k2...)
	p.OtherPlayer().Grave = append(p.OtherPlayer().Grave, k3...)
}

// ComputePower reports total power of cards
func (p *Player) ComputePower() int {
	pwr := 0

	//Gather all cards on table
	table := make(Cards, 0)
	table = append(table, p.RowClose...)
	table = append(table, p.RowRanged...)
	table = append(table, p.RowSiege...)

	//Compute power of each card
	for _, card := range table {
		pwr += card.Power(p)
	}

	return pwr
}

// ComputePowerOfRow reports total power of row
func (p *Player) ComputePowerOfRow(row CardRange) int {
	pwr := 0

	//Get all cards on particular table row
	var table Cards
	switch row {
	case RangeClose:
		table = p.RowClose
	case RangeRanged:
		table = p.RowRanged
	case RangeSiege:
		table = p.RowSiege
	}

	//Compute power of each card
	for _, card := range table {
		pwr += card.Power(p)
	}

	return pwr
}

// OtherPlayer reports opponent player
func (p *Player) OtherPlayer() *Player {
	if p.Game.Player1 == p {
		return p.Game.Player2
	}
	return p.Game.Player1
}

// Pass passes this round
func (p *Player) Pass() {
	p.Passed = true
}

// Lost removes single life of player
// Ends the game if player has none
func (p *Player) Lost() {
	p.Lives--

	//Set LastRoundWinner of the game
	if p.Game.Player1 == p {
		p.Game.LastRoundWinner = p.Game.Player2
	} else {
		p.Game.LastRoundWinner = p.Game.Player1
	}
}

// ResetRows empties the table (and applies Monster faction effect if possible)
func (p *Player) ResetRows() {
	//Faction related effects
	var monsterCard Card
	if p.Faction == FactionMonsters {
		if len(p.RowClose) > 0 {
			p.RowClose, monsterCard = p.RowClose.WithoutRandom()
		} else if len(p.RowRanged) > 0 {
			p.RowClose, monsterCard = p.RowRanged.WithoutRandom()
		} else if len(p.RowSiege) > 0 {
			p.RowClose, monsterCard = p.RowSiege.WithoutRandom()
		}
	}

	//Put all cards to Grave
	p.Grave = append(p.Grave, p.RowClose...)
	p.Grave = append(p.Grave, p.RowRanged...)
	p.Grave = append(p.Grave, p.RowSiege...)

	//Reset rows
	p.RowClose, p.RowRanged, p.RowSiege = make(Cards, 0), make(Cards, 0), make(Cards, 0)

	//Faction related post-effects
	if p.Faction == FactionMonsters && monsterCard != nil {
		monsterCard.PutOnTable(p)
	} else if p.Faction == FactionNorthernRealms && p.Game.LastRoundWinner == p {
		p.DrawCard()
	}

	//Reset Horns
	p.HornClose, p.HornRanged, p.HornSiege = false, false, false
}

// GiveCard gives card to players hand
func (p *Player) GiveCard(c Card) {
	if c != nil {
		p.Hand = append(p.Hand, c)
	}
}

// DrawCard give card to players hand from players deck
func (p *Player) DrawCard() {
	if len(p.Deck) > 0 {
		index := rand.Intn(len(p.Deck)) //random card index from rest of the deck
		card := p.Deck[index]
		p.Hand = append(p.Hand, card)
		p.Deck = append(p.Deck[:index], p.Deck[index+1:]...)
		p.Game.MakeEvent(card, nil, EventToHand, p)
	}
}

// Reset resets hand, sets player lives to 2
func (p *Player) Reset() {
	//Reset Lifes to 2
	p.Lives = 2

	//Reset CurrentDeck to "Original" Deck
	p.Deck = p.Cards

	//Go trough Deck & generate GUIDs
	for _, c := range p.Deck {
		c.SetGUID(GetNextGUID())
	}
	if p.Leader != nil {
		p.Leader.SetGUID(GetNextGUID())
	}

	//Give 10 random cards from CurrentDeck to Hand
	for i := 0; i < 10; i++ {
		p.DrawCard()
	}

	//Check for Leader-related effects
	if p.Leader != nil && p.Leader.LeaderEffect == LeaderFxDrawExtraCard {
		p.Leader.Play(p, nil)
	}
}
