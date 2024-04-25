package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	boardColor       = rl.NewColor(63, 142, 106, 255)
	boardLineColor   = rl.NewColor(255, 255, 255, 255)
	playerBlackColor = rl.NewColor(32, 32, 32, 255)
	playerWhiteColor = rl.NewColor(228, 228, 228, 255)
	validSpaceColor  = rl.NewColor(196, 196, 32, 255)
)

func isValidPosition(row int, col int) bool {
	return row >= 0 && row < tilesPerRow && col >= 0 && col < tilesPerRow
}

func getTileValueAt(row int, col int) TileValue {
	return gridboard[row][col]
}

func setTileValueAt(row int, col int, tileValue TileValue) {
	gridboard[row][col] = tileValue
}

func setCaptureValue(row int, col int, numCaptures int) {
	tileCaptureValues[row][col] = numCaptures
}

func isValidNextMove(row int, col int) bool {
	return tileCaptureValues[row][col] > 0
}

func initializeBoard() {

	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {
			gridboard[r][c] = TileEmpty
		}
	}

	gridboard[tilesPerRow/2][tilesPerRow/2] = TileWhite
	gridboard[tilesPerRow/2-1][tilesPerRow/2-1] = TileWhite

	gridboard[tilesPerRow/2-1][tilesPerRow/2] = TileBlack
	gridboard[tilesPerRow/2][tilesPerRow/2-1] = TileBlack

}

/*
func copyBoard() [tilesPerRow][tilesPerRow]TileValue {
	var copiedBoard [tilesPerRow][tilesPerRow]TileValue
	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {
			copiedBoard[r][c] = getTileValueAt(r, c)
		}
	}
	return copiedBoard
}
*/
