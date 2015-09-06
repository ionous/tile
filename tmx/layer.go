package tmx

import "strconv"

// All <tileset> tags shall occur before the first <layer> tag so that parsers may rely on having the tilesets before needing to resolve tiles.
type TileLayer struct {
	// LayerData:
	//  x and y coordinates are in tiles. Defaults to 0 and can no longer be changed in Tiled Qt.
	// width: The width of the layer in tiles. Traditionally required, but as of Tiled Qt always the same as the map width.
	// height: The height of the layer in tiles. Traditionally required, but as of Tiled Qt 	always the same as the map height.
	LayerData
	Properties *Properties  `xml:"properties,omitempty" json:",omitempty"`
	Data       EmbeddedData `xml:"data"`
}

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
