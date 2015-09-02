package tile

import (
	"image"
)

type Image struct {
	Image                 image.Image
	TileWidth, TileHeight int // size of an individual tile in pixels
}

func (t Image) IsValid() bool {
	return t.Image != nil
}

func (t Image) NumTiles() (tilesX, tilesY int) {
	if t.Image != nil {
		bounds := t.Image.Bounds()
		tilesX, tilesY = bounds.Dx()/t.TileWidth, bounds.Dy()/t.TileHeight
	}
	return
}

func (t Image) GetXY(index int) (tileX int, tileY int, okay bool) {
	if tilesWide, tilesHigh := t.NumTiles(); tilesWide > 0 {
		tileX, tileY = index%tilesWide, index/tilesWide
		okay = tileY < tilesHigh
	}
	return tileX, tileY, okay

}

func (t Image) Bounds(tileX, tileY int) (ret image.Rectangle) {
	if tileX >= 0 && tileY >= 0 {
		tilesWide, tilesHigh := t.NumTiles()
		if tileX < tilesWide && tileY < tilesHigh {
			px, py := tileX*t.TileWidth, tileY*t.TileHeight
			ret = image.Rect(px, py, px+t.TileWidth, py+t.TileHeight)
		}
	}
	return
}
