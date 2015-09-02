package tile

import (
	"github.com/ionous/tile/flare"
	"sort"
)

// Remap collapses and reindexes a tileset to only what is used.
type Remap struct {
	indices []int
	remap   map[int]int
}

// Remap returns the remapped index if found.
func (r Remap) Remap(originalIndex int) (tile int, okay bool) {
	if originalIndex == 0 {
		tile, okay = 0, true
	} else {
		tile, okay = r.remap[originalIndex]
	}
	return tile, okay
}

// NewRemap from (flare) data.
func NewRemap(layers flare.Layers, filter func(flare.Layer) bool) Remap {
	known := map[int]bool{}
	indices := []int{}
	for _, layer := range layers {
		if filter == nil || filter(layer) {
			for _, r := range layer.Data {
				if r != 0 {
					r := flare.Index(r)
					if exists := known[r]; !exists {
						known[r] = true
						indices = append(indices, r-1)
					}
				}
			}
		}
	}
	remap := make(map[int]int)
	sort.Ints(indices)
	for i, r := range indices {
		remap[r+1] = i + 1
	}
	return Remap{indices, remap}
}
