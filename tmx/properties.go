package tmx

//Properties wraps any number of custom properties. Can be used as a child of the map, tile (when part of a tileset), layer, objectgroup and object elements.
type Properties struct {
	Property []Property `xml:"property"`
}

//When the property spans contains newlines, the current versions of Tiled Java and Tiled Qt will write out the value as characters contained inside the property element rather than as the value attribute. However, it is at the moment not really possible to edit properties consisting of multiple lines with Tiled.
//It is possible that a future version of the TMX format will switch to always saving property values inside the element rather than as an attribute.
type Property struct {
	Name  string `xml:"name,attr""`
	Value string `xml:"value,attr""`
}
