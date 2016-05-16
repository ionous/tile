package tile

import (
	"fmt"
	"image"
	"image/draw"
	"math"
)

// type Clipper interface {
// 	// width and height of each tile (in pixels)
// 	TileSize() (int, int)
// 	// given a global tile index, find the sheet and cell.
// 	GetTile(TileId) (TileRef, bool)
// 	// given a global sheet index, retreive the sheet.
// 	GetSheet(SheetIndex) (Sheet, bool)
// }

// Clip image data from the passed data
func Clip(tiles TileSets, remap Remap) (
	ret image.Image,
	err error,
) {
	// determine the number of tiles necessary to fit all of the indicies we have used.
	clipper := NewClipper(tiles)
	fcnt := float64(remap.NumTiles())
	tilesWide := int(math.Floor(math.Sqrt(fcnt)))
	tilesHigh := int(math.Ceil(fcnt / float64(tilesWide)))
	//
	tileWidth, tileHeight := tiles.TileSize()
	rect := image.Rect(0, 0, tilesWide*tileWidth, tilesHigh*tileHeight)
	if rect.Empty() {
		err = fmt.Errorf("clipping empty tile")
	} else {
		out := image.NewRGBA(rect)
		outPt := image.Pt(0, 0)

		for _, id := range remap.Tiles() {
			if tile, ok := clipper.GetTile(id); !ok {
				err = fmt.Errorf("couldnt find tile %v", id)
				break
			} else if sheet, ok := clipper.GetSheet(tile.Sheet); !ok {
				err = fmt.Errorf("couldnt find sheet %v", tile)
				break
			} else if srcRect, ok := sheet.Bounds(tile.Cell); !ok {
				err = fmt.Errorf("couldnt find cell %v", tile)
				break
			} else {
				if srcRect.Empty() {
					panic("clipping empty bounds")
				}
				ux, uy := outPt.X*tileWidth, outPt.Y*tileHeight
				outRect := image.Rect(ux, uy, ux+srcRect.Dx(), uy+srcRect.Dy())
				draw.Draw(out, outRect, sheet.Image, srcRect.Min, draw.Src)
				// move to the next output tile
				if x := outPt.X + 1; x < tilesWide {
					outPt.X = x
				} else {
					outPt.X, outPt.Y = 0, outPt.Y+1
				}
			}
		}
		if err == nil {
			ret = out
		}
	}
	return ret, err
}

func NewClipper(ts TileSets) *Clipper {
	cache := NewCache(ts)
	return &Clipper{ts, 0, 0, cache}
}

type Clipper struct {
	tiles      TileSets
	sheetIndex SheetIndex
	startingId TileId
	cache      *Cache
}

// global tile index.
func (clip *Clipper) GetTile(id TileId) (TileRef, bool) {
	return clip.find(id)
}

// global sheet index.
func (clip *Clipper) GetSheet(sheet SheetIndex) (ret Sheet, okay bool) {
	if tile, e := clip.cache.LoadTileSheet(sheet); e == nil {
		ret = tile
		okay = true
	}
	return ret, okay
}

func (clip *Clipper) find(id TileId) (ret TileRef, okay bool) {
	// reset
	index := id
	if index < clip.startingId {
		clip.sheetIndex = 0
		clip.startingId = 0
	}

	// search through all tile sheets to find the matching tile id
	for clip.sheetIndex <= clip.tiles.MaxSheet() {
		if sheet, e := clip.cache.LoadTileSheet(clip.sheetIndex); e != nil {
			panic(e)
			break
		} else {
			tilesWide, tilesHigh := sheet.NumTiles()
			numTiles := tilesWide * tilesHigh
			if subIndex := int(index - clip.startingId); subIndex < numTiles {
				cell := CellIndex(subIndex)
				ret = TileRef{clip.sheetIndex, cell}
				okay = true
				break
			} else {
				clip.startingId = clip.startingId.Advance(numTiles)
				clip.sheetIndex++
			}
		}
	}
	return
}
