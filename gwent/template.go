package gwent

type CardTemplate struct {
	Id int64 `json:"id"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Power int `json:"power,omitempty"`
	Hero bool `json:"hero,omitempty"`
	Type string `json:"type,omitempty"`
	Range string `json:"range,omitempty"`
	Faction string `json:"faction,omitempty"`
	Ability string `json:"ability,omitempty"`
	MaxCount int `json:"max,omitempty"`
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