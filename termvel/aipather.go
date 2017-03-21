package termvel

import (
	"fmt"
	"strings"

	"github.com/beefsack/go-astar"
)

// Kind* constants refer to PathTile kinds for input and output.
const (
	// KindPlain (.) is a plain PathTile with a movement cost of 1.
	KindPlain = iota
	// KindRiver (~) is a river PathTile with a movement cost of 2.
	KindRiver
	// KindMountain (M) is a mountain PathTile with a movement cost of 3.
	KindMountain
	// KindForest (^)
	KindForest
	// KindBlocker (X) is a PathTile which blocks movement.
	KindBlocker
	// KindFrom (F) is a PathTile which marks where the path should be calculated
	// from.
	KindFrom
	// KindTo (T) is a PathTile which marks the goal of the path.
	KindTo
	// KindPath (●) is a PathTile to represent where the path is in the output.
	KindPath
	KindOcean
	KindNeutral
)

// KindRunes map PathTile kinds to output runes.
var KindRunes = map[int]rune{
	KindPlain:    '.',
	KindRiver:    '~',
	KindMountain: 'M',
	KindForest:   '^',
	KindBlocker:  'X',
	KindFrom:     'F',
	KindTo:       'T',
	KindPath:     '●',
	KindOcean:    'O',
	KindNeutral:  ' ',
}

// RuneKinds map input runes to PathTile kinds.
var RuneKinds = map[rune]int{
	'.': KindPlain,
	'~': KindRiver,
	'M': KindMountain,
	'X': KindBlocker,
	'F': KindFrom,
	'T': KindTo,
	'^': KindForest,
	'O': KindOcean,
	' ': KindNeutral,
}

// KindCosts map PathTile kinds to movement costs.
var KindCosts = map[int]float64{
	KindPlain:    1.0,
	KindFrom:     1.0,
	KindTo:       1.0,
	KindRiver:    8.0,
	KindMountain: 3.0,
	KindForest:   3.0,
	KindOcean:    2.0,
	KindNeutral:  1.0,
}

// A PathTile is a PathTile in a grid which implements Pather.
type PathTile struct {
	// Kind is the kind of PathTile, potentially affecting movement.
	Kind int
	// X and Y are the coordinates of the PathTile.
	X, Y int
	// W is a reference to the WorldGrid that the PathTile is a part of.
	W WorldGrid
}

// PathNeighbors returns the neighbors of the PathTile, excluding blockers and
// PathTiles off the edge of the board.
func (t *PathTile) PathNeighbors() []astar.Pather {
	//log.Printf("%v,%v\n", t.X, t.Y)
	neighbors := []astar.Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {

		if n := t.W.PathTile(t.X+offset[0], t.Y+offset[1]); n != nil &&
			n.Kind != KindBlocker {
			neighbors = append(neighbors, n)
		}

	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring PathTile.
func (t *PathTile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*PathTile)
	return KindCosts[toT.Kind]
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *PathTile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*PathTile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

// WorldGrid is a two dimensional map of PathTiles.
type WorldGrid map[int]map[int]*PathTile

// PathTile gets the PathTile at the given coordinates in the WorldGrid.
func (w WorldGrid) PathTile(x, y int) *PathTile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

// SetPathTile sets a PathTile at the given coordinates in the WorldGrid.
func (w WorldGrid) SetPathTile(t *PathTile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*PathTile{}
	}
	w[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

// FirstOfKind gets the first PathTile on the board of a kind, used to get the from
// and to PathTiles as there should only be one of each.
func (w WorldGrid) FirstOfKind(kind int) *PathTile {
	for _, row := range w {
		for _, t := range row {
			if t.Kind == kind {
				return t
			}
		}
	}
	return nil
}

// From gets the from PathTile from the WorldGrid.
func (w WorldGrid) From() *PathTile {
	return w.FirstOfKind(KindFrom)
}

// To gets the to PathTile from the WorldGrid.
func (w WorldGrid) To() *PathTile {
	return w.FirstOfKind(KindTo)
}

// RenderPath renders a path on top of a WorldGrid.
func (w WorldGrid) RenderPath(path []astar.Pather) string {
	width := len(w)
	if width == 0 {
		return ""
	}
	height := len(w[0])
	pathLocs := map[string]bool{}
	for _, p := range path {
		pT := p.(*PathTile)
		pathLocs[fmt.Sprintf("%d,%d", pT.X, pT.Y)] = true
	}
	rows := make([]string, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			t := w.PathTile(x, y)
			r := ' '
			if pathLocs[fmt.Sprintf("%d,%d", x, y)] {
				r = KindRunes[KindPath]
			} else if t != nil {
				r = KindRunes[t.Kind]
			}
			rows[y] += string(r)
		}
	}
	return strings.Join(rows, "\n")
}

// ParseWorldGrid parses a textual representation of a WorldGrid into a WorldGrid map.
func ParseWorldGrid(input string) WorldGrid {
	w := WorldGrid{}
	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, raw := range row {
			kind, ok := RuneKinds[raw]
			if !ok {
				kind = KindBlocker
			}
			w.SetPathTile(&PathTile{
				Kind: kind,
			}, x, y)
		}
	}
	return w
}
