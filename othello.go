package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func getTileCoordFromMousePosition(mousePosition rl.Vector2) (int, int) {

	var row int = (int(mousePosition.Y) - int(topY)) / tileSize
	var col int = (int(mousePosition.X) - int(topX)) / tileSize

	return row, col
}

func setPlayerCounts(board *GameBoard) {
	countBlack = 0
	countWhite = 0
	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {
			tileValue := getTileValueAt(board, r, c)
			if tileValue == TileBlack {
				countBlack++
			} else if tileValue == TileWhite {
				countWhite++
			}

		}
	}
}

func drawCapturedSpaces(toBeFlipped [18][2]int, count int) {
	for i := 0; i < count; i++ {
		r := toBeFlipped[i][0]
		c := toBeFlipped[i][1]
		x := int32(c*tileSize) + topX
		y := int32(r*tileSize) + topY
		rl.DrawCircle(x+tileSize/2, y+tileSize/2, tileSize/6, rl.Red)
	}
}

func drawBoard(board *GameBoard) {
	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {

			// draw tile
			x := int32(c*tileSize) + topX
			y := int32(r*tileSize) + topY
			rl.DrawRectangle(x, y, tileSize, tileSize, boardColor)
			rl.DrawRectangleLines(x, y, tileSize, tileSize, boardLineColor)

			// draw circle, if occupied
			tileValue := getTileValueAt(board, r, c)

			if tileValue == TileEmpty {
				if isValidNextMove(r, c) {
					rl.DrawCircleLines(x+tileSize/2, y+tileSize/2, tileSize/3, validSpaceColor)
				}
			} else {
				var playerColor color.RGBA
				if tileValue == TileWhite {
					playerColor = playerWhiteColor
				} else {
					playerColor = playerBlackColor
				}
				rl.DrawCircle(x+tileSize/2, y+tileSize/2, tileSize/3, playerColor)
			}

		}
	}
}

func setNextValidMoves(board *GameBoard, currentPlayer TileValue) bool {
	if currentPlayer == TileEmpty {
		return false
	}
	doesValidMoveExist := false

	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {
			if getTileValueAt(board, r, c) != TileEmpty {
				setCaptureValue(r, c, 0)
				continue
			}

			numCaptures, _ := numCapturesForPlayerOnSpace(board, currentPlayer, r, c)
			setCaptureValue(r, c, numCaptures)
			if numCaptures > 0 {
				doesValidMoveExist = true
			}
		}
	}

	return doesValidMoveExist
}

func numCapturesForPlayerOnSpace(board *GameBoard, player TileValue, row int, col int) (int, [18][2]int) {
	var toBeFlipped [18][2]int
	if row < 0 || row >= tilesPerRow || col < 0 || col >= tilesPerRow || player == TileEmpty {
		return 0, toBeFlipped
	}
	if getTileValueAt(board, row, col) != TileEmpty {
		return 0, toBeFlipped
	}

	var total int = 0
	var (
		nextRow int
		nextCol int
	)

	var opponent TileValue = getOpponent(player)

	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1},
		{0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		dr, dc := dir[0], dir[1]
		nextRow, nextCol = row+dr, col+dc
		var counted int = 0
		for isValidPosition(nextRow, nextCol) {
			tileValue := getTileValueAt(board, nextRow, nextCol)
			if tileValue == opponent {
				toBeFlipped[total+counted][0] = nextRow
				toBeFlipped[total+counted][1] = nextCol
				counted += 1
			} else if tileValue == player {
				total += counted
				break
			} else {
				break
			}
			nextRow += dr
			nextCol += dc
		}
	}

	return total, toBeFlipped

}
