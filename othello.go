package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	gridboard      [tilesPerRow][tilesPerRow]TileValue
	validNextSteps [tilesPerRow][tilesPerRow]bool
	countWhite     int
	countBlack     int
)

const (
	topY int32 = uiHeight
	topX int32 = 0
)

func getTileCoordFromMousePosition(mousePosition rl.Vector2) (int, int) {

	var row int = (int(mousePosition.Y) - int(topY)) / tileSize
	var col int = (int(mousePosition.X) - int(topX)) / tileSize

	return row, col
}

func setPlayerCounts() {
	countBlack = 0
	countWhite = 0
	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {
			tileValue := getTileValueAt(r, c)
			if tileValue == TileBlack {
				countBlack++
			} else if tileValue == TileWhite {
				countWhite++
			}

		}
	}
}

func drawBoard() {
	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {

			// draw tile
			x := int32(c*tileSize) + topX
			y := int32(r*tileSize) + topY
			rl.DrawRectangle(x, y, tileSize, tileSize, boardColor)
			rl.DrawRectangleLines(x, y, tileSize, tileSize, boardLineColor)

			// draw circle, if occupied
			tileValue := getTileValueAt(r, c)

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

func setNextValidMoves(currentPlayer TileValue) bool {
	if currentPlayer == TileEmpty {
		return false
	}
	doesValidMoveExist := false

	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {
			if getTileValueAt(r, c) != TileEmpty {
				setValidNextMove(r, c, false)
				continue
			}

			numCaptures := numCapturesForPlayerOnSpace(currentPlayer, r, c, false)
			if numCaptures > 0 {
				setValidNextMove(r, c, true)
				doesValidMoveExist = true
			} else {
				setValidNextMove(r, c, false)
			}
		}
	}

	return doesValidMoveExist
}

func numCapturesForPlayerOnSpace(player TileValue, row int, col int, doFlip bool) int {
	if row < 0 || row >= tilesPerRow || col < 0 || col >= tilesPerRow || player == TileEmpty {
		return 0
	}

	var total int = 0
	var (
		nextRow int
		nextCol int
	)

	var opponent TileValue
	if player == TileBlack {
		opponent = TileWhite
	} else {
		opponent = TileBlack
	}

	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1},
		{0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	var toBeFlipped [8][2]int

	for _, dir := range directions {
		dr, dc := dir[0], dir[1]
		nextRow, nextCol = row+dr, col+dc
		var counted int = 0
		for isValidPosition(nextRow, nextCol) {
			tileValue := getTileValueAt(nextRow, nextCol)
			if tileValue == opponent {
				toBeFlipped[counted][0] = nextRow
				toBeFlipped[counted][1] = nextCol
				counted += 1
			} else if tileValue == player {
				total += counted
				if doFlip {
					for i := 0; i < counted; i++ {
						setTileValueAt(toBeFlipped[i][0], toBeFlipped[i][1], player)
					}
				}
				break
			} else {
				break
			}
			nextRow += dr
			nextCol += dc
		}
	}

	return total

}
