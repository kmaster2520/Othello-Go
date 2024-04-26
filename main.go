package main

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	gridboard         GameBoard
	tileCaptureValues [tilesPerRow][tilesPerRow]int
	countWhite        int
	countBlack        int
)

const (
	topY int32 = uiHeight
	topX int32 = 0
)

var (
	currentState  GameState
	currentPlayer TileValue

	mousePosition rl.Vector2
	mouseClicked  = false
)

func restartGame() {
	currentState = GameInProgress
	currentPlayer = TileBlack
	initializeBoard(&gridboard)
	setPlayerCounts(&gridboard)
	setNextValidMoves(&gridboard, currentPlayer)
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, gameName)
	rl.SetTargetFPS(30)
	restartGame()
}

func quit() {
	defer rl.CloseWindow()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)

	drawBoard(&gridboard)

	rl.DrawText("Current Player", 0, 0, 20, rl.White)
	var playerColor color.RGBA
	if currentPlayer == TileWhite {
		playerColor = playerWhiteColor
	} else {
		playerColor = playerBlackColor
	}
	rl.DrawCircle(tileSize/4*3, tileSize/4+24, tileSize/4, playerColor)

	rl.DrawCircle(tileSize*3, tileSize/4+24, tileSize/6, playerBlackColor)
	rl.DrawText(fmt.Sprint(countBlack), tileSize*4-24, tileSize/4+24, 20, rl.White)
	rl.DrawCircle(tileSize*5, tileSize/4+24, tileSize/6, playerWhiteColor)
	rl.DrawText(fmt.Sprint(countWhite), tileSize*6-24, tileSize/4+24, 20, rl.White)

	if currentState == GameOver {
		rl.DrawText("Game Over\n\nPress Space to Restart", screenWidth/4, screenHeight/2, 32, rl.SkyBlue)
	}

	rl.EndDrawing()
}

func input() {
	if rl.IsMouseButtonPressed(0) {
		mouseClicked = true
		mousePosition = rl.GetMousePosition()
	}
}

func update() {
	if currentState == GameInProgress {
		var (
			tileRow int = -1
			tileCol int = -1
		)
		if currentPlayer == TileBlack && mouseClicked {
			tileRow, tileCol = getTileCoordFromMousePosition(mousePosition)
		} else if currentPlayer == TileWhite {
			tileRow, tileCol = determineNextMoveForAIPlayer(&gridboard, currentPlayer)
			fmt.Printf("AI choose: %d %d (valid? %t)\n", tileRow+1, tileCol+1,
				getTileValueAt(&gridboard, tileRow, tileCol) == TileEmpty)
		}

		if isValidNextMove(tileRow, tileCol) {
			numCapturesForPlayerOnSpace(&gridboard, currentPlayer, tileRow, tileCol, true)
			setTileValueAt(&gridboard, tileRow, tileCol, currentPlayer)

			setPlayerCounts(&gridboard)

			if currentPlayer == TileBlack {
				currentPlayer = TileWhite
				//aiRow, aiCol := determineNextMoveForAIPlayer(&gridboard, currentPlayer)
				//fmt.Printf("%d %d\n", aiRow+1, aiCol+1)
			} else {
				currentPlayer = TileBlack
			}
			doesValidMoveExist := setNextValidMoves(&gridboard, currentPlayer)
			if !doesValidMoveExist {
				currentState = GameOver
			}

		}
	}

	if currentState == GameOver && rl.IsKeyDown(rl.KeySpace) {
		restartGame()
	}

	mouseClicked = false
}

func main() {
	defer quit()

	for !rl.WindowShouldClose() {

		input()
		update()

		render()

	}
}
