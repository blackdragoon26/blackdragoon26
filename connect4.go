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
	RepoUser     = "blackdragoon26" // Change this if your user/repo changes
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
	startIndex := strings.Index(content, StartMarker)
	endIndex := strings.Index(content, EndMarker)
	if startIndex == -1 || endIndex == -1 {
		panic("Board markers not found in README")
	}

	boardSection := content[startIndex+len(StartMarker) : endIndex]
	lines := strings.Split(strings.TrimSpace(boardSection), "\n")

	// 4. Parse Grid from Markdown Table
	grid := make([][]string, 0)
	
	// We scan lines to find the actual rows with emojis
	for _, line := range lines {
		// Clean the line of pipes
		cleanLine := strings.ReplaceAll(line, "|", " ")
		items := strings.Fields(cleanLine)

		// Heuristic: A valid game row has exactly 7 items and contains our game pieces
		// We skip headers (1 2 3), separators (---), and button rows (⬇️)
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

	// Safety check if grid read failed
	if len(grid) != Rows {
		fmt.Println("Error reading grid. Found rows:", len(grid))
		// Fallback: Create empty board if parsing fails entirely to prevent crash
		if len(grid) == 0 {
			grid = resetBoard()
		} else {
			os.Exit(1)
		}
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
		os.Exit(0) // Exit success so action doesn't fail, but nothing changes
	}

	// 7. Check Win
	winner := ""
	if checkWin(grid, currentPlayer) {
		winner = currentPlayer
	}

	// 8. Reconstruct Output (The Pretty Table)
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

	// Table Header
	sb.WriteString("| 1 | 2 | 3 | 4 | 5 | 6 | 7 |\n")
	sb.WriteString("|:---:|:---:|:---:|:---:|:---:|:---:|:---:|\n") // Center alignment

	// Board Rows
	for _, row := range grid {
		sb.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}

	// Buttons Row
	sb.WriteString("|")
	for i := 0; i < Cols; i++ {
		link := fmt.Sprintf("https://github.com/%s/%s/issues/new?title=connect4%%7C%d&body=Just+push+submit", RepoUser, RepoName, i)
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
