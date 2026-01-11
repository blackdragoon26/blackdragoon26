package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

const (
	Rows         = 6
	Cols         = 7
	Empty        = "⚪"
	PlayerRed    = "🔴"
	PlayerYellow = "🟡"
	StartMarker  = ""
	EndMarker    = ""
	ReadmeFile   = "README.md"
)

func main() {
	// 1. Get Move from Environment Variable
	title := os.Getenv("ISSUE_TITLE")
	if !strings.HasPrefix(title, "connect4|") {
		fmt.Println("Not a game issue. Exiting.")
		return
	}

	colStr := strings.TrimPrefix(title, "connect4|")
	col, err := strconv.Atoi(colStr)
	if err != nil || col < 0 || col >= Cols {
		fmt.Println("Invalid column. Exiting.")
		os.Exit(1)
	}

	// 2. Read README
	contentBytes, err := os.ReadFile(ReadmeFile)
	if err != nil {
		panic(err)
	}
	content := string(contentBytes)

	// 3. Extract Board
	startIndex := strings.Index(content, StartMarker)
	endIndex := strings.Index(content, EndMarker)
	if startIndex == -1 || endIndex == -1 {
		panic("Board markers not found in README")
	}

	boardSection := content[startIndex+len(StartMarker) : endIndex]
	boardLines := strings.Split(strings.TrimSpace(boardSection), "\n")
	
	// Skip the first line (status message)
	if len(boardLines) > 0 && strings.Contains(boardLines[0], "Last move:") {
		boardLines = boardLines[1:]
	}

	// Parse Grid
	grid := make([][]string, Rows)
	for r := 0; r < Rows; r++ {
		// Split by space to handle emojis correctly
		rowItems := strings.Fields(boardLines[r])
		// Ensure row has correct width (padding if needed)
		if len(rowItems) < Cols {
			for i := len(rowItems); i < Cols; i++ {
				rowItems = append(rowItems, Empty)
			}
		}
		grid[r] = rowItems
	}

	// 4. Determine Player Turn
	redCount, yellowCount := 0, 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == PlayerRed {
				redCount++
			} else if cell == PlayerYellow {
				yellowCount++
			}
		}
	}

	currentPlayer := PlayerRed
	if redCount > yellowCount {
		currentPlayer = PlayerYellow
	}

	// 5. Execute Move (Drop Piece)
	placed := false
	for r := Rows - 1; r >= 0; r-- {
		if grid[r][col] == Empty {
			grid[r][col] = currentPlayer
			placed = true
			break
		}
	}

	if !placed {
		fmt.Println("Column full. No move made.")
		os.Exit(0)
	}

	// 6. Check Win (Simplified Logic)
	winner := ""
	if checkWin(grid, currentPlayer) {
		winner = currentPlayer
	}

	// 7. Reconstruct Board String
	var newBoardStr strings.Builder
	newBoardStr.WriteString(StartMarker + "\n")
	
	if winner != "" {
		newBoardStr.WriteString(fmt.Sprintf("**GAME OVER! %s WINS! Resetting board...**\n", winner))
		// Reset Grid Logic could go here, but let's just show the win for now
		// To auto-reset, just overwrite 'grid' with empty rows before printing
		grid = resetBoard()
	} else {
		nextPlayer := PlayerYellow
		if currentPlayer == PlayerYellow {
			nextPlayer = PlayerRed
		}
		newBoardStr.WriteString(fmt.Sprintf("Last move: %s in col %d. Next turn: %s\n", currentPlayer, col, nextPlayer))
	}

	for _, row := range grid {
		newBoardStr.WriteString(strings.Join(row, " ") + "\n")
	}
	newBoardStr.WriteString(EndMarker)

	// 8. Write Back to File
	newContent := content[:startIndex] + newBoardStr.String() + content[endIndex+len(EndMarker):]
	err = os.WriteFile(ReadmeFile, []byte(newContent), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Board updated successfully.")
}

func checkWin(grid [][]string, p string) bool {
	// Horizontal, Vertical, Diagonal Check
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if grid[r][c] == p {
				if c+3 < Cols && grid[r][c+1] == p && grid[r][c+2] == p && grid[r][c+3] == p { return true }
				if r+3 < Rows && grid[r+1][c] == p && grid[r+2][c] == p && grid[r+3][c] == p { return true }
				if r+3 < Rows && c+3 < Cols && grid[r+1][c+1] == p && grid[r+2][c+2] == p && grid[r+3][c+3] == p { return true }
				if r-3 >= 0 && c+3 < Cols && grid[r-1][c+1] == p && grid[r-2][c+2] == p && grid[r-3][c+3] == p { return true }
			}
		}
	}
	return false
}

func resetBoard() [][]string {
	g := make([][]string, Rows)
	for r := 0; r < Rows; r++ {
		row := make([]string, Cols)
		for c := 0; c < Cols; c++ {
			row[c] = Empty
		}
		g[r] = row
	}
	return g
}
