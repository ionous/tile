package tmx

type Object struct {
	Id     int     `xml:"id,attr"`     // Unique ID of the object. Each object that is placed on a map gets a unique id. Even if an object was deleted, no object gets the same ID. Can not be changed in Tiled Qt. (since Tiled 0.11)
	X      float32 `xml:"x,attr"`      // x: The x coordinate of the object in pixels.
	Y      float32 `xml:"y,attr"`      // y: The y coordinate of the object in pixels.
	Width  float32 `xml:"width,attr"`  // width: The width of the object in pixels (defaults to 0).
	Height float32 `xml:"height,attr"` // height: The height of the object in pixels (defaults to 0).
}

// <objectgroup>
// name: The name of the object group.
// color: The color used to display the objects in this group.
// x: The x coordinate of the object group in tiles. Defaults to 0 and can no longer be changed in Tiled Qt.
// y: The y coordinate of the object group in tiles. Defaults to 0 and can no longer be changed in Tiled Qt.
// width: The width of the object group in tiles. Meaningless.
// height: The height of the object group in tiles. Meaningless.
// opacity: The opacity of the layer as a value from 0 to 1. Defaults to 1.
// visible: Whether the layer is shown (1) or hidden (0). Defaults to 1.
// draworder: Whether the objects are drawn according to the order of appearance ("index") or sorted by their y-coordinate ("top-down"). Defaults to "top-down".
// The object group is in fact a map layer, and is hence called "object layer" in Tiled Qt.

// Can contain: properties, object

// <object>
// name: The name of the object. An arbitrary string.
// type: The type of the object. An arbitrary string.
// rotation: The rotation of the object in degrees clockwise (defaults to 0). (since 0.10)
// gid: An reference to a tile (optional).
// visible: Whether the object is shown (1) or hidden (0). Defaults to 1. (since 0.9)
// While tile layers are very suitable for anything repetitive aligned to the tile grid, sometimes you want to annotate your map with other information, not necessarily aligned to the grid. Hence the objects have their coordinates and size in pixels, but you can still easily align that to the grid when you want to.

// You generally use objects to add custom information to your tile map, such as spawn points, warps, exits, etc.

// When the object has a gid set, then it is represented by the image of the tile with that global ID. Currently that means width and height are ignored for such objects. The image alignment currently depends on the map orientation. In orthogonal orientation it's aligned to the bottom-left while in isometric it's aligned to the bottom-center.

// Can contain: properties, ellipse (since 0.9), polygon, polyline, image

// <ellipse>
// Used to mark an object as an ellipse. The existing x, y, width and height attributes are used to determine the size of the ellipse.

// <polygon>
// points: A list of x,y coordinates in pixels.
// Each polygon object is made up of a space-delimited list of x,y coordinates. The origin for these coordinates is the location of the parent object. By default, the first point is created as 0,0 denoting that the point will originate exactly where the object is placed.

// <polyline>
// points: A list of x,y coordinates in pixels.
// A polyline follows the same placement definition as a polygon object.
