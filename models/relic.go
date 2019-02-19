package models

type Relic struct {
	Era     string       `json:"era"`
	Name    string       `json:"name"`
	Rewards RelicRewards `json:"rewards"`
	Vaulted bool         `json:"vaulted"`
}

type RelicRewards struct {
	Common   []string `json:"common"`
	Uncommon []string `json:"uncommon"`
	Rare     []string `json:"rare"`
}
