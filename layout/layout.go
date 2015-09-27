package layout

import "fmt"

var _ = fmt.Sprint()

// Tiled doesnt have a hierarchy, but it should.
type L struct {
	Name   string
	Object string `json:",omitempty"` // optional
	Layers []L    `json:",omitempty"` // recursive
	Hidden bool   `json:",omitempty"`
}

func Find(layers []L, match []string) (ret *L, retPath []string) {
	return find(match, nil, nil, layers)
}

func find(match, inPath, outPath []string, layers []L) (ret *L, retPath []string) {
	for _, r := range layers {
		currIn := append(inPath, r.InName())
		currOut := append(outPath, r.OutName())
		d := Compare(currIn, match)
		//		fmt.Println(currIn, d)
		if d == 0 {
			ret = &r
			retPath = currOut
			break
		} else if d < 0 {
			if r, p := find(match, currIn, currOut, r.Layers); r != nil {
				ret, retPath = r, p
				break
			}
		}
	}
	return
}

func (l *L) OutName() (ret string) {
	if l.Object != "" {
		ret = l.Object
	} else {
		ret = l.Name
	}
	return ret
}

func (l *L) InName() string {
	return l.Name
}
