package tmx

import (
	"fmt"
	"image/color"
	"strconv"
)

type ImageLayer struct {
	LayerData
	Properties *Properties `xml:"properties,omitempty" json:",omitempty"`
	Image      *Image      `xml:"image,omitempty" json:",omitempty"`
}

type ColorString string

type Image struct {
	Source string        `xml:"source,attr"`                   /// The reference to the tileset image file (Tiled supports most common image formats).
	Trans  ColorString   `xml:"trans,attr" json:",omitempty"`  //trans: Defines a specific color that is treated as transparent (example value: "#FF00FF" for magenta).
	Width  int           `xml:"width,attr" json:",omitempty"`  //: The image width in pixels (optional, used for tile index correction when the image changes)
	Height int           `xml:"height,attr" json:",omitempty"` // : The image height in pixels (optional)
	Format string        `xml:"format,attr" json:",omitempty"` // Used for embedded images, in combination with a data child element. Valid values are file extensions like png, gif, jpg, bmp, etc. (since 0.9)
	Data   *EmbeddedData `xml:"data,omitempty" json:",omitempty"`
}

func (h ColorString) Parse() (ret color.RGBA, err error) {
	if len(h) < 6 {
		err = fmt.Errorf("couldnt parse color string %s", h)
	} else {
		// the tmx format sometimes doesnt have the starting hash.
		if h[0] == '#' {
			h = h[1:]
		}
		if rgb, e := strconv.ParseUint(string(h), 16, 32); e != nil {
			err = e
		} else {
			ret = color.RGBA{
				uint8(rgb >> 16), uint8((rgb >> 8) & 0xFF), uint8(rgb & 0xFF), 0xFF,
			}
		}
	}
	return ret, err
}
