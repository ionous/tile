package tile

import (
	"github.org/ionous/tile/flare"
)

type Iterator struct {
	tiles        flare.Tilesets
	tileIdx, ofs int
	cache        *Cache
}

// the flare iterator needs the image cache:
// the number of tiles per image isnt recorded into the data file.
func MakeIterator(tiles flare.Tilesets, cache *Cache) Iterator {
	return Iterator{tiles: tiles, cache: cache}
}

type TileIndex struct {
	Tile, Index int
}

func (it *Iterator) Find(index int) (ret TileIndex, okay bool) {
	// reset
	if index < it.ofs {
		it.tileIdx = 0
		it.ofs = 0
	}

	// search through all tile sheets
	for it.tileIdx < len(it.tiles) {
		if tile, e := it.cache.GetImage(it.tileIdx); e != nil {
			panic(e)
			break
		} else {
			tilesWide, tilesHigh := tile.NumTiles()
			numTiles := tilesWide * tilesHigh
			if subIndex := index - it.ofs; subIndex < numTiles {
				okay, ret = true, TileIndex{it.tileIdx, subIndex}
				break
			} else {
				it.ofs += numTiles
				it.tileIdx++
			}
		}
	}
	return
}
