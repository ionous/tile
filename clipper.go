package tile

import (
	"fmt"
	"github.org/ionous/tile/flare"
	"image"
	"image/draw"
	"math"
)

// Clip image data from the passed (flare) data
func Clip(header flare.Header,
	tilesets flare.Tilesets,
	imagePath string,
	remap Remap,
) (
	img image.Image,
	err error,
) {
	cache := NewCache(imagePath, tilesets)
	fcnt := float64(len(remap.indices))
	tilesWide := int(math.Floor(math.Sqrt(fcnt)))
	tilesHigh := int(math.Ceil(fcnt / float64(tilesWide)))
	out := image.NewRGBA(image.Rect(0, 0, tilesWide*header.TileWidth, tilesHigh*header.TileHeight))
	outPt := image.Pt(0, 0)

	it := MakeIterator(tilesets, cache)
	for _, index := range remap.indices {
		if tileIndex, ok := it.Find(index); !ok {
			err = fmt.Errorf("couldn't find index %d", index)
			break
		} else if tile, e := cache.GetImage(tileIndex.Tile); e != nil {
			err = e
			break
		} else {
			ux, uy := outPt.X*header.TileWidth, outPt.Y*header.TileHeight
			//
			if tileX, tileY, ok := tile.GetXY(tileIndex.Index); !ok {
				err = fmt.Errorf("tile out of range?! %d", tileIndex.Index)
				break
			} else if srcRect := tile.Bounds(tileX, tileY); srcRect.Empty() {
				err = fmt.Errorf("tile empty?! %d, %d", tileX, tileY)
				break
			} else {
				outRect := image.Rect(ux, uy, ux+srcRect.Dx(), uy+srcRect.Dy())
				draw.Draw(out, outRect, tile.Image, srcRect.Min, draw.Src)
				// move to the next output tile
				if x := outPt.X + 1; x < tilesWide {
					outPt.X = x
				} else {
					outPt.X, outPt.Y = 0, outPt.Y+1
				}
			}
		}
	}
	return out, err
}
