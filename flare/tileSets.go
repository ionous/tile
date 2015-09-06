package flare

import (
	"github.com/ionous/tile"
	"path"
)

type TileSets struct {
	header   Header
	tilesets Tilesets
	path     string
}

// return an object supplying the tile.TileSet inteface
func NewTileSets(header Header, ts Tilesets, path string) *TileSets {
	return &TileSets{header, ts, path}
}

func (ts *TileSets) TileSize() (int, int) {
	// tile := ts.tilesets[idx]
	// return tile.Width, tile.Height
	return ts.header.TileWidth, ts.header.TileHeight
}

func (ts *TileSets) MaxSheet() tile.SheetIndex {
	return tile.SheetIndex(len(ts.tilesets) - 1)
}

func (ts *TileSets) Path(idx tile.SheetIndex) string {
	tile := ts.tilesets[idx]
	return path.Join(ts.path, tile.Path)
}
