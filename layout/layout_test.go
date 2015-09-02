package layout

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIterator(t *testing.T) {
	nodes := [][]string{
		[]string{"ground"},
		[]string{"doors"},
		[]string{"doors", "east"},
		[]string{"doors", "west"},
		[]string{"doors", "north"},
		[]string{"walls"},
	}
	it := NewIterator(root)
	for _, expect := range nodes {
		path := it.Path()
		require.EqualValues(t, expect, path)
		it.Next()
	}
}

var root []L = []L{
	L{Name: "ground"},
	L{Name: "doors", Layers: []L{
		L{Name: "east", Object: "automat-other-door"},
		L{Name: "west", Object: "automat-hall-door"},
		L{Name: "north", Object: "automat-deck-door"},
	}},
	L{Name: "walls", Layers: []L{
		L{Name: "lower"},
		L{Name: "upper"},
	}},
}
