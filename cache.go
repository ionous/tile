package tile

import (
	"fmt"
	"github.com/ionous/tile/flare"
	"image"
	"os"
	"path"
)

type Cache struct {
	path     string
	tilesets flare.Tilesets
	img      Image
	idx      int
}

func NewCache(path string, tilesets flare.Tilesets) *Cache {
	return &Cache{path: path, tilesets: tilesets}
}

func (cache *Cache) GetImage(idx int) (ret Image, err error) {
	if cache.img.IsValid() && cache.idx == idx {
		ret = cache.img
	} else if numTiles := len(cache.tilesets); idx >= numTiles {
		err = fmt.Errorf("index of out of range %d %d", idx, numTiles)
	} else {
		tile := cache.tilesets[idx]
		full := path.Join(cache.path, tile.Path)
		if r, e := os.Open(full); e != nil {
			err = e
		} else {
			defer r.Close()
			if img, _, e := image.Decode(r); e != nil {
				err = e
			} else {
				cache.img = Image{img, tile.Width, tile.Height}
				cache.idx = idx
				ret = cache.img
			}
		}
	}
	return
}
