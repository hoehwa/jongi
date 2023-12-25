package listfancy

import (
	"sync"
)

type randomItemGenerator struct {
	titles     []string
	descs      []string
	titleIndex int
	descIndex  int
	mtx        *sync.Mutex
	shuffle    *sync.Once
}

func (r *randomItemGenerator) reset(titles []string, descs []string) {
	r.mtx = &sync.Mutex{}
	r.shuffle = &sync.Once{}

	r.titles = titles // 	"Artichoke",
	// 	"Baking Flour",
	// 	"Bananas",
	// 	"Barley",
	// 	"Bean Sprouts",
	// 	"Bitter Melon",
	// 	"Black Cod",
	// 	"Blood Orange",
	// 	"Brown Sugar",
	// 	"Cashew Apple",
	// 	"Cashews",
	// 	"Cat Food",
	// 	"Coconut Milk",
	// 	"Cucumber",
	// 	"Curry Paste",
	// 	"Currywurst",
	// 	"Dill",
	// 	"Dragonfruit",
	// 	"Dried Shrimp",
	// 	"Eggs",
	// 	"Fish Cake",
	// 	"Furikake",
	// 	"Garlic",
	// 	"Gherkin",
	// 	"Ginger",
	// 	"Granulated Sugar",
	// 	"Grapefruit",
	// 	"Green Onion",
	// 	"Hazelnuts",
	// 	"Heavy whipping cream",
	// 	"Honey Dew",
	// 	"Horseradish",
	// 	"Jicama",
	// 	"Kohlrabi",
	// 	"Leeks",
	// 	"Lentils",
	// 	"Licorice Root",
	// 	"Meyer Lemons",
	// 	"Milk",
	// 	"Molasses",
	// 	"Muesli",
	// 	"Nectarine",
	// 	"Niagamo Root",
	// 	"Nopal",
	// 	"Nutella",
	// 	"Oat Milk",
	// 	"Oatmeal",
	// 	"Olives",
	// 	"Papaya",
	// 	"Party Gherkin",
	// 	"Peppers",
	// 	"Persian Lemons",
	// 	"Pickle",
	// 	"Pineapple",
	// 	"Plantains",
	// 	"Pocky",
	// 	"Powdered Sugar",
	// 	"Quince",
	// 	"Radish",
	// 	"Ramps",
	// 	"Star Anise",
	// 	"Sweet Potato",
	// 	"Tamarind",
	// 	"Unsalted Butter",
	// 	"Watermelon",
	// 	"Weißwurst",
	// 	"Yams",
	// 	"Yeast",
	// 	"Yuzu",
	// 	"Snow Peas",
	// }

	r.descs = descs
}

func (r *randomItemGenerator) next(titles []string, descs []string) item {
	if r.mtx == nil {
		r.reset(titles, descs)
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	i := item{
		title:       r.titles[r.titleIndex],
		description: r.descs[r.descIndex],
	}

	r.titleIndex++
	if r.titleIndex >= len(r.titles) {
		r.titleIndex = 0
	}

	r.descIndex++
	if r.descIndex >= len(r.descs) {
		r.descIndex = 0
	}

	return i
}
