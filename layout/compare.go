package layout

import (
	"bytes"
)

// compare two string arrays
func (a Path) Compare(b []string) (ret int) {
	la, lb := len(a), len(b)
	if sub := la - lb; sub != 0 {
		ret = sub
	} else {
		for i := 0; i < la; i++ {
			sa, sb := a[i], b[i]
			if str := bytes.Compare([]byte(sa), []byte(sb)); str != 0 {
				ret = str
				break
			}
		}
	}
	return
}
