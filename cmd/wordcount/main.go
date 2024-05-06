package main

import (
	"fmt"
	"runtime"
	"os"
	"strconv"
	"strings"
	"word-count/pkg/fileparser"
	"word-count/pkg/mapreduce"
)

func main() {
	runtime.GOMAXPROCS(8)

	if len(os.Args) < 4 || len(os.Args) > 5 {
		fmt.Println("[Usage] wordcount <filename> <num splits> <query word> <optional: num repeats>")
		return
	}

	filename := os.Args[1]
	nSplit, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid number of splits")
		return
	}
	word := strings.ToLower(os.Args[3])
	var nRepeat int
	if len(os.Args) == 5 {
		nRepeat, err = strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println("Invalid number of repeats")
			return
		}
	}

	inputs, err := fileparser.SplitFile(filename, nSplit)
	if err != nil {
		fmt.Println(err)
		return
	}
	// inputs := make([][]string, 0)
	// inputs = append(inputs, fileparser.ParseString("Hello World Hello"))
	// inputs = append(inputs, fileparser.ParseString("World World Hello"))
	// word := "world"

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

	for i := 0; i < nRepeat - 1; i++ {
		mapreduce.MapReduce[[]string, string, int, int](inputs, mapper, reducer)
	}
	result := mapreduce.MapReduce[[]string, string, int, int](inputs, mapper, reducer)
	fmt.Println(result[word])
}