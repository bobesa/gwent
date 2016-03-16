package gwent

type CardTemplate struct {
	Id          int64  `json:"id"`
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

func GetLeaderEffectByName(name string) int {
	return LEADER_FX_NONE
}

func GetUnitAbilityByName(name string) int {
	return AbilityNone
}

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
