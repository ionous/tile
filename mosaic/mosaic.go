package mosaic

import "image"

type Map struct {
	Layer
}

func NewMap(name string) *Map {
	return &Map{Layer: Layer{Name: name}}
}

type Layer struct {
	Name   string
	Layers []*Layer `json:",omitempty"`
	Grid   *Grid    `json:",omitempty"`
	Image  *Image   `json:",omitempty"`
	Bounds image.Rectangle
	Hidden bool
	// x,y, rotation, etc.
	// effects.
	// pixel size?
}

func (m *Map) FindLayer(path []string) *Layer {
	return m.findLayer(path, false)
}

func (m *Map) EnsureLayer(path []string) *Layer {
	return m.findLayer(path, true)
}

func (l *Layer) Append(name string) *Layer {
	newL := &Layer{Name: name}
	l.Layers = append(l.Layers, newL)
	return newL
}

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
