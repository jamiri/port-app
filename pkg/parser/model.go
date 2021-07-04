package parser

type Port struct {
	Name        string      `json:"name"`
	Coordinates Coordinates `json:"coordinates"`
	City        string      `json:"city"`
	Country     string      `json:"country"`
	Alias       Alias       `json:"alias"`
	Regions     Regions     `json:"regions"`
	Province    string      `json:"province"`
	Timezone    string      `json:"timezone"`
	Unlocs      Unlocs      `json:"unlocs"`
}

type Coordinates []float32
type Alias []string
type Regions []string
type Unlocs []string
