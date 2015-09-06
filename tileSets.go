package tile

type TileSets interface {
	// TileSize in pixels of the sheets in this tileset
	TileSize() (int, int)
	// NumTileSets, returns sheet index to allow comparisions.
	MaxSheet() SheetIndex
	// Path of the passed tilset index.
	Path(SheetIndex) string
}
