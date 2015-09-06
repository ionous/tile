package tile

import (
	"image"
)

// Sheet of images each of pre-defined size.
type Sheet struct {
	Image                 image.Image
	Path                  string
	TileWidth, TileHeight int // size of an individual tile in pixels
}

func (t Sheet) IsValid() bool {
	return t.Image != nil
}

// Number of tiles contained by the sheet.
func (t Sheet) NumTiles() (tilesX, tilesY int) {
	if t.Image != nil {
		bounds := t.Image.Bounds()
		tilesX, tilesY = bounds.Dx()/t.TileWidth, bounds.Dy()/t.TileHeight
	}
	return
}

// TileXY coordinates in whole tiles based of the passed cell.
func (t Sheet) TileXY(cell CellIndex) (tileX int, tileY int, okay bool) {
	if tilesWide, tilesHigh := t.NumTiles(); tilesWide > 0 {
		tileX, tileY = int(cell)%tilesWide, int(cell)/tilesWide
		okay = tileY < tilesHigh
	}
	return tileX, tileY, okay
}

// Bounds (rectangular) on the image sheet of the passed cell.
func (t Sheet) Bounds(cell CellIndex) (ret image.Rectangle, okay bool) {
	if x, y, ok := t.TileXY(cell); ok {
		ret, okay = t.boundsXY(x, y), true
	}
	return
}

func (t Sheet) boundsXY(tileX, tileY int) (ret image.Rectangle) {
	if tileX >= 0 && tileY >= 0 {
		tilesWide, tilesHigh := t.NumTiles()
		if tileX < tilesWide && tileY < tilesHigh {
			px, py := tileX*t.TileWidth, tileY*t.TileHeight
			ret = image.Rect(px, py, px+t.TileWidth, py+t.TileHeight)
		}
	}
	return
}
