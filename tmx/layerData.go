package tmx

import "strconv"

type LayerData struct {
	// Name of the image layer.
	Name string `xml:"name,attr"`
	// X position of the image layer
	X int `xml:"x,attr" json:",omitempty"`
	// Y position of the image layer
	Y int `xml:"y,attr" json:",omitempty"`
	// Width of the image layer in tiles. Meaningless.
	Width int `xml:"width,attr"`
	// Height of the image layer in tiles. Meaningless.
	Height int `xml:"height,attr"`
	// Opacity of the layer as a value from 0 to 1. Defaults to 1.
	Opacity Opacity `xml:"opacity,attr" json:",omitempty"`
	// Visible if the layer is shown (1) or hidden (0). Defaults to 1.
	Visible Visible `xml:"visible,attr" json:",omitempty"`
}

type Visible string
type Opacity string

func (v Visible) Value() bool {
	return v == "" || v == "1"
}

func (o Opacity) Value() float64 {
	val := 1.0
	if o != "" {
		if f, err := strconv.ParseFloat(string(o), 64); err != nil {
			val = f
		}
	}
	return val
}
