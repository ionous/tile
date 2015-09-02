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

// The diagonal is done first, followed by the horizontal and vertical flips.
const (
	FLIPPED_DIAGONALLY_FLAG   uint32 = 0x20000000 // (x/y axis swap)
	FLIPPED_HORIZONTALLY_FLAG        = 0x80000000
	FLIPPED_VERTICALLY_FLAG          = 0x40000000
)

func Index(data uint32) int {
	r := data & ^(FLIPPED_HORIZONTALLY_FLAG |
		FLIPPED_VERTICALLY_FLAG |
		FLIPPED_DIAGONALLY_FLAG)
	return int(r)
}
