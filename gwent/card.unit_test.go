package gwent

import "testing"

func TestUnitCards(t *testing.T) {	
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FACTION_NORTHERN_REALMS, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30)), MakePlayer("test 2", FACTION_MONSTERS, GenerateDeckWithUnitCards(RANGE_CLOSE,5,30))
	
	//Create & reset game
	g := MakeGame(p1, p2)
	
	//Play unit card
	card := &CardUnit{ Hero:false, Power: 5, Range: RANGE_CLOSE }
	p1.GiveCard(card)
	p1.Play(card, nil)
	
	//Apply weather (power = 1)
	p1.Game.WeatherClose = true
	if card.GetPower(p1) != 1 {
		t.Error("Power of card should drop to 1 when weather is applied")
	}
	p1.Game.WeatherClose = false
	if card.GetPower(p1) != 5 {
		t.Error("Power of card should be back to 5 when weather is cleared")
	}
	
	//Apply horn (double power)
	p1.HornClose = true
	if card.GetPower(p1) != 10 {
		t.Error("Power of card should raise to 10 when horn is applied")
	}
	p1.Game.WeatherClose = true
	if card.GetPower(p1) != 2 {
		t.Error("Power of card should drop to 2 when horn & weather is applied")
	}
	
	//Card is hero (ignore weather)
	card.Hero = true
	if card.GetPower(p1) != 10 {
		t.Error("Power of card should be 10 when horn & weather is applied on hero card")
	}	
	
	//End game
	g.Next();
}