package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Board represents bingo board
type Board struct {
	Values    [5][]string // board values, -1 if marked
	WinNumber int         // index of the number that was just called when the board won
}

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

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	winNumbers := strings.Split(scanner.Text(), ",")
	// skip '\n'
	scanner.Scan()

	boards := make([]Board, 0)
	for scanner.Scan() {
		board := Board{
			Values:    [5][]string{},
			WinNumber: -1,
		}
		for i := 0; i < 5; i++ {
			// single digit number starts with space, trim double space and leading space
			board.Values[i] = strings.Split(strings.TrimPrefix(strings.ReplaceAll(scanner.Text(), "  ", " "), " "), " ")
			scanner.Scan()
		}
		boards = append(boards, board)
	}

	for i := range boards {
		boards[i].markBoard(winNumbers)
	}

	// first part
	wn := boards[0].WinNumber
	winBoardInd := 0
	for i := 1; i < len(boards); i++ {
		if boards[i].WinNumber < wn {
			wn = boards[i].WinNumber
			winBoardInd = i
		}
	}

	unmarkedSum := 0
	winBoard := boards[winBoardInd]
	for i := range winBoard.Values {
		for _, v := range winBoard.Values[i] {
			if v != "-1" {
				d, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				unmarkedSum += d
			}
		}
	}

	winNumber, err := strconv.Atoi(winNumbers[winBoard.WinNumber])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("final score of first win board is %d\n", unmarkedSum*winNumber)

	// second part
	wn = boards[0].WinNumber
	winBoardInd = 0
	for i := 1; i < len(boards); i++ {
		if boards[i].WinNumber > wn {
			wn = boards[i].WinNumber
			winBoardInd = i
		}
	}

	unmarkedSum = 0
	winBoard = boards[winBoardInd]
	for i := range winBoard.Values {
		for _, v := range winBoard.Values[i] {
			if v != "-1" {
				d, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				unmarkedSum += d
			}
		}
	}

	winNumber, err = strconv.Atoi(winNumbers[winBoard.WinNumber])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("final score of latest win board is %d\n", unmarkedSum*winNumber)
}

func (b *Board) markBoard(winNumbers []string) {
	for w := range winNumbers {
		if w > 4 && b.isBingo() {
			break
		}

		for i := range b.Values {
			for j := range b.Values[i] {
				if b.Values[i][j] == winNumbers[w] {
					b.Values[i][j] = "-1"
				}
			}
		}
		b.WinNumber = w
	}
}

func (b *Board) isBingo() bool {
	for i := range b.Values {
		marked := 0
		for j := range b.Values[i] {
			if b.Values[i][j] == "-1" {
				marked++
			}
		}
		if marked == 5 {
			return true
		}
	}

	for i := range b.Values {
		marked := 0
		for j := range b.Values[i] {
			if b.Values[j][i] == "-1" {
				marked++
			}
		}
		if marked == 5 {
			return true
		}
	}

	return false
}
