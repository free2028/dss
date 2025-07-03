package main

import (
	"dss/dirstats"
	"fmt"
	"os"
)

func main() {
	// 1. 获取目标目录
	targetDirs := []string{"."}
	if len(os.Args) > 1 {
		targetDirs = os.Args[1:]
	}

	allResults := make([]dirstats.DirectoryResult, 0, len(targetDirs))

	// 2. 循环处理，调用核心库
	for _, dir := range targetDirs {
		stats, err := dirstats.Collect(dir)
		if err != nil {
			fmt.Printf("Error reading directory '%s': %v\n", dir, err)
			continue
		}

		allResults = append(allResults, dirstats.DirectoryResult{
			Path:  dir,
			Stats: stats,
		})
	}

	// 3. 调用格式化函数并打印结果
	output := dirstats.FormatResults(allResults)
	fmt.Print(output)
}
