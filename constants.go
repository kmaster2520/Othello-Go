package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	tileSize     = 80
	tilesPerRow  = 8
	uiHeight     = 100
	screenWidth  = tilesPerRow * tileSize
	screenHeight = screenWidth + uiHeight
	gameName     = "Othello"
)

var (
	bkgColor = rl.NewColor(0, 0, 0, 255)
)

type TilePosition struct {
	Row int
	Col int
}

type TileValue int8

const (
	TileEmpty TileValue = iota
	TileBlack
	TileWhite
)

type GameBoard [tilesPerRow][tilesPerRow]TileValue

type GameState int8

const (
	GameInitial GameState = iota
	GameInProgress
	GameAnimateCapture
	GameOver
)
