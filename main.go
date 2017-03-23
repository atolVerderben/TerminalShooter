package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/atolVerderben/TerminalShooter/termvel"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rawtext := ""

	g := termvel.NewGame()
	for {
		fmt.Println("")
		fmt.Println("Terminal Arena Shooter")
		fmt.Println("")
		fmt.Println("=======================================================================")
		fmt.Println("")
		fmt.Println("Controls:")
		fmt.Println("WASD: Move")
		fmt.Println("Arrow Keys: Aim")
		fmt.Println("Space: Shoot")
		fmt.Println("E: Detonate Bullet")
		fmt.Println("Q: Detonate All Bullets")
		fmt.Println("Mouse Click: Shoot in Direction")
		fmt.Println("Tab: Toggle Control Schemes (Click to shoot or Click to move)")
		fmt.Println("Backspace: Center Camera (only for large arenas)")
		fmt.Println("1: Switch to 'Classic' Controls")
		fmt.Println("2: Switch to 'New' Controls (Default Controls)")
		fmt.Println("Crtl+C: Quit")
		fmt.Println("")
		fmt.Println("=======================================================================")
		fmt.Println("")
		fmt.Println("Rules:")
		fmt.Println("A direct hit with a bullet is -2 HP")
		fmt.Println("A hit with an explosion is -1 HP")
		fmt.Println("All player begin with 6HP. Last standing player is the winner.")
		fmt.Println("")
		fmt.Println("=======================================================================")
		fmt.Println("")
		fmt.Println("Pick Arena Size: (S)mall or (L)arge")
		fmt.Println("Then Press ENTER to begin...")

		rawtext, _ = reader.ReadString('\n')

		switch []rune(strings.ToLower(rawtext))[0] {
		case 's':
			g.ArenaSize = "small"
			break
		case 'l':
			g.ArenaSize = "large"
			break
		default:
			g.ArenaSize = "small"
			break
		}
		g.Run()
		g.StopUpdate()
		fmt.Println("")
		fmt.Println("")
		fmt.Println("=======================================================================")
		fmt.Println("Thanks for playing. MUCH more polish and new features coming soon....")
		fmt.Println("=======================================================================")
		fmt.Println("Would you like to play again? (Y)es (N)o")

		rawtext, _ = reader.ReadString('\n')

		switch []rune(strings.ToLower(rawtext))[0] {
		case 'y':
			break
		case 'n':
			os.Exit(0)
		default:
			fmt.Println("You didn't say no, so I'm going to assume you want some more! ☺")
		}
	}
}
