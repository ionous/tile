package tmx

// The tilewidth and tileheight properties determine the general grid size of the map. The individual tiles may have different sizes. Larger tiles will extend at the top and right (anchored to the bottom left).
// A map contains three different kinds of layers. Tile layers were once the only type, and are simply called layer, object layers have the objectgroup tag and image layers use the imagelayer tag. The order in which these layers appear is the order in which the layers are rendered by Tiled.
type Map struct {
	//Map         xml.XMLName `xml:"map"`
	Version     string      `xml:"version,attr"`         // 1.0
	Orientation string      `xml:"orientation,attr"`     //Orientation supports "orthogonal", "isometric" and "staggered" (since 0.9) at the moment.
	Width       int         `xml:"width,attr"`           // Width in tiles.
	Height      int         `xml:"height,attr"`          // Height in tiles.
	TileWidth   int         `xml:"tilewidth,attr"`       // Width of a tile. ( in pixels )
	TileHeight  int         `xml:"tileheight,attr"`      // Height of a tile. ( in pixels )
	Background  ColorString `xml:"backgroundcolor,attr"` // Background color of the map. (since 0.9, optional)
	// order in which tiles on tile layers are rendered. Valid values are right-down (the default), right-up, left-down and left-up. In all cases, the map is drawn row-by-row. (since 0.10, but only supported for orthogonal maps at the moment)
	Order string `xml:"renderorder,attr"`

	Properties *Properties `xml:"properties,omitempty" json:",omitempty"`
	TileSet    []TileSet   `xml:"tileset" json:",omitempty"`

	// the txm format has three different tags for layers
	// ( <layer>, <imagelayer>, and <objectgroup> )
	// if we called out the tags individually we would loose their inter-ordering; instead we load them all into a single "uniform layer"
	// meaning: we cant serialize directly back out...
	Layers []UniformLayer `xml:",any" json:",omitempty"`
}

func (m *Map) ReverseLayers() {
	a := m.Layers
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	m.Layers = a
}
