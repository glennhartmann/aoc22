package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/glennhartmann/aoc22/src/common"
)

const width = 7

type shape [][]byte

// 0 is bottom
var shapes = []shape{
	shape{
		[]byte{'@', '@', '@', '@'},
	},
	shape{
		[]byte{'.', '@', '.'},
		[]byte{'@', '@', '@'},
		[]byte{'.', '@', '.'},
	},
	shape{
		// upside down, visually
		[]byte{'@', '@', '@'},
		[]byte{'.', '.', '@'},
		[]byte{'.', '.', '@'},
	},
	shape{
		[]byte{'@'},
		[]byte{'@'},
		[]byte{'@'},
		[]byte{'@'},
	},
	shape{
		[]byte{'@', '@'},
		[]byte{'@', '@'},
	},
}

type row [width]byte

func (r row) isEmpty() bool {
	return r == emptyLine
}

var emptyLine = row{'.', '.', '.', '.', '.', '.', '.'}

type direction byte

const (
	left  direction = '<'
	right direction = '>'
	down  direction = 'v'
)

func main() {
	r := bufio.NewReader(os.Stdin)
	jets, err := r.ReadString('\n')
	if err != nil {
		panic("unable to read")
	}
	jets = strings.TrimSpace(jets)

	screen := make([]row, 0, 5000) // 0 is bottom, grows upwards

	rocksStopped := 0
	needNextShape := true
	shapeCtr := -1
	jetCtr := -1

	var bottom, right int
	for {
		jetCtr++
		if needNextShape {
			shapeCtr++
			screen, bottom, right = nextShape(screen, shapes[shapeCtr%len(shapes)])
			//logScreen(screen)
		}

		jetDir := direction(jets[jetCtr%len(jets)])
		log.Printf("bottom: %d, right: %d, dir: %c", bottom, right, byte(jetDir))
		screen, _, _, right = move(screen, bottom, right, jetDir)
		//logScreen(screen)

		log.Printf("bottom: %d, right: %d, dir: v", bottom, right)
		screen, needNextShape, bottom, _ = move(screen, bottom, right, down)
		screen = trimScreenHeight(screen)
		//logScreen(screen)

		if needNextShape {
			rocksStopped++
		}
		if rocksStopped == 2022 {
			break
		}
	}
	logScreen(screen)
	log.Printf("height: %d", len(screen))
}

func trimScreenHeight(screen []row) []row {
	i := 0
	for i = range screen {
		if !screen[len(screen)-1-i].isEmpty() {
			break
		}
	}
	return screen[:len(screen)-i]
}

func nextShape(screen []row, shape shape) (rrow []row, btm, rght int) {
	screen = trimScreenHeight(screen)
	screen = append(screen, emptyLine, emptyLine, emptyLine)
	bottom := -1
	right := -1

	for i := range shape {
		r := []byte{'.', '.'}
		r = append(r, shape[i]...)
		lr := len(r)
		right = common.Max(lr-1, right)
		for j := 0; j < width-lr; j++ {
			r = append(r, '.')
		}
		screen = append(screen, *(*row)(r))
		if bottom < 0 {
			bottom = len(screen) - 1
		}
	}

	return screen, bottom, right
}

func move(screen []row, bottom, rght int, dir direction) ([]row, bool, int, int) {
	scr := append([]row(nil), screen...)
	for i := bottom; i < len(scr); i++ {
		switch dir {
		case left:
			for j := range scr[i] {
				if scr[i][j] != '@' {
					continue
				}
				if j == 0 || scr[i][j-1] == '#' {
					return screen, false, bottom, rght
				}
				scr[i][j-1] = '@'
				scr[i][j] = '.'
			}
		case right:
			for j := rght; j >= 0; j-- {
				if scr[i][j] != '@' {
					continue
				}
				if j == len(scr[i])-1 || scr[i][j+1] == '#' {
					return screen, false, bottom, rght
				}
				scr[i][j+1] = '@'
				scr[i][j] = '.'
			}
		case down:
			for j := range scr[i] {
				if scr[i][j] != '@' {
					continue
				}
				if i == 0 || scr[i-1][j] == '#' {
					return stopRock(screen, bottom, rght), true, bottom, rght
				}
				scr[i-1][j] = '@'
				scr[i][j] = '.'
			}
		}
	}

	switch dir {
	case left:
		rght--
	case right:
		rght++
	case down:
		bottom--
	}

	return scr, false, bottom, rght
}

func stopRock(screen []row, bottom, rght int) []row {
	for i := bottom; i < len(screen); i++ {
		for j := 0; j <= rght; j++ {
			if screen[i][j] == '@' {
				screen[i][j] = '#'
			}
		}
	}
	return screen
}

func logScreen(screen []row) {
	screen = trimScreenHeight(screen)
	for i := range screen {
		log.Printf("%s  row %d", string(screen[len(screen)-1-i][:]), i)
	}
	log.Print("")
}
