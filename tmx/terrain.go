package tmx

//TerrainTypes defines an array of terrain types, which can be referenced from the terrain attribute of the tile element.
type TerrainTypes struct {
	Terrain []Terrain
}

type Terrain struct {
	Name       string      `xml:"name,attr"` // Name of the image layer.
	TileId     string      `xml:"tile,attr"` // Tile-id of the tile that represents the terrain visually.
	Properties *Properties `xml:"properties"`
}
