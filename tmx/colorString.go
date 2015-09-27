package tmx

import (
	"fmt"
	"image/color"
	"strconv"
)

type ColorString string

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
