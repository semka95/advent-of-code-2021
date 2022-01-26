package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type VentLine [4]int
type Diagram [1000][1000]int

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	diagram := Diagram{}
	ventLines := make([]VentLine, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := strings.Split(strings.ReplaceAll(scanner.Text(), " -> ", ","), ",")
		vl := VentLine{}
		for i := range s {
			v, err := strconv.Atoi(s[i])
			if err != nil {
				log.Fatal(err)
			}
			vl[i] = v
		}
		ventLines = append(ventLines, vl)
	}

	for i := range ventLines {
		fill(&diagram, &ventLines[i])
	}

	sum := 0
	for i := range diagram {
		for j := range diagram {
			if diagram[i][j] > 1 {
				sum++
			}
		}
	}

	// fmt.Println(diagram[822][976], diagram[822][975])
	fmt.Println(sum)
}

func fill(d *Diagram, ventLine *VentLine) {
	var a, b int
	var isHorizontal bool
	if ventLine[0] == ventLine[2] && ventLine[1] == ventLine[3] {
		if ventLine[0] > ventLine[2] {
			ventLine[0], ventLine[1], ventLine[2], ventLine[3] = ventLine[2], ventLine[3], ventLine[0], ventLine[1]
		}
		for i := ventLine[0]; i <= ventLine[2]; i++ {
			d[i][i]++
		}
	}
	if ventLine[0] == ventLine[2] {
		if ventLine[1] > ventLine[3] {
			ventLine[1], ventLine[3] = ventLine[3], ventLine[1]
		}
		a = ventLine[1]
		b = ventLine[3]
		isHorizontal = false
	} else if ventLine[1] == ventLine[3] {
		if ventLine[0] > ventLine[2] {
			ventLine[0], ventLine[2] = ventLine[2], ventLine[0]
		}
		a = ventLine[0]
		b = ventLine[2]
		isHorizontal = true
	} else {
		switch {
		case ventLine[0] < ventLine[2] && ventLine[1] > ventLine[3]:
			for x, y := ventLine[0], ventLine[1]; x <= ventLine[2]; x, y = x+1, y-1 {
				d[x][y]++
			}
			return
		case ventLine[0] > ventLine[2] && ventLine[1] < ventLine[3]:
			for x, y := ventLine[0], ventLine[1]; x >= ventLine[2]; x, y = x-1, y+1 {
				d[x][y]++
			}
			return
		case ventLine[0] > ventLine[2] && ventLine[1] > ventLine[3]:
			for x, y := ventLine[0], ventLine[1]; x >= ventLine[2]; x, y = x-1, y-1 {
				d[x][y]++
			}
			return
		case ventLine[0] < ventLine[2] && ventLine[1] < ventLine[3]:
			for x, y := ventLine[0], ventLine[1]; x <= ventLine[2]; x, y = x+1, y+1 {
				d[x][y]++
			}
			return
		}
	}

	for ; a <= b; a++ {
		if isHorizontal {
			d[a][ventLine[1]]++
			continue
		}
		d[ventLine[0]][a]++
	}
}
