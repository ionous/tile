package tmx

type LayerType string

const (
	EmptyLayer  LayerType = "empty layer"
	TileLayer             = "tile layer"
	ImageLayer            = "image layer"
	ObjectLayer           = "object layer"
)

// the contents of the tmx layer, imagelayer, and objectgroup tags
type UniformLayer struct {
	LayerData
	Properties  *Properties   `xml:"properties,omitempty" json:",omitempty"`
	Data        *EmbeddedData `xml:"data" json:",omitempty"`
	Image       *Image        `xml:"image,omitempty" json:",omitempty"`
	ObjectGroup []Object      `xml:"object,omitempty" json:",omitempty"`
}

func (u UniformLayer) LayerType() (ret LayerType) {
	switch {
	case u.Data != nil:
		ret = TileLayer
	case u.Image != nil:
		ret = ImageLayer
	case len(u.ObjectGroup) > 0:
		ret = ObjectLayer
	default:
		ret = EmptyLayer
	}
	return ret
}
