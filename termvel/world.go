package termvel

import (
	"bufio"
	"os"

	tl "github.com/JoelOtter/termloop"
)

//ReadAStarFile will be used to read in a file and convert to WorldGrid
//As of right now it just uses the built in testmap
func ReadAStarFile(filename string) *WorldGrid {

	/*lines, _ := readLines(filename)
	//content, err := ioutil.ReadFile(filename)
	//if err != nil {
	lineText := ""
	for _, line := range lines {
		lineText += line + "\n"
	}*/
	wg := ParseWorldGrid(testmap)
	/*log.Printf("%v", lineText)
	reader := bufio.NewReader(os.Stdin)
	rawtext, _ := reader.ReadString('\n')
	log.Printf("%v\n", rawtext)*/
	return &wg
	/*} else {
		log.Printf("%v", err)
		reader := bufio.NewReader(os.Stdin)
		rawtext, _ := reader.ReadString('\n')
		log.Printf("%v\n", rawtext)
	}*/
	//return nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//World represents both the drawn level and the a* world for pathfinding
type World struct {
	*tl.BaseLevel
	Grid          *WorldGrid
	Height, Width int
} //⚰️
