package gwent

// LeaderFx defines type of leader effect
type LeaderFx int

// List of leader effects (or abilities)
const (
	LeaderFxNone = LeaderFx(iota)

	// Northern Realms leader effects

	// LeaderFxPlayWeatherRanged instantly plays an Impenetrable Fog card from your deck.
	LeaderFxPlayWeatherRanged
	// LeaderFxWeatherClear clears any weather effects (on both sides) currently in play.
	LeaderFxWeatherClear
	// LeaderFxHornSiege doubles the strength of all Siege units, unless a Commander's Horn is already in play on that row
	LeaderFxHornSiege
	// LeaderFxDestroySiege10Plus destroy opponent's strongest Siege units if the strength of that row is 10 or higher.
	LeaderFxDestroySiege10Plus

	//Nilfgaardian Empire leader effects

	// LeaderFxPlayWeatherSiege picks a Torrential Rain card from your deck and play it instantly.
	LeaderFxPlayWeatherSiege
	// LeaderFxLook3OpponentCards looks at 3 random cards from your opponent's hand.
	LeaderFxLook3OpponentCards
	// LeaderFxCancel cancels your opponent's Leader Ability.
	LeaderFxCancel
	// LeaderFxDrawOpponentGrave draws a card from your opponent's discard pile.
	LeaderFxDrawOpponentGrave

	//Scoia'tael leader effects

	// LeaderFxPlayWeatherClose picks a Biting Frost card from your deck and play it instantly.
	LeaderFxPlayWeatherClose
	// LeaderFxDrawExtraCard draws an extra card at the beginning of the battle.
	LeaderFxDrawExtraCard
	// LeaderFxHornRanged doubles the strength of all your Ranged Combat units (unless a Commander's Horn is also present on that row).
	LeaderFxHornRanged
	// LeaderFxDestroyClose10Plus destroys your enemy's strongest Close Combat unit(s) if the combined strength of all his or her Close Combat units is 10 or more.
	LeaderFxDestroyClose10Plus

	//Monsters leader effects

	// LeaderFxPlayWeatherAny picks any weather card from your deck and play it instantly.
	LeaderFxPlayWeatherAny
	// LeaderFxHornClose doubles the strength of all your Close Combat units (unless a Commander's Horn is also present on that row).
	LeaderFxHornClose
	// LeaderFxDiscard2Draw1 discards 2 cards and draw 1 card of your choice from your deck.
	LeaderFxDiscard2Draw1
	// LeaderFxDrawGrave restores a card from your discard pile to your hand.
	LeaderFxDrawGrave
)
