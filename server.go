package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AbilityStats represents any statistics pertaining to a Champion's Ability
type AbilityStats struct {
	Type  string `json:"type"`  // Type is the kind of statistic this ability has - ex. Damage, Attack Speed Slow, Storm Duration
	Value string `json:"value"` // Value is the associated value with the kind of statistic - ex. 500, 5%, 3s
}

// Ability represents any abilities a Champion may have
type Ability struct {
	Name         string         `json:"name"`        // Name is the name of the ability
	Description  string         `json:"description"` // Description is a brief explanation of the ability
	Type         string         `json:"type"`        // Type describes whether this ability is Active or Passive
	ManaCost     int            `json:"manaCost"`    // ManaCost describes how much mana this ability costs to cast
	ManaStart    int            `json:"manaStart"`   // ManaStart is how much mana the champion starts with - I think? Maybe this should just be under Champion rather than Ability
	AbilityStats []AbilityStats `json:"stats"`       // AbilityStats represents any statistics pertaining to a Champion's Ability
}

// Offense represents the base offensive statistics for a Champion
type Offense struct {
	Damage      int     `json:"damage"`      // Damage is a base value for how much an auto-attack from this Champion does
	AttackSpeed float32 `json:"attackSpeed"` // AttackSpeed is the value at which this Champion attacks per second
	Dps         int     `json:"dps"`         // DPS is the Damage-Per-Second, can be calculated by multiplying Damage * AttackSpeed
	Range       int     `json:"range"`       // Range is the number of hexes a unit can attack from
}

// Defense represents the base defensive statistics for a Champion
type Defense struct {
	Health      int `json:"health"`      // Health is the base life value for a Champion
	Armor       int `json:"armor"`       // Armor is the base armor value for a Champion
	MagicResist int `json:"magicResist"` // MagicResist is the base magic resistance value for a Champion.
}

// ChampionStats represents the base Offense and Defense statistics for a Champion
type ChampionStats struct {
	Offense Offense `json:"offense"` // Offense represents the base offensive statistics for a Champion
	Defense Defense `json:"defense"` // Defense represents the base defensive statistics for a Champion
}

// Champion represents a unit from Teamfight Tactics
type Champion struct {
	Key              string        `json:"key"`     // Key is a logical key which represents this Champion object
	Name             string        `json:"name"`    // Name is the name of this Champion unit
	Origin           []string      `json:"origin"`  // Origin is where this Champion came from
	Class            []string      `json:"class"`   // Class is the class(es) to which this champion belongs to
	Cost             int           `json:"cost"`    // Cost is the gold value of this Champion
	Ability          Ability       `json:"ability"` // Ability represents any abilities a Champion may have
	ChampionStats    ChampionStats `json:"stats"`   // ChampionStats represents the base Offense and Defense statistics for a Champion
	RecommendedItems []string      `json:"items"`   // RecommendedItems is a list of curated items which should work well on this Champion
}

func main() {
	// HandlerFunction to test Champions struct and get familiar with JSON + Go
	http.HandleFunc("/champions/aatrox", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		aatrox := Champion{
			Key:    "Aatrox",
			Name:   "Aatrox",
			Origin: []string{"Demon", "Pirate"},
			Class:  []string{"Blademaster", "Gunslinger"},
			Cost:   3,
			Ability: Ability{
				Name:        "The Darkin Blade",
				Description: "Aatrox cleaves the area in front of him, dealing damage to enemies inside it.",
				Type:        "Active",
				ManaCost:    100,
				ManaStart:   0,
				AbilityStats: []AbilityStats{
					{
						Type:  "Damage",
						Value: "350 / 575 / 850",
					},
					{
						Type:  "Storm Duration",
						Value: "8s",
					},
				},
			},
			ChampionStats: ChampionStats{
				Offense: Offense{
					Damage:      65,
					AttackSpeed: 0.65,
					Dps:         42,
					Range:       1,
				},
				Defense: Defense{
					Health:      750,
					Armor:       25,
					MagicResist: 20,
				},
			},
			RecommendedItems: []string{
				"titanichydra",
				"phantomdancer",
				"dragonsclaw",
			},
		}

		json.NewEncoder(w).Encode(aatrox) // Encode our object into JSON before sending the response
	})

	// HandlerFunction to handle creating new champions in the Database
	http.HandleFunc("/champions/create", func(w http.ResponseWriter, r *http.Request) {
		var champ Champion                     // Initialize a variable to hold the decoded JSON
		json.NewDecoder(r.Body).Decode(&champ) // Decode the JSON from the request body

		fmt.Fprintf(w, "Name: %s, Ability Description: %s", champ.Name, champ.Ability.Description) // Send back a test message to ensure decoding occured successfully
	})

	http.ListenAndServe(":8080", nil) // Start HTTP server listening on port 8080
}
