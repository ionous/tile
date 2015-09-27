package tmx

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTmxReader(t *testing.T) {
	d, e := NewDecoder(strings.NewReader(tmxData)).Decode()
	if assert.NoError(t, e) {
		assert.Equal(t, "1.0", d.Version)
		if out, e := json.MarshalIndent(d, "", "  "); assert.NoError(t, e) {
			assert.NotEmpty(t, out)
			t.Log(string(out))
		}
	}
}

// spring.tmx
var tmxData string = `<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" renderorder="right-down" width="27" height="23" tilewidth="32" tileheight="32" backgroundcolor="#000000" nextobjectid="1">
 <tileset firstgid="1" name="grass" tilewidth="32" tileheight="32" tilecount="18">
  <image source="../../../../../../alice-app/data/tiles/lpc/grass.png" width="96" height="192"/>
 </tileset>
 <tileset firstgid="19" name="grassalt" tilewidth="32" tileheight="32" tilecount="18">
  <image source="../../../../../../alice-app/data/tiles/lpc/grassalt.png" width="96" height="192"/>
 </tileset>
 <tileset firstgid="37" name="watergrass" tilewidth="32" tileheight="32" tilecount="18">
  <image source="../../../../../../alice-app/data/tiles/lpc/watergrass.png" width="96" height="192"/>
 </tileset>
 <tileset firstgid="55" name="inside" tilewidth="32" tileheight="32" tilecount="100">
  <image source="../../../../../../alice-app/data/tiles/lpc/inside.png" width="320" height="320"/>
 </tileset>
 <tileset firstgid="155" name="cabinets" tilewidth="32" tileheight="32" tilecount="78">
  <image source="../../../../../../alice-app/data/tiles/lpc/cabinets.png" width="192" height="416"/>
 </tileset>
 <tileset firstgid="233" name="Interior-Furniture" tilewidth="32" tileheight="32" tilecount="256">
  <image source="../../../../../../alice-app/data/tiles/Skorpio's Pack/Interior-Furniture.png" width="512" height="512"/>
 </tileset>
 <tileset firstgid="489" name="Interior-Walls-Beige" tilewidth="32" tileheight="32" tilecount="100">
  <image source="../../../../../../alice-app/data/tiles/Skorpio's Pack/Interior-Walls-Beige.png" width="320" height="320"/>
 </tileset>
 <tileset firstgid="589" name="Objects" tilewidth="32" tileheight="32" tilecount="64">
  <image source="../../../../../../alice-app/data/tiles/Skorpio's Pack/Objects.png" width="256" height="256"/>
 </tileset>
 <tileset firstgid="653" name="Pipes-RustyWalls" tilewidth="32" tileheight="32" tilecount="88">
  <image source="../../../../../../alice-app/data/tiles/Skorpio's Pack/Pipes-RustyWalls.png" width="256" height="352"/>
 </tileset>
 <tileset firstgid="741" name="victoria" tilewidth="32" tileheight="32" tilecount="55">
  <image source="../../../../../../alice-app/data/tiles/lpc/victoria.png" width="352" height="160"/>
 </tileset>
 <tileset firstgid="796" name="dungeon" tilewidth="32" tileheight="32" tilecount="104">
  <image source="../../../../../../alice-app/data/tiles/lpc/dungeon.png" width="416" height="256"/>
 </tileset>
 <tileset firstgid="900" name="slime" tilewidth="32" tileheight="32" tilecount="12">
  <image source="../../../../../../alice-app/data/tiles/lpc-sprites/monsters/slime.png" width="96" height="128"/>
 </tileset>
 <tileset firstgid="912" name="princess" tilewidth="32" tileheight="32" tilecount="144">
  <image source="../../../../../../alice-app/data/tiles/lpc-sprites/people/princess.png" width="576" height="256"/>
 </tileset>
 <tileset firstgid="1056" name="Interior-Walls-Blue" tilewidth="32" tileheight="32" tilecount="100">
  <image source="../../../../../../../Documents/tiles/Skorpio's Pack/Interior-Walls-Blue.png" width="320" height="320"/>
 </tileset>
 <tileset firstgid="1156" name="waterfall" tilewidth="32" tileheight="32" tilecount="15">
  <image source="../../../../../../alice-app/data/tiles/lpc/waterfall.png" width="96" height="160"/>
 </tileset>
 <tileset firstgid="1171" name="water" tilewidth="32" tileheight="32" tilecount="18">
  <image source="../../../../../../alice-app/data/tiles/lpc/water.png" width="96" height="192"/>
 </tileset>
 <tileset firstgid="1189" name="mountains" tilewidth="32" tileheight="32" tilecount="108">
  <image source="../../../../../../alice-app/data/tiles/lpc/mountains.png" width="384" height="288"/>
 </tileset>
 <tileset firstgid="1297" name="dirt2" tilewidth="32" tileheight="32" tilecount="18">
  <image source="../../../../../../alice-app/data/tiles/lpc/dirt2.png" width="96" height="192"/>
 </tileset>
 <tileset firstgid="1315" name="airplane" tilewidth="32" tileheight="32" tilecount="88">
  <image source="../../../../../../alice-app/data/tiles/sithjester/airplane.png" width="256" height="352"/>
 </tileset>
 <tileset firstgid="1403" name="treetop" tilewidth="32" tileheight="32" tilecount="42">
  <image source="../../../../../../alice-app/data/tiles/lpc/treetop.png" width="192" height="224"/>
 </tileset>
 <tileset firstgid="1445" name="trunk" tilewidth="32" tileheight="32" tilecount="18">
  <image source="../../../../../../alice-app/data/tiles/lpc/trunk.png" width="192" height="96"/>
 </tileset>
 <tileset firstgid="1463" name="treetop_small" tilewidth="32" tileheight="32" tilecount="16">
  <image source="../../../../../../alice-app/data/tiles/lpc/treetop_small.png" width="128" height="128"/>
 </tileset>
 <tileset firstgid="1479" name="trunk_small" tilewidth="32" tileheight="32" tilecount="8">
  <image source="../../../../../../alice-app/data/tiles/lpc/trunk_small.png" width="128" height="64"/>
 </tileset>
 <tileset firstgid="1487" name="bridges" tilewidth="32" tileheight="32" tilecount="42">
  <image source="../../../../../../alice-app/data/tiles/lpc/bridges.png" width="192" height="224"/>
 </tileset>
 <tileset firstgid="1529" name="environment_assets" tilewidth="32" tileheight="32" tilecount="49">
  <image source="../../../../../../alice-app/data/tiles/turin/environment_assets.png" width="224" height="224"/>
 </tileset>
 <tileset firstgid="1578" name="vine" tilewidth="32" tileheight="32" tilecount="12">
  <image source="../../../../../../alice-app/data/tiles/vine.png" width="128" height="96"/>
 </tileset>
 <tileset firstgid="1590" name="man_eater_flower" tilewidth="32" tileheight="32" tilecount="6">
  <image source="../../../../../../alice-app/data/tiles/man_eater_flower.png" width="64" height="96"/>
 </tileset>
 <layer name="ground-cliff" width="27" height="23" opacity="0.26">
  <data encoding="base64" compression="zlib">
   eJxjYBgFo2AUjILBBbazIPCoXaN2jdo1ahct7NpJgZsI2XWYhWjcQKxakHsBqewsmg==
  </data>
 </layer>
 <layer name="ground-dirt" width="27" height="23">
  <properties>
   <property name="prop" value="hello"/>
  </properties>
  <data encoding="base64" compression="zlib">
   eJxjYBgFo2AUjALagO0s5GFy7aKHHnrZRUl4kGMXuYCQXmrFLzF2IdtHKaCGGcPVLlxxSWpcU1IukGIXpWUQsXaBxHcB8WEaxicA0xc7Ow==
  </data>
 </layer>
 <layer name="waterfall" width="27" height="23" visible="0">
  <data encoding="base64" compression="zlib">
   eJxjYBgFo2AUjIJRMApGwXABc1kQmB52ddPRLnr5y4yBIQGGaW1XO9A/HUDcSYS/AD8fCP4=
  </data>
 </layer>
 <layer name="fill" width="27" height="23">
  <data encoding="base64" compression="zlib">
   eJxrYmVgaBrFo3gYYRAYaDeMYkRcUApGql3UBPS0C90+WgJ6+AUGuGln9Khdo3YRDQCz0aqY
  </data>
 </layer>
 <layer name="grass" width="27" height="23">
  <data encoding="base64" compression="zlib">
   eJxjYBgFo2AUjALaAHYg5iARc5JpFxcQc5OIeci0SwbJDG0g1iGAdcmwSxKIpYBYAskuPSL0GZBhF8g/slBMC7tgfoH5B9keSuzCln5g5iNjatjFwICZftDtoadd6PKGQGxEABtTyV9MQMwHxcgAxAfGwQFmBtz5C90cBSyYGUkNofTFC7WXH49dTDjkkfXjMoMUgM8eagEAZ+0OcA==
  </data>
 </layer>
 <layer name="spring" width="27" height="23">
  <data encoding="base64" compression="zlib">
   eJxjYBgFo2AUjILhAbhxYFrZBQLaQKxDAOtSyS4Q0CeADahgF8w+QnZRwx5sdplCMTXswpYWaOWvgbaLVgDdLiYg5sOhloOB4QAzFe3koYJZxACQf/jpZBclAABp2Aum
  </data>
 </layer>
 <layer name="trunks" width="27" height="23">
  <data encoding="base64" compression="zlib">
   eJxjYBj8YBkrfj4xYCUZeqgFyHEvvQCy2yh1J6XxhK5+sMUZrrCiRvqkJUAPx6WsEDcuZyWOT4n5q4F4DRCvhaohxMflRmLs3wikNwHxZiL5MIDLDbQA2NwwEOmc1DgmBTQNsvCltl/xmUdtv+IzD1d6JhdQ27yRBrjpZA8APyI65g==
  </data>
 </layer>
 <layer name="trees" width="27" height="23">
  <data encoding="base64" compression="zlib">
   eJztVEkKAzEMyyWv7r5Pl0dXUEpdITvuTHobgUkUcGQ5SykfLGopS8Sqvvga4waxreULO/A94lA1z4BzjhhPiDPtcQEfENeqeQacc8N4RzwcnxHnHil4fVM+I676yvqcb7XZZ8RVX1mf89WZ2fq8ufKh9CMvXJ83Vz4UlBerb+vz5p4P9qu8sH6rX2pf79xa+tl+qfXMvbHI9ovXPf9j/oMWbP2Z9z4Ftv4x/2gP3Td6e4326+31lzs3Ff+4czNmZPEEKQ15OQ==
  </data>
 </layer>
 <layer name="objects-bridge" width="27" height="23">
  <data encoding="base64" compression="zlib">
   eJxjYBgFo2AUjIJRMApGwXAA51kZGC4g4YustLPrKtDsa1DzQfR1Gto1CiAAAJNDB5Y=
  </data>
 </layer>
 <imagelayer name="objects-vials" x="732" y="528">
  <image source="../../../../../../alice-app/data/images/test-tubes.png"/>
 </imagelayer>
 <imagelayer name="objects-toolbox" x="733" y="502">
  <image source="../../../../../../alice-app/data/images/toolbox.png"/>
 </imagelayer>
 <layer name="chara" width="27" height="23">
  <data encoding="base64" compression="zlib">
   eJztzbEJADAMBLEDu7b33zWBTJGvTgsIJEnSL6fgVuaahu3MpbwH4KoCJQ==
  </data>
 </layer>
 <layer name="vine-limb" width="27" height="23">
  <data encoding="base64" compression="zlib">
   eJxjYBgFo2AUjIJRMJyAPhvDAiBOMIDStLLHGGL+Aj2gHXo0tgvJzgMgTGt7RsEoGIwAADg1CHM=
  </data>
 </layer>
 <layer name="vine-head" width="27" height="23">
  <data encoding="base64" compression="zlib">
   eJxjYBgFo2AUjIJRMNzAdlYGhh2s9LFrN9CePXSyaxSMgpEMANLCAvs=
  </data>
 </layer>
</map>`
