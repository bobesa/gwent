package gwent

import "testing"

func TestScorchCards(t *testing.T) {
	//Generate scorch deck
	deck := Cards{}
	for i := 0; i < 10; i++ {
		deck = append(deck, &CardScorch{})
	}
	
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FACTION_NORTHERN_REALMS, deck), MakePlayer("test 2", FACTION_MONSTERS, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30))
	
	//Create & reset game
	g := MakeGame(p1, p2)
	
	//Play all unit cards of player 2
	for i := 0; i < 10; i++ {
		p2.Play(p2.Hand[0],nil);
	}
	
	//Power check
	if p2.ComputePower() != 50 {
		t.Error("Player 2 power should be 50")
	}
	
	//Play Scorch (which should kill all cards on table)
	p1.Play(p1.Hand[0],nil);
	
	//Power check
	if p2.ComputePower() != 0 {
		t.Error("Player 2 power should be 0")
	}
	
	g.Next()
}