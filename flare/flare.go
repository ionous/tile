package flare

type Map struct {
	Header   Header
	Tilesets Tilesets
	Layers   Layers
}

type Header struct {
	Width, Height         int // size of the layers in tiles
	TileWidth, TileHeight int // size of an individual tile in pixels
	Orientation           string
}

type Tilesets []Tileset

type Tileset struct {
	Path          string
	Width, Height int // size of an individual tile in pixels
	// Margin, Padding int
}

type Layer struct {
	Type string
	Data []uint32
}

type Layers []Layer
