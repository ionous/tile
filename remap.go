package tile

import (
	"sort"
)

// Remap collapses and reindexes a tileset to only what is used.
type Remap struct {
	tiles []TileId
	remap map[TileId]TileId
}

// Remap returns the remapped index if found.
func (r Remap) RemapId(originalId TileId) (tile TileId, okay bool) {
	if originalId == 0 {
		tile, okay = 0, true
	} else {
		tile, okay = r.remap[originalId]
	}
	return tile, okay
}

// NumTiles in use by the remapper.
func (r Remap) NumTiles() int {
	return len(r.tiles)
}

// Tiles contains the list of tiles (global tile tiles) in use by the remapper.
// The list is sorted from smallest to largest index.
func (r Remap) Tiles() []TileId {
	return r.tiles
}

// RemapBuilder helps construct a Remap from multiple tile layers.
type RemapBuilder struct {
	known map[TileId]bool
	tiles []TileId
}

// NewRemapBuilder
func NewRemapBuilder() RemapBuilder {
	known := map[TileId]bool{}
	tiles := []TileId{}
	return RemapBuilder{known, tiles}
}

// AddTile to the pending remap.
func (b *RemapBuilder) AddTile(tile TileId) {
	if tile != 0 {
		if exists := b.known[tile]; !exists {
			b.known[tile] = true
			// FIX: why -1?
			b.tiles = append(b.tiles, tile-1)
		}
	}
}

// RemapTiles returns a Remap consisting of all tiles added so far.
func (b *RemapBuilder) RemapIds() Remap {
	remap := make(map[TileId]TileId)
	tiles := make([]TileId, len(b.tiles))
	copy(tiles, b.tiles)
	sort.Sort(TileSlice(tiles))
	//
	for i, r := range tiles {
		// FIX: why +1 on left-side?
		remap[r+1] = TileId(i + 1)
	}
	return Remap{tiles, remap}
}

type TileSlice []TileId

func (p TileSlice) Len() int           { return len(p) }
func (p TileSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p TileSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
