package gwent

import "testing"

func TestWeatherCards(t *testing.T) {
	//Generate weather deck
	deck := Cards{}
	for i := 0; i < 10; i++ {
		deck = append(deck, &CardWeather{Target:RangeClose + CardRange(i%3)})
	}
	
	//Prepare players & cards
	p1, p2 := MakePlayer("test 1", FactionNorthernRealms, deck), MakePlayer("test 2", FactionMonsters, GenerateDeckWithUnitCards(RangeClose,5,30))
	
	//Create & reset game
	g := MakeGame(p1, p2)
	
	//Play all cards (this should apply all weather effects)
	for i := 0; i < 10; i++ {
		p1.Play(p1.Hand[0],nil);
	}
	
	//Check weather application
	if !g.WeatherClose || !g.WeatherRanged || !g.WeatherSiege {
		t.Error("All weather effects must be active")
	}
	
	//Give clear weather card & play it
	p1.GiveCard(&CardWeather{Target:RangeNone})
	p1.Play(p1.Hand[0],nil);
	
	//Check that all weather are clear
	if g.WeatherClose || g.WeatherRanged || g.WeatherSiege {
		t.Error("All weather effects must be in-active after casting Clear Weather")
	}
	
	//Give weather card, play it & pass both players
	p1.GiveCard(&CardWeather{Target:RangeClose})
	p1.Play(p1.Hand[0],nil); //Apply weather
	p1.Pass()
	p2.Pass()
	g.Next()
	
	//Check no weather
	if g.WeatherClose || g.WeatherRanged || g.WeatherSiege {
		t.Error("All weather effects must be in-active after new round")
	}
}