package gwent

type CardFaction int

const (
	FactionNeutral        = CardFaction(iota)
	FactionNilfgaard      //Wins any round ending in draw
	FactionNorthernRealms //Grants an extra card upon winning round
	FactionMonsters       //Keeps random Unit Card out after each round
	FactionScoiatael      //Decides who takes first turn
)

type CardRange int

const (
	RangeNone = CardRange(iota)
	RangeClose
	RangeRanged
	RangeSiege
)

type CardType int

const (
	TypeBasic = CardType(iota)
	TypeHorn
	TypeScorch
	TypeWeather
	TypeLeader
)
