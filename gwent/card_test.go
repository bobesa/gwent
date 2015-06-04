package gwent

import "testing"

func GenerateUnitCard(cardRange, cardPower int) (*CardUnit) {
	return &CardUnit{
		Name: "Unit Card",
		Type: TYPE_BASIC,
		Range: cardRange,
		Power: cardPower,
		Hero: false,
	}
}

func GenerateDeckWithUnitCards(cardRange, cardPower, count int) (Cards) {
	cards := Cards{}
	for i := 0; i < count; i++ {
		cards = append(cards, GenerateUnitCard(cardRange, cardPower))
	}
	return cards
}

func TestCombatSimple(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FACTION_NILFGAARD, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30)), MakePlayer("test 2", FACTION_NILFGAARD, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30))
	
	//Create & reset game
	g := MakeGame(p1, p2)
	
	//Play game
	p1.Play(p1.Hand[0],nil); g.Next()
	p2.Play(p2.Hand[0],nil); g.Next()
	p1.Play(p1.Hand[0],nil); g.Next()
	p2.Pass(); g.Next()
	p1.Pass()
	
	if p1.ComputePower() != 10 || p2.ComputePower() != 5 {
		t.Error("Power is different than it should")
	}
	
	g.Next()
	
	if p1.ComputePower() != 0 || p2.ComputePower() != 0 {
		t.Error("There should be zero power after end of the round")
	} else if p2.Lifes != 1 {
		t.Error("Player should lose a life after defeat")
	}
}

func TestCombatDraw(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FACTION_NORTHERN_REALMS, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30)), MakePlayer("test 2", FACTION_NORTHERN_REALMS, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30))
	
	//Create & reset game
	g := MakeGame(p1, p2)
	
	//Play game
	p2.Pass()
	p1.Pass()
	g.Next()
	
	if p1.Lifes > 1 || p2.Lifes > 1 {
		t.Error("Both players should lose life upon draw")
	}
}

func TestFactionBonusNilfgaard(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FACTION_NILFGAARD, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30)), MakePlayer("test 2", FACTION_NORTHERN_REALMS, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30))
	
	//Create & reset game
	g := MakeGame(p1, p2)
	
	//Play game
	p1.Pass()
	p2.Pass()
	g.Next()
	
	if p2.Lifes != 1 {
		t.Error("Player should lose a life after a draw")
	}
}

func TestFactionBonusMonsters(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FACTION_MONSTERS, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30)), MakePlayer("test 2", FACTION_NORTHERN_REALMS, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30))
	
	//Create & reset game
	g := MakeGame(p1, p2)
	
	//Play 10 cards, then pass
	for c := 0; c < 10; c++ {
		p1.Play(p1.Hand[0],nil);
	}
	p1.Pass()
	p2.Pass()
	g.Next()
	
	if p1.ComputePower() == 0 && len(p1.RowClose) != 1 {
		t.Error("Player should keep one card")
	}
}

func TestFactionBonusNorthernRealms(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FACTION_NORTHERN_REALMS, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30)), MakePlayer("test 2", FACTION_MONSTERS, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30))
	
	//Create & reset game
	g := MakeGame(p1, p2)
	
	//Play 1 card, then pass
	p1.Play(p1.Hand[0],nil);
	p1.Pass()
	p2.Pass()
	g.Next()
	
	if len(p1.Hand) != 10 {
		t.Error("Player should have 10 cards in hand")
	}
}