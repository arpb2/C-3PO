package level

const (
	OrientationNorth Orientation = "N"
	OrientationSouth Orientation = "S"
	OrientationEast  Orientation = "E"
	OrientationWest  Orientation = "W"

	AxisX Axis = "x"
	AxisY Axis = "y"
)

type Definition struct {
	Origin           Origin      `json:"origin"`
	Destination      Destination `json:"destination"`
	MinimalDimension Dimension   `json:"minimal_dimensions"`
	Collectibles     []Element   `json:"collectibles"`
	Gates            []Gate      `json:"gates"`
	Pads             []Element   `json:"pads"`
}

type Orientation string
type Axis string

type Origin struct {
	Position    Position    `json:"position"`
	Orientation Orientation `json:"orientation"`
}

type Position struct {
	X    int  `json:"x"`
	Y    int  `json:"y"`
	Axis Axis `json:"axis,omitempty"`
}

type Destination struct {
	Position   Position    `json:"position"`
	Active     bool        `json:"active"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type Dimension struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

type Element struct {
	Position Position `json:"position"`
	Type     string   `json:"type"`
}

type Gate struct {
	Element
	OpeningValue string `json:"opening_value,omitempty"`
}
