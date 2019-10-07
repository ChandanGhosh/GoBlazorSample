package models

type Weather struct {
	Date			string		`json:"date"`
	TemperatureC	int			`json:"temperatureC"`
	TemperatureF	int			`json:"temperatureF"`
	Summary			string		`json:"summary"`
}

type Weathers struct {
	Weathers	[]Weather		`json:"weathers"`
}
