package flare

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFlareReader(t *testing.T) {
	d, e := NewMapReader().Read(strings.NewReader(flareData))
	if assert.NoError(t, e); e == nil {
		assert.Equal(t, 17, d.Header.Height)
		assert.Equal(t, 19, d.Header.Width)
		assert.Equal(t, 32, d.Header.TileWidth)
		assert.Equal(t, 32, d.Header.TileHeight)
		assert.Equal(t, "orthogonal", d.Header.Orientation)
		assert.Len(t, d.Tilesets, 13)
		set := d.Tilesets[3]
		assert.Equal(t, "tiles/lpc/inside.png", set.Path)
		assert.Equal(t, 32, set.Width)
		assert.Equal(t, 32, set.Height)
		assert.Len(t, d.Layers, 2)
		l := d.Layers[1]
		assert.Equal(t, "automat-other-door", l.Type)
		assert.Len(t, l.Data, d.Header.Height*d.Header.Width)
		assert.EqualValues(t, 121, l.Data[len(l.Data)-1])
	}
}

var flareData string = `[header]
width=19
height=17
tilewidth=32
tileheight=32
orientation=orthogonal

[tilesets]
tileset=tiles/lpc/grass.png,32,32,0,0
tileset=tiles/lpc/grassalt.png,32,32,0,0
tileset=tiles/lpc/watergrass.png,32,32,0,0
tileset=tiles/lpc/inside.png,32,32,0,0
tileset=tiles/lpc/cabinets.png,32,32,0,0
tileset=tiles/Skorpio's Pack/Interior-Furniture.png,32,32,0,0
tileset=tiles/Skorpio's Pack/Interior-Walls-Beige.png,32,32,0,0
tileset=tiles/Skorpio's Pack/Objects.png,32,32,0,0
tileset=tiles/Skorpio's Pack/Pipes-RustyWalls.png,32,32,0,0
tileset=tiles/lpc/victoria.png,32,32,0,0
tileset=tiles/lpc/dungeon.png,32,32,0,0
tileset=tiles/lpc-sprites/monsters/slime.png,32,32,0,0
tileset=tiles/lpc-sprites/people/princess.png,32,32,0,0

[layer]
type=ground
data=
549,519,519,519,519,519,519,519,519,519,519,519,519,519,519,519,519,519,521,
539,490,490,490,490,564,490,490,490,490,490,490,490,490,490,490,490,491,531,
539,500,500,500,500,574,500,0,0,500,500,500,500,500,500,500,500,501,531,
539,510,510,510,510,584,510,0,0,510,510,510,510,510,510,510,510,511,531,
539,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,492,493,
539,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,492,493,
539,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,492,493,
539,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,492,493,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,492,493,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,492,493,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,492,493,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
539,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
539,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
539,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
539,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,492,493,
569,540,540,540,540,540,540,540,540,540,540,540,540,540,540,540,540,540,541

[layer]
type=automat-other-door
data=
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,375,374,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,359,358,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,375,374,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,391,390,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,121`
