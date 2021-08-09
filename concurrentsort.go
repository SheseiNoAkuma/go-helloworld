package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const numberOfRoutines = 4

func main() {

	slice := loadSlice()
	chunks := chunkSlice(slice)

	channel := make(chan []int, 4)
	for _, chunk := range chunks {
		go sortSlice(chunk, channel)
	}

	var orderedSlice []int
	for i := 0; i < len(chunks); i++ {
		orderedSlice = mergeKeepingOrder(orderedSlice, <-channel)
	}
	fmt.Printf("list ordered: %v\n", orderedSlice)
}

func mergeKeepingOrder(sliceOne []int, sliceTwo []int) []int {

	if sliceOne == nil {
		return sliceTwo
	}

	var merged []int
	indexOne := 0
	indexTwo := 0
	for {
		sliceOneFinished := indexOne > len(sliceOne)-1
		sliceTwoFinished := indexTwo > len(sliceTwo)-1
		if sliceOneFinished && sliceTwoFinished {
			break
		} else if sliceOneFinished {
			merged = append(merged, sliceTwo[indexTwo])
			indexTwo++
		} else if sliceTwoFinished {
			merged = append(merged, sliceOne[indexOne])
			indexOne++
		} else {
			vOne := sliceOne[indexOne]
			vTwo := sliceTwo[indexTwo]

			if vOne < vTwo {
				merged = append(merged, vOne)
				indexOne++
			} else {
				merged = append(merged, vTwo)
				indexTwo++
			}
		}
	}
	return merged
}

func sortSlice(slice []int, channel chan []int) {
	fmt.Printf("sorting the chunk %v\n", slice)
	BubbleSort(slice)
	channel <- slice
}

func BubbleSort(slice []int) {

	hasSwap := false
	for {
		for index, value := range slice {
			if index+1 < len(slice) && value > slice[index+1] {
				Swap(slice, index)
				hasSwap = true
			}
		}

		if hasSwap {
			hasSwap = false
		} else {
			break
		}
	}
}

func Swap(slice []int, index int) {
	tmpValue := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = tmpValue
}

func chunkSlice(slice []int) [][]int {

	chunkSize := len(slice) / numberOfRoutines
	if len(slice)%numberOfRoutines != 0 {
		chunkSize++
	}

	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func loadSlice() []int {

	var slice []int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("enter your list separated by space: ")
	scanner.Scan()
	line := scanner.Text()
	for _, x := range strings.Fields(line) {
		intVar, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, intVar)
	}
	return slice
}
