package main

func determineNextMoveForAIPlayer(board *GameBoard, aiPlayer TileValue) (int, int) {
	nextRow, nextCol, _ := minimax(*board, aiPlayer, aiPlayer, true, 0, 4)
	return nextRow, nextCol
}

func getScore(board *GameBoard, targetValue TileValue) int {
	var total int = 0
	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {
			tileValue := getTileValueAt(board, r, c)
			if tileValue == targetValue {
				total++
			}
		}
	}
	return total
}

func minimax(board GameBoard, aiPlayer TileValue, currentPlayer TileValue, doMax bool, depth int, maxDepth int) (int, int, int) {

	if depth >= maxDepth {
		return 0, 0, getScore(&board, aiPlayer)
	}

	var (
		bestRow      int = -1
		bestCol      int = -1
		bestScore    int = 0
		nextBoard    GameBoard
		bestScoreSet bool = false
	)

	for r := 0; r < tilesPerRow; r++ {
		for c := 0; c < tilesPerRow; c++ {
			copyBoard(&board, &nextBoard)
			if numCapturesForPlayerOnSpace(&nextBoard, currentPlayer, r, c, true) > 0 {
				_, _, mmax := minimax(nextBoard, aiPlayer, getOpponent(currentPlayer), !doMax, depth+1, maxDepth)
				if (doMax && mmax > bestScore) || (!doMax && mmax < bestScore) || (!bestScoreSet) {
					bestRow = r
					bestCol = c
					bestScore = mmax
				}
			}
		}
	}

	return bestRow, bestCol, bestScore
}
