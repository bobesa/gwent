package gwent

import "testing"

func TestHornDecoy(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeClose, 5, 10)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 10))

	//Create & reset game
	g := MakeGame(p1, p2)

	//Play 2 unit cards
	unit := p1.Hand[0]
	p1.Play(unit, nil)
	p1.Play(p1.Hand[0], nil)

	decoy := &CardDecoy{}
	p1.GiveCard(decoy)
	p1.Play(decoy, unit)

	//Check hand
	if p1.Hand.Has(decoy) {
		t.Error("Decoy should be on table's close row")
	}
	if !p1.Hand.Has(unit) {
		t.Error("Unit card should be in the hand")
	}

	//Check field
	if !p1.RowClose.Has(decoy) {
		t.Error("Decoy should NOT be in the hand")
	}
	if p1.RowClose.Has(unit) {
		t.Error("Unit card should NOT  be on table's close row")
	}

	//End game
	g.Next()
}
