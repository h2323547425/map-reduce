package main

import (
	"fmt"
	"runtime"
	// "os"
	// "strconv"
	// "strings"
	"word-count/pkg/fileparser"
	"word-count/pkg/mapreduce"
)

func main() {
	// if len(os.Args) != 4 {
	// 	fmt.Println("[Usage] wordcount <filename> <num splits> <query word>")
	// 	return
	// }

	// filename := os.Args[1]
	// nSplit, err := strconv.Atoi(os.Args[2])
	// if err != nil {
	// 	fmt.Println("Invalid number of splits")
	// 	return
	// }
	// word := strings.ToLower(os.Args[3])

	// inputs, err := fileparser.SplitFile(filename, nSplit)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	runtime.GOMAXPROCS(2)

	inputs := make([][]string, 0)
	inputs = append(inputs, fileparser.ParseString("Hello World Hello"))
	inputs = append(inputs, fileparser.ParseString("World World Hello"))
	word := "world"

	mapper := func(input []string) map[string]int {
		result := make(map[string]int)
		for _, word := range input {
			result[word]++
		}
		return result
	}
	reducer := func(key string, values []int) int {
		result := 0
		for _, value := range values {
			result += value
		}
		return result
	}
	result := mapreduce.MapReduce[[]string, string, int, int](inputs, mapper, reducer)
	fmt.Println(result[word])
}