package gwent

import "testing"

func TestHornCards(t *testing.T) {	
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeClose,5,10)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose,5,10))
	
	//Create & reset game
	g := MakeGame(p1, p2)
	
	//Play all unit cards
	for i := 0; i < 10; i++ {
		p1.Play(p1.Hand[0],nil);
	}
	
	//Give player 3 horn cards
	p1.GiveCard(&CardHorn{})
	p1.GiveCard(&CardHorn{})
	p1.GiveCard(&CardHorn{})
	
	//Cast horn cards & test horn application
	p1.PlayOnRow(p1.Hand[0], RangeClose)
	if !p1.HornClose {
		t.Error("Horn on close combat units should be active")
	}
	p1.PlayOnRow(p1.Hand[0], RangeRanged)
	if !p1.HornRanged {
		t.Error("Horn on ranged combat units should be active")
	}
	p1.PlayOnRow(p1.Hand[0], RangeSiege)
	if !p1.HornSiege {
		t.Error("Horn on siege combat units should be active")
	}
	
	//Check horn application
	if !p1.HornClose || !p1.HornRanged || !p1.HornSiege {
		t.Error("All horn effects must be active")
	}
	
	//Check that player 2 has any horn effects active
	if p2.HornClose || p2.HornRanged || p2.HornSiege {
		t.Error("Horn effect should not be happening on player 2")
	}
	
	//Next round
	p1.Pass()
	p2.Pass()
	g.Next()
	
	//Check no horn
	if p1.HornClose || p1.HornRanged || p1.HornSiege || p2.HornClose || p2.HornRanged || p2.HornSiege {
		t.Error("Horn effects should be cleared after next round")
	}
}