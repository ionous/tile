package mosaic

import "image"

type Map struct {
	Layer
}

func NewMap(name string) *Map {
	return &Map{Layer: Layer{Name: name}}
}

type Layer struct {
	Name       string
	Layers     []*Layer         `json:",omitempty"`
	Grid       *Grid            `json:",omitempty"`
	Image      *Image           `json:",omitempty"`
	Shapes     *Shapes          `json:",omitempty"`
	Bounds     *image.Rectangle `json:",omitempty"`
	Properties map[string]bool  `json:",omitempty"`
	// x,y, rotation, etc.
	// effects.
	// pixel size?
}

// FindLayer returns the identified layer, if it exists.
func (m *Map) FindLayer(path []string) *Layer {
	return m.findLayer(path, false)
}

// EnsureLayer returns the identified layer, creating all or any of the hierarchy of layers necessary.
func (m *Map) EnsureLayer(path []string) *Layer {
	return m.findLayer(path, true)
}

// Append a new layer of the passed name; doesn't verify whether or not the name is unique.
func (l *Layer) Append(name string) *Layer {
	newL := &Layer{Name: name}
	l.Layers = append(l.Layers, newL)
	return newL
}

// findLayer unifies the code for FindLayer and EnsureLayer.
func (l *Layer) findLayer(path []string, create bool) (ret *Layer) {
	if len(path) == 0 {
		ret = l
	} else {
		next, path := path[0], path[1:]
		for _, l := range l.Layers {
			if next == l.Name {
				ret = l.findLayer(path, create)
				break
			}
		}
		if ret == nil && create {
			ret = l.Append(next).findLayer(path, create)
		}
	}
	return ret
}
