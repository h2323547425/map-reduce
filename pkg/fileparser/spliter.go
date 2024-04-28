package fileparser

import (
	"os"
	"strings"
)

func SplitFile(filePath string, numSplits int) ([][]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	
	lines := strings.Split(string(content), "\n")
	return split(lines, numSplits), nil
}

func split(lines []string, numSplits int) [][]string {
	splits := make([][]string, numSplits)
	size := len(lines)
	div := size / numSplits
	rem := size % numSplits
	cur := 0

	for i := 0; i < numSplits; i++ {
		step := div
		if rem > 0 {
			step++
			rem--
		}
		splits[i] = make([]string, 0)
		for _, line := range lines[cur : cur+step] {
			splits[i] = append(splits[i], ParseString(line)...)
		}

		cur += step
	}

	return splits
}