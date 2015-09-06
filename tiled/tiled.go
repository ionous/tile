//
// Package tiled defines constants and types for common Tiled(http://www.mapeditor.org/) data.
//
package tiled

import "github.com/ionous/tile"

// A tile in a tiled map contains a global-tile-id, and the rotation, mirroring of the placed tile.
type Tile uint32

func (t Tile) TileId() tile.TileId {
	data := uint32(t)
	r := data & ^(FLIPPED_HORIZONTALLY_FLAG |
		FLIPPED_VERTICALLY_FLAG |
		FLIPPED_DIAGONALLY_FLAG)
	return tile.TileId(r)
}

// The diagonal is done first, followed by the horizontal and vertical flips.
const (
	FLIPPED_DIAGONALLY_FLAG   uint32 = 0x20000000 // (x/y axis swap)
	FLIPPED_HORIZONTALLY_FLAG        = 0x80000000
	FLIPPED_VERTICALLY_FLAG          = 0x40000000
)
