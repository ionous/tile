package tmx

// If there are multiple <tileset> elements, they are in ascending order of their firstgid attribute. The first tileset always has a firstgid value of 1 and it can be assumed that there are no gaps in the valid range of global tile IDs.
type TileSet struct {
	FirstGid   int    `xml:"firstgid,attr"`                  // 1.0 The first global tile ID of this tileset (this global ID maps to the first tile in this tileset).
	Source     string `xml:"source,attr" json:",omitempty"`  // If this tileset is stored in an external TSX (Tile Set XML) file, this attribute refers to that file. That TSX file has the same structure as the <tileset> element described here. (There is the firstgid attribute missing and this source attribute is also not there. These two attributes are kept in the TMX map, since they are map specific.)
	Name       string `xml:"name,attr"`                      // The name of this tileset.
	TileWidth  int    `xml:"tilewidth,attr"`                 // The (maximum) width of the tiles in this tileset.
	TileHeight int    `xml:"tileheight,attr"`                // The (maximum) height of the tiles in this tileset.
	Spacing    int    `xml:"spacing,attr" json:",omitempty"` // The spacing in pixels between the tiles in this tileset (applies to the tileset image).
	Margin     int    `xml:"margin,attr" json:",omitempty"`  // The margin around the tiles in this tileset (applies to the tileset image).

	TileOffset   *TileOffset   `xml:"tileoffset,omitempty" json:",omitempty"`
	Properties   *Properties   `xml:"properties,omitempty" json:",omitempty"`
	Image        *Image        `xml:"image,omitempty" json:",omitempty"`
	TerrainTypes *TerrainTypes `xml:"terraintypes,omitempty" json:",omitempty"`
	Tile         []TileSetTile `xml:"tile" json:",omitempty"`
}

type TileOffset struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"` // 1.0
}

type TileSetTile struct {
	// id: The local tile ID within its tileset.
	// terrain: Defines the terrain type of each corner of the tile, given as comma-separated indexes in the terrain types array in the order top-left, top-right, bottom-left, bottom-right. Leaving out a value means that corner has no terrain. (optional) (since 0.9)
	// probability: A percentage indicating the probability that this tile is chosen when it competes with others while editing with the terrain tool. (optional) (since 0.9)
	Properties  *Properties  `xml:"properties,omitempty"`
	Image       *Image       `xml:"image,omitempty"`
	ObjectGroup *ObjectGroup `xml:"objectgroup,omitempty"`
	Animation   *Animation   `xml:"animation,omitempty"`
}

// Animation contains a list of animation frames.
// As of Tiled 0.10, each tile can have exactly one animation associated with it. In the future, there could be support for multiple named animations on a tile.
type Animation struct {
	Frame []Frame `xml:"frame"`
	// Can contain: frame
}

type Frame struct {
	TileId   string  `xml:"tile,attr"`     // TileId of a tile within the parent tileset.
	Duration float32 `xml:"duration,attr"` // Duration in milliseconds this frame should be displayed before advancing to the next frame.
}
