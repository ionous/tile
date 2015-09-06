package tmx

import (
	"github.com/ionous/tile"
	"path"
)

// implements tile.TileSets
type TileSets struct {
	imagePath             string
	TileWidth, TileHeight int // in pixels
	Sheets                []TileSet
}

// FIX FIX FIX
// we have to consolidate alll of the added sets into a single set. ad/or have a find -- across all tile sheets

func NewTileSets(tmap Map, imagePath string) *TileSets {
	return &TileSets{imagePath, tmap.TileWidth, tmap.TileHeight, tmap.TileSet}
}

// FIX FIX this is probably wrong
// a Max TileSize would be correct
// otherwise, its per sheet
// if there were different tile sizes
// the remapper/clipper wold have to group them by size
func (ts *TileSets) TileSize() (int, int) {
	return ts.TileWidth, ts.TileHeight
}

func (ts *TileSets) MaxSheet() tile.SheetIndex {
	return tile.SheetIndex(len(ts.Sheets) - 1)
}

func (ts *TileSets) Path(idx tile.SheetIndex) string {
	tile := ts.Sheets[idx]
	return path.Join(ts.imagePath, tile.Image.Source)
}
