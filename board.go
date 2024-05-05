package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	boardColor       = rl.NewColor(63, 142, 106, 255)
	boardLineColor   = rl.NewColor(255, 255, 255, 255)
	playerBlackColor = rl.NewColor(32, 32, 32, 255)
	playerWhiteColor = rl.NewColor(228, 228, 228, 255)
	validSpaceColor  = rl.NewColor(196, 196, 32, 255)
)

func getOpponent(value TileValue) TileValue {
	if value == TileBlack {
		return TileWhite
	} else if value == TileWhite {
		return TileBlack
	}
	return TileEmpty
}

func isValidPosition(row int, col int) bool {
	return row >= 0 && row < tilesPerRow && col >= 0 && col < tilesPerRow
}

func getTileValueAt(board *GameBoard, row int, col int) TileValue {
	return board[row][col]
}

func setTileValueAt(board *GameBoard, row int, col int, tileValue TileValue) {
	board[row][col] = tileValue
}

func setTileValues(board *GameBoard, tileValue TileValue, spaces [18][2]int, count int) {
	for i := 0; i < count; i++ {
		setTileValueAt(board, spaces[i][0], spaces[i][1], tileValue)
	}
}

func setCaptureValue(row int, col int, numCaptures int) {
	tileCaptureValues[row][col] = numCaptures
}

func isValidNextMove(row int, col int) bool {
	return isValidPosition(row, col) && tileCaptureValues[row][col] > 0
}

func initializeBoard(board *GameBoard) {

	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {
			board[r][c] = TileEmpty
		}
	}

	board[tilesPerRow/2][tilesPerRow/2] = TileWhite
	board[tilesPerRow/2-1][tilesPerRow/2-1] = TileWhite

	board[tilesPerRow/2-1][tilesPerRow/2] = TileBlack
	board[tilesPerRow/2][tilesPerRow/2-1] = TileBlack
}

func copyBoard(board *GameBoard, newBoard *GameBoard) {
	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {
			newBoard[r][c] = board[r][c]
		}
	}
}
