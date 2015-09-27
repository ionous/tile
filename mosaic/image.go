package mosaic

type Image struct {
	// scale?
	// offset?
	Source string
}

func NewImage(src string) *Image {
	return &Image{Source: src}
}
