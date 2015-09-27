package tmx

type Image struct {
	Source string        `xml:"source,attr"`                   /// The reference to the tileset image file (Tiled supports most common image formats).
	Trans  ColorString   `xml:"trans,attr" json:",omitempty"`  //trans: Defines a specific color that is treated as transparent (example value: "#FF00FF" for magenta).
	Width  int           `xml:"width,attr" json:",omitempty"`  //: The image width in pixels (optional, used for tile index correction when the image changes)
	Height int           `xml:"height,attr" json:",omitempty"` // : The image height in pixels (optional)
	Format string        `xml:"format,attr" json:",omitempty"` // Used for embedded images, in combination with a data child element. Valid values are file extensions like png, gif, jpg, bmp, etc. (since 0.9)
	Data   *EmbeddedData `xml:"data,omitempty" json:",omitempty"`
}
