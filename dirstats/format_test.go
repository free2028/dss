package dirstats

import (
	"strings"
	"testing"
)

func TestFormatResults(t *testing.T) {
	t.Run("formats single directory result correctly", func(t *testing.T) {
		results := []DirectoryResult{
			{
				Path: "/tmp/dir1",
				Stats: DirectoryStats{
					DirCount:  1,
					FileCount: 2,
					ExtCount:  map[string]int{"go": 1, "txt": 1},
				},
			},
		}

		actualOutput := FormatResults(results)
		// 使用反引号` `定义多行字符串，更清晰
		expectedOutput := `Directory: /tmp/dir1
dir:  1
.go:  1
.txt: 1
Total files: 2
`
		if actualOutput != expectedOutput {
			t.Errorf("Expected output:\n---\n%s\n---\nGot:\n---\n%s\n---", expectedOutput, actualOutput)
		}
	})

	t.Run("formats multiple directory results with aligned columns", func(t *testing.T) {
		results := []DirectoryResult{
			{
				Path: "projectA",
				Stats: DirectoryStats{
					DirCount:  2,
					FileCount: 3,
					ExtCount:  map[string]int{"js": 2, "no_extension": 1},
				},
			},
			{
				Path: "projectB",
				Stats: DirectoryStats{
					DirCount:  1,
					FileCount: 2,
					ExtCount:  map[string]int{"css": 1, "js": 1},
				},
			},
		}

		actualOutput := FormatResults(results)
		// 注意扩展名是按字母排序的（.css, .js, no_extension），并且列是对齐的
		expectedOutput := `Directory: projectA
dir:          2
.js:          2
no_extension: 1
Total files: 3

Directory: projectB
dir:          1
.css:         1
.js:          1
Total files: 2
`
		// 使用 strings.TrimSpace 忽略可能由编辑器引入的尾部空白差异
		if strings.TrimSpace(actualOutput) != strings.TrimSpace(expectedOutput) {
			t.Errorf("Expected output:\n---\n%s\n---\nGot:\n---\n%s\n---", expectedOutput, actualOutput)
		}
	})

	t.Run("handles empty results", func(t *testing.T) {
		results := []DirectoryResult{}
		actualOutput := FormatResults(results)
		if actualOutput != "" {
			t.Errorf("Expected empty string for empty results, got '%s'", actualOutput)
		}
	})
}
