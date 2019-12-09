package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	defaultFilename = "lines.txt"
)

func parseFile(fileName string) ([]string, []string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	line1 := strings.Split(lines[0], ",")
	line2 := strings.Split(lines[1], ",")

	return line1, line2, err
}

func trimFirstChar(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func mapCoords(line []string) [][]int {
	coords := [][]int{{0, 0}}

	for _, move := range line {
		mtype := string(move[0])
		amount, err := strconv.Atoi(trimFirstChar(move))
		if err != nil {
			log.Fatal(err)
		}

		lastMove := make([]int, 2)
		copy(lastMove, coords[len(coords)-1])

		switch mtype {
		case "U":
			lastMove[1] = lastMove[1] + amount
			break
		case "D":
			lastMove[1] = lastMove[1] - amount
			break
		case "L":
			lastMove[0] = lastMove[0] + amount
			break
		case "R":
			lastMove[0] = lastMove[0] - amount
			break
		}

		coords = append(coords, lastMove)
	}

	return coords
}

func line(coord1 []int, coord2 []int) []int {
	var calcs []int
	a := (coord1[1] - coord2[1])
	b := (coord2[0] - coord1[0])
	c := ((coord1[0] * coord2[1]) - (coord2[0] * coord1[1]))

	calcs = append(calcs, a, b, -c)

	return calcs
}

func lineIntersection(line1 []int, line2 []int) []int {
	d := (line1[0] * line2[1]) - (line1[1] * line2[0])
	dx := (line1[2] * line2[1]) - (line1[1] * line2[2])
	dy := (line1[0] * line2[2]) - (line1[2] * line2[0])

	if d != 0 {
		return []int{dx / d, dy / d}
	} else {
		return nil
	}
}



func intersects(coords1 [][]int, coords2 [][]int) {
	var intersecting [][]int

	for i, coord1a := range coords1 {
		if i != 0 {
			line1 := line(coord1a, coords1[i-1])
			for j, coord2a := range coords2 {
				if j != 0 {
					line2 := line(coord2a, coords2[j-1])

					lineIntersect := lineIntersection(line1, line2)

					if lineIntersect != nil {
						intersecting = append(intersecting, lineIntersect)
					}
				}
			}
		}
	}

	fmt.Println(intersecting)
}

func main() {
	line1, line2, err := parseFile(defaultFilename)
	if err != nil {
		log.Fatal(err)
	}

	intersects(mapCoords(line1), mapCoords(line2))

}
