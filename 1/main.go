package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data, err := readData("input")
	if err != nil {
		log.Fatal(err)
	}

	result, err := firstPart(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d measurments are larger than previous measurment\n", result)

	result, err = secondPart(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d sums that are larger than the previous sum\n", result)
}

func readData(fileName string) ([]int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	data := make([]int, 0)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}

	return data, nil
}

func secondPart(data []int) (int, error) {
	counter := 0
	previous := sum(data[:3])
	for i := 1; i < len(data)+2; i++ {
		current := sum(data[i : i+3])
		if current > previous {
			counter++
		}
		previous = current
	}

	return counter, nil
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func firstPart(data []int) (int, error) {
	counter := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			counter++
		}
	}

	return counter, nil
}
