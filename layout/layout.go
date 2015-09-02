package layout

// Tiled doesnt have a hierarchy, but it should.
type L struct {
	Name   string
	Object string `json:",omitempty"` // optional
	Layers []L    `json:",omitempty"` // recursive
	Hidden bool   `json:",omitempty"`
}
