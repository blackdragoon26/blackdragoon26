package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	Rows         = 6
	Cols         = 7
	Empty        = "⚪"
	PlayerRed    = "🔴"
	PlayerYellow = "🟡"
	// We stick to the standard markers. 
	// The script logic change below will fix the position issue.
	StartMarker  = ""
	EndMarker    = ""
	ReadmeFile   = "README.md"
	RepoUser     = "blackdragoon26" 
	RepoName     = "blackdragoon26"
)

func main() {
	// 1. Get Move
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

	// 3. Extract Board Section
	// CRITICAL FIX: Use LastIndex to find the board at the BOTTOM of the file
	// This ignores any broken/ghost markers stuck at the top.
	startIndex := strings.LastIndex(content, StartMarker)
	
	if startIndex == -1 {
		fmt.Println("Markers not found. Please ensure is in README.")
		os.Exit(1)
	}

	// Find the EndMarker that comes AFTER the StartMarker we found
	rest OfContent := content[startIndex:]
	endIndexOffset := strings.Index(restOfContent, EndMarker)
	if endIndexOffset == -1 {
		panic("Start marker found but End marker missing")
	}
	endIndex := startIndex + endIndexOffset

	boardSection := content[startIndex+len(StartMarker) : endIndex]
	lines := strings.Split(strings.TrimSpace(boardSection), "\n")

	// 4. Parse Grid
	grid := make([][]string, 0)
	for _, line := range lines {
		cleanLine := strings.ReplaceAll(line, "|", " ")
		items := strings.Fields(cleanLine)
		if len(items) == Cols {
			isGameRow := true
			for _, item := range items {
				if item != Empty && item != PlayerRed && item != PlayerYellow {
					isGameRow = false
					break
				}
			}
			if isGameRow {
				grid = append(grid, items)
			}
		}
	}

	if len(grid) != Rows {
		grid = resetBoard()
	}

	// 5. Determine Player Turn
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

	// 6. Execute Move
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

	// 7. Check Win
	winner := ""
	if checkWin(grid, currentPlayer) {
		winner = currentPlayer
	}

	// 8. Reconstruct Output
	var sb strings.Builder
	sb.WriteString(StartMarker + "\n")
	
	if winner != "" {
		sb.WriteString(fmt.Sprintf("**GAME OVER! %s WINS! Resetting board...**\n", winner))
		grid = resetBoard()
	} else {
		nextPlayer := PlayerYellow
		if currentPlayer == PlayerYellow {
			nextPlayer = PlayerRed
		}
		sb.WriteString(fmt.Sprintf("Last move: %s in col %d. Next turn: %s\n", currentPlayer, col, nextPlayer))
	}

	sb.WriteString("| 1 | 2 | 3 | 4 | 5 | 6 | 7 |\n")
	sb.WriteString("|:---:|:---:|:---:|:---:|:---:|:---:|:---:|\n")

	for _, row := range grid {
		sb.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}

	friendlyMessage := "👋 **Click 'Submit new issue' below to play your move!**\n\nI am a bot powered by GitHub Actions. I will automatically update the board and close this issue in about 30 seconds."
	encodedBody := url.QueryEscape(friendlyMessage)
	
	sb.WriteString("|")
	for i := 0; i < Cols; i++ {
		link := fmt.Sprintf("https://github.com/%s/%s/issues/new?title=connect4%%7C%d&body=%s", RepoUser, RepoName, i, encodedBody)
		sb.WriteString(fmt.Sprintf(" [%s](%s) |", "⬇️", link))
	}
	sb.WriteString("\n")
	sb.WriteString(EndMarker)

	// 9. Write File
	newContent := content[:startIndex] + sb.String() + content[endIndex+len(EndMarker):]
	err = os.WriteFile(ReadmeFile, []byte(newContent), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Board updated successfully.")
}

func checkWin(grid [][]string, p string) bool {
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
