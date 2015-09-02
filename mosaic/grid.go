package mosaic

import "image"

type TileType int8
type RotType int8

// Grid contains tile data, based on image.Gray.
// unimplemented:
//  func ColorModel() color.Model
//  func At(x, y int) color.Color
//  func Set(x, y int, c color.Color)
//  func SubImage(r Rectangle) Image
type Grid struct {
	// Cells holds the image's pixels, as gray values. The pixel at
	// (x, y) starts at Cells[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
	//Cells []Cell
	Tile []TileType `json:",omitempty"`
	Rot  []RotType  `json:",omitempty"`
	// Stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect bounds in cells.
	Rect image.Rectangle
	// CellSize in pixels.
	CellSize image.Point
	// tint, etc.
}

type Cell struct {
	Tile TileType
	Rot  RotType
	// tint, etc.
}

func (p *Grid) Bounds() image.Rectangle {
	return p.Rect
}

func (p *Grid) CellAt(x, y int) Cell {
	if !(image.Point{x, y}.In(p.Rect)) {
		return Cell{}
	}
	i := p.CellsOffset(x, y)
	return Cell{p.Tile[i], p.Rot[i]}
}

// CellsOffset returns the index of the first element of Cells that corresponds to the pixel at (x, y).
func (p *Grid) CellsOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*1
}

func (p *Grid) SetCell(x, y int, c Cell) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.CellsOffset(x, y)
	p.Tile[i] = c.Tile
	p.Rot[i] = c.Rot
}

// SubImage returns an image representing the portion of the image p visible through r.
// The returned value shares pixels with the original image.
func (p *Grid) SubGrid(r image.Rectangle) *Grid {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are image.Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Cells[i:] expression below can panic.
	if r.Empty() {
		return &Grid{}
	}
	i := p.CellsOffset(r.Min.X, r.Min.Y)
	return &Grid{
		Tile:     p.Tile[i:],
		Rot:      p.Rot[i:],
		Stride:   p.Stride,
		Rect:     r,
		CellSize: p.CellSize,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Grid) Opaque() bool {
	return true
}

// NewGrid returns a Grid with the given tile bounds.
func NewGrid(cellSize image.Point, r image.Rectangle) *Grid {
	w, h := r.Dx(), r.Dy()
	tile := make([]TileType, 1*w*h)
	rot := make([]RotType, 1*w*h)
	return &Grid{tile, rot, 1 * w, r, cellSize}
}
