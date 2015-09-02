package tmx

import (
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strings"
)

type Encoding string

//Whatever format you choose for your layer data, you will always end up with so called "global tile IDs" (gids). They are global, since they may refer to a tile from any of the tilesets used by the map. In order to find out from which tileset the tile is you need to find the tileset with the highest firstgid that is still lower or equal than the gid. The tilesets are always stored with increasing firstgids.
type EmbeddedData struct {
	Encoding    string `xml:"encoding,attr"` //: The encoding used to encode the tile layer data. When used, it can be "base64" and "csv" at the moment.
	Compression string `xml:"compression,attr"`
	// compression: The compression used to compress the tile layer data. Tiled Qt supports "gzip" and "zlib".
	// When no encoding or compression is given, the tiles are stored as individual XML tile elements. Next to that, the easiest format to parse is the "csv" (comma separated values) format.
	Data string `xml:",innerxml"`
	//The base64-encoded and optionally compressed layer data is somewhat more complicated to parse. First you need to base64-decode it, then you may need to decompress it. Now you have an array of bytes, which should be interpreted as an array of unsigned 32-bit integers using little-endian byte ordering.
}

func (d EmbeddedData) Decompress(out []uint32) (err error) {
	if d.Encoding != "base64" {
		err = fmt.Errorf("encoding format '%s' not supported", d.Encoding)
	} else if d.Compression != "zlib" {
		err = fmt.Errorf("compression format '%s' not supported", d.Compression)
	} else {
		if r, e := zlib.NewReader(base64.NewDecoder(base64.StdEncoding, strings.NewReader(d.Data))); e != nil {
			err = e
		} else {
			defer r.Close()
			err = binary.Read(r, binary.LittleEndian, &out)
		}
	}
	return err
}
