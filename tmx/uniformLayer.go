package tmx

type UniformLayer struct {
	LayerData
	Properties *Properties   `xml:"properties,omitempty" json:",omitempty"`
	Data       *EmbeddedData `xml:"data" json:",omitempty"`
	Image      *Image        `xml:"image,omitempty" json:",omitempty"`
}
