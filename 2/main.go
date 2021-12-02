package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	result, err := firstPart(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("multiplied final horizontal position by your final depth is %d\n", result)

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	result, err = secondPart(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("multiplied final horizontal position by your final depth is %d\n", result)
}

func firstPart(f io.Reader) (int, error) {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	depth := 0
	pos := 0

	for scanner.Scan() {
		switch scanner.Text() {
		case "forward":
			v, err := getValue(scanner)
			if err != nil {
				return -1, err
			}
			pos += v
		case "down":
			v, err := getValue(scanner)
			if err != nil {
				return -1, err
			}
			depth += v
		case "up":
			v, err := getValue(scanner)
			if err != nil {
				return -1, err
			}
			depth -= v
		}

	}

	return depth * pos, nil
}

func secondPart(f io.Reader) (int, error) {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	depth := 0
	pos := 0
	aim := 0

	for scanner.Scan() {
		switch scanner.Text() {
		case "forward":
			v, err := getValue(scanner)
			if err != nil {
				return -1, err
			}
			pos += v
			depth += aim * v
		case "down":
			v, err := getValue(scanner)
			if err != nil {
				return -1, err
			}
			aim += v
		case "up":
			v, err := getValue(scanner)
			if err != nil {
				return -1, err
			}
			aim -= v
		}

	}

	return depth * pos, nil
}

func getValue(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	v, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return -1, err
	}

	return v, nil
}
