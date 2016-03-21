package gwent

const (
	LeaderFxNone = iota

	//Northern Realms
	LeaderFxPlayWeatherRanged  //Instantly play an Impenetrable Fog card from your deck.
	LeaderFxWeatherClear       //Clear any weather effects (on both sides) currently in play.
	LeaderFxHornSiege          //Doubles the strength of all Siege units, unless a Commander's Horn is already in play on that row
	LeaderFxDestroySiege10Plus //Destroy opponent's strongest Siege units if the strength of that row is 10 or higher.

	//Nilfgaardian Empire
	LeaderFxPlayWeatherSiege   //Pick a Torrential Rain card from your deck and play it instantly.
	LeaderFxLook3OpponentCards //Look at 3 random cards from your opponent's hand.
	LeaderFxCancel             //Cancel your opponent's Leader Ability.
	LeaderFxDrawOpponentGrave  //Draw a card from your opponent's discard pile.

	//Scoia'tael
	LeaderFxPlayWeatherClose   //Pick a Biting Frost card from your deck and play it instantly.
	LeaderFxDrawExtraCard      //Draw an extra card at the beginning of the battle.
	LeaderFxHornRanged         //Doubles the strength of all your Ranged Combat units (unless a Commander's Horn is also present on that row).
	LeaderFxDestroyClose10Plus //Destroy your enemy's strongest Close Combat unit(s) if the combined strength of all his or her Close Combat units is 10 or more.

	//Monsters
	LeaderFxPlayWeatherAny //Pick any weather card from your deck and play it instantly.
	LeaderFxHornClose      //Double the strength of all your Close Combat units (unless a Commander's Horn is also present on that row).
	LeaderFxDiscard2Draw1  //Discard 2 cards and draw 1 card of your choice from your deck.
	LeaderFxDrawGrave      //Restore a card from your discard pile to your hand.
)
