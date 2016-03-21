package gwent

// CardTemplate defines our json structure for loading the cards
type CardTemplate struct {
	ID          int64  `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Power       int    `json:"power,omitempty"`
	Hero        bool   `json:"hero,omitempty"`
	Type        string `json:"type,omitempty"`
	Range       string `json:"range,omitempty"`
	Faction     string `json:"faction,omitempty"`
	Ability     string `json:"ability,omitempty"`
	MaxCount    int    `json:"max,omitempty"`
}

// GetLeaderEffectByName reports LeaderFx from given name
func GetLeaderEffectByName(name string) LeaderFx {
	return LeaderFxNone
}

// GetUnitAbilityByName reports our Ability from given name
func GetUnitAbilityByName(name string) int {
	return AbilityNone
}

// GetFactionByName reports our CardFaction from given name
func GetFactionByName(name string) CardFaction {
	switch name {
	case "nilfgaard":
		return FactionNilfgaard
	case "northern_realms":
		return FactionNorthernRealms
	case "monsters":
		return FactionMonsters
	case "scoiatael":
		return FactionScoiatael
	default:
		return FactionNeutral
	}
}

// Make creates an instance of Card from given template
func (t CardTemplate) Make() Card {
	switch t.Type {
	case "leader":
		return &CardLeader{
			BasicCard: BasicCard{
				CardName:        t.Name,
				CardDescription: t.Description,
			},
			LeaderFaction: GetFactionByName(t.Faction),
			LeaderEffect:  GetLeaderEffectByName(t.Ability),
		}
	}
	return nil
}
