package tile

// Global index for a tile in the group of all loaded tile sheets.
type TileId int

// Index for a tile sheet in a group of all loaded tile sheets
// ( FIX: change to path? )
type SheetIndex int

type CellIndex int

// Tile(Sheet) and (Sheet)Index
type TileRef struct {
	Sheet SheetIndex
	Cell  CellIndex
}

func (t TileId) Advance(i int) TileId {
	return TileId(int(t) + i)
}
