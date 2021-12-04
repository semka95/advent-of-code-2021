package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	zero rune = '0'
	one  rune = '1'
)

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

	data := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result, err := partOne(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the power consumption of the submarine is %d\n", result)

	result, err = partTwo(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the life support rating of the submarine is %d\n", result)
}

func partOne(data []string) (int64, error) {
	ones := [12]int{}
	zeroes := [12]int{}
	for _, v := range data {
		for i, r := range v {
			switch r {
			case zero:
				zeroes[i]++
			case one:
				ones[i]++
			}
		}
	}

	gamma := ""
	epsilon := ""
	for i := 0; i < 12; i++ {
		if ones[i] > zeroes[i] {
			gamma += "1"
			epsilon += "0"
			continue
		}
		gamma += "0"
		epsilon += "1"
	}

	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return -1, err
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return -1, err
	}

	return g * e, nil
}

func partTwo(data []string) (int64, error) {
	oxygen, err := rec(0, data, true)
	if err != nil {
		return -1, err
	}
	co2, err := rec(0, data, false)
	if err != nil {
		return -1, err
	}
	return oxygen * co2, nil
}

func rec(iter int, data []string, more bool) (int64, error) {
	if len(data) == 1 {
		result, err := strconv.ParseInt(data[0], 2, 64)
		if err != nil {
			return -1, err
		}
		return result, nil
	}

	ones := make([]string, 0, len(data)/2)
	zeroes := make([]string, 0, len(data)/2)
	for _, v := range data {
		switch v[iter] {
		case 48:
			zeroes = append(zeroes, v)
		case 49:
			ones = append(ones, v)
		}
	}
	iter++

	if (more && len(ones) > len(zeroes)) || (!more && len(ones) < len(zeroes)) || (more && len(ones) == len(zeroes)) {
		return rec(iter, ones, more)
	}

	return rec(iter, zeroes, more)
}
