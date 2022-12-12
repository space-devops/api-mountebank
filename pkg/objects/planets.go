package objects

type Planets struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PlanetList struct {
	Planets []Planets `json:"planets"`
}

type Planet struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Radius     Radius   `json:"equatorial_radius"`
	Mass       Mass     `json:"mass"`
	Volume     Volume   `json:"volume"`
	Density    Density  `json:"density"`
	Satellites []string `json:"satellites"`
}

type Radius struct {
	Value float64 `json:"value"`
	Unit  string  `json:"metric_unit"`
}

type Mass struct {
	Value float64 `json:"value"`
	Unit  string  `json:"metric_unit"`
}

type Volume struct {
	Value float64 `json:"value"`
	Unit  string  `json:"metric_unit"`
}

type Density struct {
	Value float64 `json:"value"`
	Unit  string  `json:"metric_unit"`
}
