package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"BeesInTheTrap/internal/game"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func printWithDelay(message string, delay time.Duration) {
	fmt.Println(message)
	time.Sleep(delay)
}

func printTitle() {
	fmt.Printf("\n%s=== BEES IN THE TRAP ===%s\n", colorYellow, colorReset)
}

func printDivider() {
	fmt.Printf("%s----------------------------------------%s\n", colorBlue, colorReset)
}

func main() {
	// Clear screen
	fmt.Print("\033[H\033[2J")

	printTitle()
	printWithDelay(fmt.Sprintf("%sDestroy the hive before the bees sting you to death!%s", colorCyan, colorReset), 500*time.Millisecond)
	printDivider()
	printWithDelay(fmt.Sprintf("%sCommands:%s", colorPurple, colorReset), 300*time.Millisecond)
	printWithDelay(fmt.Sprintf("  %shit%s - Attack the beehive", colorGreen, colorReset), 300*time.Millisecond)
	printWithDelay(fmt.Sprintf("  %sauto%s - Toggle auto-play mode", colorGreen, colorReset), 300*time.Millisecond)
	printWithDelay(fmt.Sprintf("  %sstatus%s - Show game status", colorGreen, colorReset), 300*time.Millisecond)
	printWithDelay(fmt.Sprintf("  %sexit%s - Quit the game", colorGreen, colorReset), 300*time.Millisecond)
	printDivider()
	printWithDelay(fmt.Sprintf("\n%sGood luck!%s\n", colorYellow, colorReset), 500*time.Millisecond)

	g := game.NewGame()
	scanner := bufio.NewScanner(os.Stdin)

	for !g.IsGameOver() {
		if !g.AutoPlayMode {
			fmt.Printf("\n%sEnter command%s (%shit%s, %sauto%s, %sstatus%s, %sexit%s): ",
				colorWhite, colorReset,
				colorGreen, colorReset,
				colorGreen, colorReset,
				colorGreen, colorReset,
				colorGreen, colorReset)
			scanner.Scan()
			input := strings.ToLower(strings.TrimSpace(scanner.Text()))

			switch input {
			case "hit":
				message := g.PlayerTurn()
				printWithDelay(fmt.Sprintf("%s%s%s", colorCyan, message, colorReset), 300*time.Millisecond)

				if g.IsGameOver() {
					break
				}

				message = g.BeeTurn()
				printWithDelay(fmt.Sprintf("%s%s%s", colorRed, message, colorReset), 300*time.Millisecond)

			case "auto":
				g.ToggleAutoPlay()
				if g.AutoPlayMode {
					printWithDelay(fmt.Sprintf("%sAuto-play mode activated. Press Enter to stop.%s", colorPurple, colorReset), 300*time.Millisecond)
					go func() {
						scanner.Scan()
						g.ToggleAutoPlay()
						printWithDelay(fmt.Sprintf("%sAuto-play mode deactivated.%s", colorPurple, colorReset), 300*time.Millisecond)
					}()
				}

			case "status":
				printDivider()
				printWithDelay(fmt.Sprintf("%s%s%s", colorYellow, g.Status(), colorReset), 300*time.Millisecond)
				printDivider()

			case "exit":
				printWithDelay(fmt.Sprintf("%sExiting game...%s", colorRed, colorReset), 300*time.Millisecond)
				return

			default:
				printWithDelay(fmt.Sprintf("%sUnknown command. Try 'hit', 'auto', 'status', or 'exit'.%s", colorRed, colorReset), 300*time.Millisecond)
			}
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}

	printDivider()
	printWithDelay(fmt.Sprintf("\n%s=== GAME OVER ===%s", colorYellow, colorReset), 500*time.Millisecond)
	printWithDelay(fmt.Sprintf("%s%s%s", colorCyan, g.GetResult(), colorReset), 500*time.Millisecond)
	printDivider()
	printWithDelay(fmt.Sprintf("\n%s=== GAME SUMMARY ===%s", colorYellow, colorReset), 500*time.Millisecond)
	printWithDelay(fmt.Sprintf("%s%s%s", colorWhite, g.Summary(), colorReset), 500*time.Millisecond)
	printDivider()
}
