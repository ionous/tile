package mosaic

type Map struct {
	Layer
}

func NewMap(name string) *Map {
	return &Map{Layer: Layer{Name: "automat"}}
}

type Layer struct {
	Name   string
	Layers []*Layer `json:",omitempty"`
	Grid   *Grid    `json:",omitempty"`
	// x,y, rotation, etc.
	// effects.
}

func (m *Map) FindLayer(path []string) *Layer {
	return m.findLayer(path)
}

func (l *Layer) Append(name string) *Layer {
	newL := &Layer{Name: name}
	l.Layers = append(l.Layers, newL)
	return newL
}

func (l *Layer) findLayer(path []string) (ret *Layer) {
	if len(path) == 0 {
		ret = l
	} else {
		next, path := path[0], path[1:]
		for _, l := range l.Layers {
			if next == l.Name {
				ret = l.findLayer(path)
				break
			}
		}
	}
	return ret
}
