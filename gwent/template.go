package gwent

type CardTemplate struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Power int `json:"power"`
	Hero bool `json:"hero"`
	Type string `json:"type"`
	Range string `json:"range"`
	Faction string `json:"faction"`
	Ability string `json:"ability"`
}

func GetLeaderEffectByName(name string) int {
	return LEADER_FX_NONE
}

func GetUnitAbilityByName(name string) int {
	return ABILITY_NONE
}

func GetFactionByName(name string) int {
	switch(name){
	case "nilfgaard":
		return FACTION_NILFGAARD
	case "northern_realms":
		return FACTION_NORTHERN_REALMS
	case "monsters":
		return FACTION_MONSTERS
	case "scoiatael":
		return FACTION_SCOIATAEL
	default:		
		return FACTION_NEUTRAL
	}
}

func (t CardTemplate) Make() Card {
	switch(t.Type) {
	case "leader":
		return &CardLeader{
			Name: t.Name,
			Faction: GetFactionByName(t.Faction),
			Description: t.Description,
			Effect: GetLeaderEffectByName(t.Ability),
		}	
	}
	return nil
} 