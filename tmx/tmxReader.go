package tmx

import (
	"encoding/xml"
	"io"
)

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{xml.NewDecoder(r)}
}

type Decoder struct {
	xml *xml.Decoder
}

func (d *Decoder) Decode() (ret Map, err error) {
	err = d.xml.Decode(&ret)
	return ret, err
}
