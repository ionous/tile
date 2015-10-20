package tmx

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTmxObjects(t *testing.T) {
	d, e := NewDecoder(strings.NewReader(tmxObjects)).Decode()
	if assert.NoError(t, e) {
		assert.Equal(t, "1.0", d.Version)
		if out, e := json.MarshalIndent(d, "", "  "); assert.NoError(t, e) {
			assert.NotEmpty(t, out)
			t.Log(string(out))
		}
	}
}

// spring.tmx
var tmxObjects string = `<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" renderorder="right-down" width="27" height="23" tilewidth="32" tileheight="32" backgroundcolor="#000000" nextobjectid="1">
<objectgroup name="one">
  <object id="6" x="71.3333" y="95.3333" width="60.6667" height="58"/>
 </objectgroup>
 <objectgroup name="two">
  <object id="3" x="68" y="163.333" width="61.3333" height="46"/>
  <object id="7" x="65.3333" y="30" width="116.667" height="50.6667"/>
 </objectgroup>
</map>`
