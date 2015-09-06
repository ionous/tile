package tile

import (
	"fmt"
	"github.com/ionous/sashimi/util/errutil"
	"image"
	"os"
)

// Cache of tile sheet images contained by a TileSets.
type Cache struct {
	tileSets TileSets
	img      Sheet
	idx      SheetIndex
}

func NewCache(tileSets TileSets) *Cache {
	return &Cache{tileSets: tileSets}
}

func (c *Cache) LoadTileSheet(idx SheetIndex) (ret Sheet, err error) {
	width, height := c.tileSets.TileSize()

	if c.img.IsValid() && c.idx == idx {
		ret = c.img
	} else if maxSheet := c.tileSets.MaxSheet(); idx > maxSheet {
		err = fmt.Errorf("index of out of range %v > %v", idx, maxSheet)
	} else {
		path := c.tileSets.Path(idx)
		if r, e := os.Open(path); e != nil {
			err = errutil.Append(
				fmt.Errorf("error loading tilesheet %v,%s", idx, path), e)
		} else {
			defer r.Close()
			if img, _, e := image.Decode(r); e != nil {
				err = errutil.Append(
					fmt.Errorf("error decoding tilesheet %v,%s", idx, path), e)
			} else {
				c.img = Sheet{img, path, width, height}
				c.idx = idx
				ret = c.img
			}
		}
	}
	return
}
