package main

var Relics = []Relic{
	{
		"Lith",
		"A1",
		RelicRewards{
			[]string{
				"Braton Prime Barrel",
				"Forma Blueprint",
				"Vasto Prime Receiver",
			},
			[]string{
				"Saryn Prime Neuroptics Blueprint",
				"Vectis Prime Blueprint",
			},
			[]string{
				"Akstiletto Prime Blueprint",
			},
		},
		true,
	},
	{
		"Lith",
		"A2",
		RelicRewards{
			[]string{
				"Lex Prime Barrel",
				"Forma Blueprint",
				"Valkyr Prime Blueprint",
			},
			[]string{
				"Akbronco Prime Link",
				"Cernos Prime Blueprint",
			},
			[]string{
				"Akstiletto Prime Blueprint",
			},
		},
		true,
	},
}

func getRelics() []Relic {
	return Relics
}
