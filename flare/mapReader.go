package flare

import (
	"bufio"
	"fmt"
	"github.com/ionous/sashimi/util/errutil"
	"io"
	"strconv"
	"strings"
)

type MapReader struct {
	data         Map
	currentLayer Layer
	p            subParser
}

func NewMapReader() *MapReader {
	mr := MapReader{}
	mr.p = mr.newParser()
	return &mr
}

func (mr *MapReader) Read(src io.Reader) (data Map, err error) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		if s := strings.TrimSpace(scanner.Text()); len(s) > 0 {
			//fmt.Println(s)
			header := strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]")
			if header {
				header := s[1 : len(s)-1]
				if mr.p.section == nil || !mr.p.section(header) {
					err = unknownText("section", header)
					break
				}
			} else if kv := strings.Split(s, "="); len(kv) == 2 {
				name, cb := mr.p.keyvalue.name, mr.p.keyvalue.cb
				if cb == nil || !cb(kv[0], kv[1]) {
					err = unknownKeyValue(name, kv[0], kv[1])
					break
				}
			} else {
				if mr.p.rawline == nil || !mr.p.rawline(s) {
					err = unknownText("line", s)
					break
				}
			}
		}
	}
	return mr.data, err
}

func (mr *MapReader) newParser() subParser {
	p := subParser{}
	p.section = func(section string) (okay bool) {
		switch section {
		case "header":
			mr.p.keyvalue = keyValueParser{section, mr.headerData}
			okay = true
		case "tilesets":
			mr.p.keyvalue = keyValueParser{section, mr.tilesetData}
			okay = true
		case "layer":
			mr.p.keyvalue = keyValueParser{section, mr.layerData}
			mr.currentLayer = Layer{}
			okay = true
		}
		return okay
	}
	return p
}

func (mr *MapReader) headerData(key, value string) (okay bool) {
	switch key {
	case "width":
		okay = parseInt(&mr.data.Header.Width, value)
	case "height":
		okay = parseInt(&mr.data.Header.Height, value)
	case "tilewidth":
		okay = parseInt(&mr.data.Header.TileWidth, value)
	case "tileheight":
		okay = parseInt(&mr.data.Header.TileHeight, value)
	case "orientation":
		mr.data.Header.Orientation = value
		okay = true
	}
	return okay
}

func (mr *MapReader) tilesetData(key, value string) (okay bool) {
	switch key {
	case "tileset":
		if fields := strings.Split(value, ","); len(fields) == 5 {
			if values, e := parseInts(fields); e == nil {
				// path, width, height
				ts := Tileset{fields[0], values[0], values[1]}
				mr.data.Tilesets = append(mr.data.Tilesets, ts)
				okay = true
			}
		}
	}
	return okay
}

func (mr *MapReader) layerData(key, value string) (okay bool) {
	switch key {
	case "type":
		okay, mr.currentLayer.Type = true, value
	case "data":
		okay, mr.p.rawline = true, func(raw string) (rawOkay bool) {
			if done, e := parseTileLine(&mr.currentLayer, raw); e == nil {
				rawOkay = true
				if done {
					mr.data.Layers = append(mr.data.Layers, mr.currentLayer)
					mr.p.rawline = nil
				}
			}
			return rawOkay
		}
	}
	return okay
}

// parse a line of tile data.
// returns true when the last line of tile data ( no trailing comma )
// has been reached.
func parseTileLine(layer *Layer, s string) (done bool, err error) {
	if nums := strings.Split(s, ","); len(nums) > 0 {
		for _, s := range nums {
			if len(s) > 0 {
				if m, e := strconv.ParseInt(s, 10, 32); e != nil {
					err = e
					break
				} else {
					layer.Data = append(layer.Data, uint32(m))
				}
			}
		}

		if err == nil {
			if last := s[len(s)-1] != ','; last {
				done = true
			}
		}
	}
	return done, err
}

func parseInts(values []string) (ints []int, err error) {
	ints = make([]int, 0, len(values))
	for _, v := range values[1:] {
		if i, e := strconv.Atoi(v); e != nil {
			err = e
			break
		} else {
			ints = append(ints, i)
		}
	}
	return ints, err
}

func parseInt(v *int, value string) (okay bool) {
	if i, err := strconv.Atoi(value); err == nil {
		*v = i
		okay = true
	}
	return okay
}

func unknownKeyValue(reason, key, value string) error {
	return errutil.Func(func() string {
		return fmt.Sprintf("'%s' unknown key-value '%v'='%v'", reason, key, value)
	})
}

func unknownText(reason string, text string) error {
	return errutil.Func(func() string {
		return fmt.Sprintf("%s unknown text %v", reason, text)
	})
}
