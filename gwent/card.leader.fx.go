package gwent

const (
	LEADER_FX_NONE = iota

	//Northern Realms
	LEADER_FX_PLAY_WEATHER_RANGED   //Instantly play an Impenetrable Fog card from your deck.
	LEADER_FX_WEATHER_CLEAR         //Clear any weather effects (on both sides) currently in play.
	LEADER_FX_HORN_SIEGE            //Doubles the strength of all Siege units, unless a Commander's Horn is already in play on that row
	LEADER_FX_DESTROY_SIEGE_10_PLUS //Destroy opponent's strongest Siege units if the strength of that row is 10 or higher.

	//Nilfgaardian Empire
	LEADER_FX_PLAY_WEATHER_SIEGE    //Pick a Torrential Rain card from your deck and play it instantly.
	LEADER_FX_LOOK_3_OPPONENT_CARDS //Look at 3 random cards from your opponent's hand.
	LEADER_FX_CANCEL                //Cancel your opponent's Leader Ability.
	LEADER_FX_DRAW_OPPONENT_GRAVE   //Draw a card from your opponent's discard pile.

	//Scoia'tael
	LEADER_FX_PLAY_WEATHER_CLOSE    //Pick a Biting Frost card from your deck and play it instantly.
	LEADER_FX_DRAW_EXTRA_CARD       //Draw an extra card at the beginning of the battle.
	LEADER_FX_HORN_RANGED           //Doubles the strength of all your Ranged Combat units (unless a Commander's Horn is also present on that row).
	LEADER_FX_DESTROY_CLOSE_10_PLUS //Destroy your enemy's strongest Close Combat unit(s) if the combined strength of all his or her Close Combat units is 10 or more.

	//Monsters
	LEADER_FX_PLAY_WEATHER_ANY //Pick any weather card from your deck and play it instantly.
	LEADER_FX_HORN_CLOSE       //Double the strength of all your Close Combat units (unless a Commander's Horn is also present on that row).
	LEADER_FX_DISCARD_2_DRAW_1 //Discard 2 cards and draw 1 card of your choice from your deck.
	LEADER_FX_DRAW_GRAVE       //Restore a card from your discard pile to your hand.
)
