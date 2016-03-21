package gwent

import "testing"

func TestLeaderCards_HornClose(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeSiege, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 30))
	p1.Leader = &CardLeader{LeaderEffect: LeaderFxHornSiege}

	//Create & reset game
	g := MakeGame(p1, p2)

	//Check "Horn on Siege units"
	p1.Leader.LeaderEffect = LeaderFxHornSiege
	p1.Play(p1.Hand[0], nil)
	if p1.ComputePower() != 5 {
		t.Error("Combat power should NOT be doubled before siege horn leader power is used")
	}
	p1.PlayLeader(nil)
	if p1.ComputePower() != 10 {
		t.Error("Combat power should be doubled when siege horn leader power is used")
	}

	//End game
	g.Next()
}

func TestLeaderCards_HornRanged(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeRanged, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 30))
	p1.Leader = &CardLeader{LeaderEffect: LeaderFxHornRanged}

	//Create & reset game
	g := MakeGame(p1, p2)

	//Check "Horn on Ranged units"
	p1.Leader.LeaderEffect = LeaderFxHornRanged
	p1.Play(p1.Hand[0], nil)
	if p1.ComputePower() != 5 {
		t.Error("Combat power should Not be doubled before ranged horn leader power is used")
	}
	p1.PlayLeader(nil)
	if p1.ComputePower() != 10 {
		t.Error("Combat power should be doubled when ranged horn leader power is used")
	}

	//End game
	g.Next()
}

func TestLeaderCards_HornSiege(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeClose, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 30))
	p1.Leader = &CardLeader{LeaderEffect: LeaderFxHornClose}

	//Create & reset game
	g := MakeGame(p1, p2)

	//Check "Horn on Close units"
	p1.Leader.LeaderEffect = LeaderFxHornClose
	p1.Play(p1.Hand[0], nil)
	if p1.ComputePower() != 5 {
		t.Error("Combat power should Not be doubled before close horn leader power is used")
	}
	p1.PlayLeader(nil)
	if p1.ComputePower() != 10 {
		t.Error("Combat power should be doubled when close horn leader power is used")
	}

	//End game
	g.Next()
}

func TestLeaderCards_WeatherClear(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeClose, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 30))
	p1.Leader = &CardLeader{LeaderEffect: LeaderFxHornClose}

	//Create & reset game
	g := MakeGame(p1, p2)

	//Check
	weatherCardClose, weatherCardRanged := &CardWeather{Target: RangeClose}, &CardWeather{Target: RangeRanged}
	p1.Leader.LeaderEffect = LeaderFxWeatherClear
	p1.GiveCard(weatherCardClose)
	p1.GiveCard(weatherCardRanged)
	p1.Play(weatherCardClose, nil)
	p1.Play(weatherCardRanged, nil)
	if !p1.Game.WeatherClose || !p1.Game.WeatherRanged {
		t.Error("Close & Ranged weathers should be applied")
	}
	if p1.Game.WeatherSiege {
		t.Error("Siege weather should Not be applied")
	}
	p1.PlayLeader(nil)
	if p1.Game.WeatherClose || p1.Game.WeatherRanged || p1.Game.WeatherSiege {
		t.Error("No weather should be applied")
	}

	//End game
	g.Next()
}

func TestLeaderCards_WeatherClose(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeClose, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 30))
	p1.Leader = &CardLeader{LeaderEffect: LeaderFxPlayWeatherClose}

	//Create & reset game
	g := MakeGame(p1, p2)

	//Check effect without proper card
	if p1.Game.WeatherClose || p1.Game.WeatherSiege || p1.Game.WeatherRanged {
		t.Error("Weather should Not be applied")
	}
	p1.PlayLeader(nil)
	if p1.Game.WeatherClose {
		t.Error("Close weather should Not be applied (we don't have proper weather card in deck)")
	}
	if p1.Leader.CannotUse {
		t.Error("Leader effect was marked as used without available weather card")
	}

	//Give weather card & try again
	p1.Deck = append(p1.Deck, &CardWeather{Target: RangeClose})
	deckCount := len(p1.Deck)
	p1.PlayLeader(nil)
	if !p1.Game.WeatherClose {
		t.Error("Close weather should be applied")
	}
	if deckCount == len(p1.Deck) {
		t.Error("Weather card was not consumed")
	}

	//End game
	g.Next()
}

func TestLeaderCards_WeatherRanged(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeRanged, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 30))
	p1.Leader = &CardLeader{LeaderEffect: LeaderFxPlayWeatherRanged}

	//Create & reset game
	g := MakeGame(p1, p2)

	//Check effect without proper card
	if p1.Game.WeatherClose || p1.Game.WeatherSiege || p1.Game.WeatherRanged {
		t.Error("Weather should Not be applied")
	}
	p1.PlayLeader(nil)
	if p1.Game.WeatherRanged {
		t.Error("Close weather should Not be applied (we don't have proper weather card in deck)")
	}
	if p1.Leader.CannotUse {
		t.Error("Leader effect was marked as used without available weather card")
	}

	//Give weather card & try again
	p1.Deck = append(p1.Deck, &CardWeather{Target: RangeRanged})
	deckCount := len(p1.Deck)
	p1.PlayLeader(nil)
	if !p1.Game.WeatherRanged {
		t.Error("Ranged weather should be applied")
	}
	if deckCount == len(p1.Deck) {
		t.Error("Weather card was not consumed")
	}

	//End game
	g.Next()
}

func TestLeaderCards_WeatherSiege(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeSiege, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 30))
	p1.Leader = &CardLeader{LeaderEffect: LeaderFxPlayWeatherSiege}

	//Create & reset game
	g := MakeGame(p1, p2)

	//Check effect without proper card
	if p1.Game.WeatherClose || p1.Game.WeatherSiege || p1.Game.WeatherRanged {
		t.Error("Weather should Not be applied")
	}
	p1.PlayLeader(nil)
	if p1.Game.WeatherSiege {
		t.Error("Close weather should Not be applied (we don't have proper weather card in deck)")
	}
	if p1.Leader.CannotUse {
		t.Error("Leader effect was marked as used without available weather card")
	}

	//Give weather card & try again
	p1.Deck = append(p1.Deck, &CardWeather{Target: RangeSiege})
	deckCount := len(p1.Deck)
	p1.PlayLeader(nil)
	if !p1.Game.WeatherSiege {
		t.Error("Siege weather should be applied")
	}
	if deckCount == len(p1.Deck) {
		t.Error("Weather card was not consumed")
	}

	//End game
	g.Next()
}

func TestLeaderCards_DrawExtraCard(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeSiege, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 30))
	p1.Leader = &CardLeader{LeaderEffect: LeaderFxDrawExtraCard}

	//Create & reset game
	g := MakeGame(p1, p2)

	//Check amount of cards in hand
	if len(p1.Hand) != 11 {
		t.Error("Player should have 1 more card in hand")
	}

	//End game
	g.Next()
}

func TestLeaderCards_DestroyClose(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeSiege, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose, 5, 30))
	p1.Leader = &CardLeader{LeaderEffect: LeaderFxDestroyClose10Plus}

	//Create & reset game
	g := MakeGame(p1, p2)

	//Give player 2 special unit card to hand
	specialCard, otherRowCard := GenerateUnitCard(RangeClose, 4), GenerateUnitCard(RangeRanged, 5)
	p2.GiveCard(specialCard)
	p2.GiveCard(otherRowCard)
	p2.Play(specialCard, nil)
	p2.Play(otherRowCard, nil)
	for i := 0; i < 4; i++ {
		p2.Play(p2.Hand[0], nil)
	}

	//Check amount of cards in hand
	if p2.ComputePower() <= 9 {
		t.Error("Player should more cards on the field")
	}
	p1.PlayLeader(nil)
	if p2.ComputePower() != 9 {
		t.Error("Player should have only special card & other rows cards on the field after destroy spell")
	}

	//End game
	g.Next()
}

func TestLeaderCards_DestroySiege(t *testing.T) {
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, GenerateDeckWithUnitCards(RangeSiege, 5, 30)), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeSiege, 5, 30))
	p1.Leader = &CardLeader{LeaderEffect: LeaderFxDestroySiege10Plus}

	//Create & reset game
	g := MakeGame(p1, p2)

	//Give player 2 special unit card to hand
	specialCard, otherRowCard := GenerateUnitCard(RangeSiege, 4), GenerateUnitCard(RangeClose, 5)
	p2.GiveCard(specialCard)
	p2.GiveCard(otherRowCard)
	p2.Play(specialCard, nil)
	p2.Play(otherRowCard, nil)
	for i := 0; i < 4; i++ {
		p2.Play(p2.Hand[0], nil)
	}

	//Check amount of cards in hand
	if p2.ComputePower() <= 9 {
		t.Error("Player should more cards on the field")
	}
	p1.PlayLeader(nil)
	if p2.ComputePower() != 9 {
		t.Error("Player should have only special card & other rows cards on the field after destroy spell")
	}

	//End game
	g.Next()
}
