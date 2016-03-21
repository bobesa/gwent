package gwent

// CardFaction defines faction of card
type CardFaction int

// List of card factions
const (
	FactionNeutral        = CardFaction(iota)
	FactionNilfgaard      //Wins any round ending in draw
	FactionNorthernRealms //Grants an extra card upon winning round
	FactionMonsters       //Keeps random Unit Card out after each round
	FactionScoiatael      //Decides who takes first turn
)

// CardRange defines range of card
type CardRange int

// List of card ranges
const (
	RangeNone = CardRange(iota)
	RangeClose
	RangeRanged
	RangeSiege
)

// CardType defines type of card
type CardType int

// List of card types
const (
	TypeBasic = CardType(iota)
	TypeHorn
	TypeScorch
	TypeWeather
	TypeLeader
)
