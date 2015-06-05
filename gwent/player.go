package gwent

import (
	"math/rand"
)

func MakePlayer(name string, faction int, cards Cards) (*Player) {
	return &Player{
		Name: name,
		Faction: faction,
		Cards: cards,
		Passed: false,
		
		Deck: make(Cards,0),
		Hand: make(Cards,0),		
		Grave: make(Cards,0),
		RowClose: make(Cards,0),
		RowRanged: make(Cards,0),
		RowSiege: make(Cards,0),
		
		NewEvents: make(chan Event),
	}	
}

type Player struct {
	Name string
	Lifes int
	Faction int
	Leader *CardLeader
	Cards Cards
	Game *Game
	Passed bool
	HornClose, HornRanged, HornSiege bool
	
	Deck Cards
	Hand Cards
	Grave Cards
	RowClose Cards
	RowRanged Cards
	RowSiege Cards
	
	NewEvents chan Event
}

func (p *Player) IsActive() bool {
	return p == p.Game.GetCurrentPlayer()
}

func (p *Player) Play(card Card, target Card) {
	card.Play(p, target)
	p.Hand = p.Hand.Without(card)
}

func (p *Player) PlayOnRow(card Card, target int) {
	card.PlayOnRow(p, target)
	p.Hand = p.Hand.Without(card)
}

func (p *Player) PlayLeader(target Card) {
	if p.Leader != nil && !p.Leader.CannotUse {
		p.Leader.Play(p, target)
	}
}

func (p *Player) PlayLeaderOnRow(target int) {
	if p.Leader != nil && !p.Leader.CannotUse {
		p.Leader.PlayOnRow(p, target)
	}
}

func (p *Player) Horn(where int) {
	switch(where){
	case RANGE_CLOSE:
		p.HornClose = true
	case RANGE_RANGED:
		p.HornRanged = true
	case RANGE_SIEGE:
		p.HornSiege = true
	}
}

func (p *Player) Scorch() {
	//Go trough whole table and select highest power card
	maxPower := 0
	table := Cards{}
	table = append(table, p.RowClose ...)
	table = append(table, p.RowRanged ...)
	table = append(table, p.RowSiege ...)
	for _, c := range table {
		pwr := c.GetPower(p)
		if pwr > maxPower {
			maxPower = pwr
		}
	}
	table = Cards{}
	table = append(table, p.OtherPlayer().RowClose ...)
	table = append(table, p.OtherPlayer().RowRanged ...)
	table = append(table, p.OtherPlayer().RowSiege ...)
	for _, c := range table {
		pwr := c.GetPower(p.OtherPlayer())
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

func (p *Player) ComputePower() int {
	pwr := 0
	
	//Gather all cards on table
	table := make(Cards,0)
	table = append(table, p.RowClose...)
	table = append(table, p.RowRanged...)
	table = append(table, p.RowSiege...)
	
	//Compute power of each card
	for _, card := range table {
		pwr += card.GetPower(p)		
	}
	
	return pwr
}

func (p *Player) ComputePowerOfRow(row int) int {
	pwr := 0
	
	//Get all cards on particular table row
	var table Cards
	switch(row){
	case RANGE_CLOSE:
		table = p.RowClose
	case RANGE_RANGED:
		table = p.RowRanged
	case RANGE_SIEGE:
		table = p.RowSiege
	}
	
	//Compute power of each card
	for _, card := range table {
		pwr += card.GetPower(p)		
	}
	
	return pwr
}

func (p *Player) OtherPlayer() (*Player) {
	if p.Game.Player1 == p {
		return p.Game.Player2
	}
	return p.Game.Player1
}

func (p *Player) Pass() {
	p.Passed = true
}

func (p *Player) Lost() {
	p.Lifes--
	
	//Set LastRoundWinner of the game
	if p.Game.Player1 == p {
		p.Game.LastRoundWinner = p.Game.Player2		
	} else {
		p.Game.LastRoundWinner = p.Game.Player1
	}
}

func (p *Player) ResetRows() {
	//Faction related effects
	var monsterCard Card
	if p.Faction == FACTION_MONSTERS {
		if len(p.RowClose) > 0 {
			p.RowClose, monsterCard = p.RowClose.WithoutRandom()
		} else if len(p.RowRanged) > 0 {
			p.RowClose, monsterCard = p.RowRanged.WithoutRandom()			
		} else if len(p.RowSiege) > 0 {
			p.RowClose, monsterCard = p.RowSiege.WithoutRandom()			
		}
	}
	
	//Put all cards to Grave
	p.Grave = append(p.Grave, p.RowClose ...)
	p.Grave = append(p.Grave, p.RowRanged ...)
	p.Grave = append(p.Grave, p.RowSiege ...)
	
	//Reset rows
	p.RowClose = make(Cards,0)
	p.RowRanged = make(Cards,0)
	p.RowSiege = make(Cards,0)	
	
	//Faction related post-effects
	if p.Faction == FACTION_MONSTERS && monsterCard != nil {
		monsterCard.PutOnTable(p)
	} else if p.Faction == FACTION_NORTHERN_REALMS && p.Game.LastRoundWinner == p {
		p.DrawCard()
	}	
	
	//Reset Horns
	p.HornClose, p.HornRanged, p.HornSiege = false, false, false
}

func (p *Player) GiveCard(c Card) {
	if c != nil {
		p.Hand = append(p.Hand, c)
	}
}

func (p *Player) DrawCard() {	
	if len(p.Deck) > 0 {
		index := rand.Intn(len(p.Deck)) //random card index from rest of the deck
		card := p.Deck[index]
		p.Hand = append(p.Hand, card)
		p.Deck = append(p.Deck[:index], p.Deck[index+1:]...)
		p.Game.MakeEvent(card, nil, EVENT_TO_HAND, p)
	}
}

func (p *Player) Reset() {
	//Reset Lifes to 2
	p.Lifes = 2
	
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
	if p.Leader != nil && p.Leader.Effect == LEADER_FX_DRAW_EXTRA_CARD {
		p.Leader.Play(p,nil)
	}
}