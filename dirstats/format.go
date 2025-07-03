package dirstats

import (
	"fmt"
	"sort"
	"strings"
)

// DirectoryResult 包含路径和其对应的统计信息
type DirectoryResult struct {
	Path  string
	Stats DirectoryStats
}

// FormatResults 将多个目录的统计结果格式化为最终的输出字符串
func FormatResults(results []DirectoryResult) string {
	if len(results) == 0 {
		return ""
	}

	// 1. 收集所有出现过的扩展名类型
	allExtensionsSet := make(map[string]bool)
	for _, result := range results {
		for ext := range result.Stats.ExtCount {
			allExtensionsSet[ext] = true
		}
	}

	// 2. 将扩展名转换为排序后的切片，确保输出顺序稳定
	sortedExts := make([]string, 0, len(allExtensionsSet))
	for ext := range allExtensionsSet {
		sortedExts = append(sortedExts, ext)
	}
	sort.Strings(sortedExts)

	// 3. 计算全局最大名称长度，以便统一对齐
	// 这个长度是基于所有将要被显示的名称来计算的
	maxNameLen := len("dir") // 'dir' is always a potential candidate

	for _, ext := range sortedExts {
		var nameLen int
		if ext == "no_extension" {
			nameLen = len("no_extension")
		} else {
			nameLen = len(ext) + 1 // +1 for the dot '.'
		}
		if nameLen > maxNameLen {
			maxNameLen = nameLen
		}
	}

	var sb strings.Builder
	// 4. 遍历每个结果，生成格式化输出
	for i, result := range results {
		if i > 0 {
			sb.WriteString("\n")
		}

		sb.WriteString(fmt.Sprintf("Directory: %s\n", result.Path))

		// 使用全局列宽打印 "dir" 数量
		sb.WriteString(fmt.Sprintf("%-*s %d\n", maxNameLen+1, "dir"+":", result.Stats.DirCount))

		// 严格按照全局排序的扩展名列表来打印，确保顺序和对齐一致
		for _, ext := range sortedExts {
			if count, ok := result.Stats.ExtCount[ext]; ok && count > 0 {
				displayName := ext
				if ext != "no_extension" {
					displayName = "." + ext
				}
				sb.WriteString(fmt.Sprintf("%-*s %d\n", maxNameLen+1, displayName+":", count))
			}
		}

		sb.WriteString(fmt.Sprintf("Total files: %d\n", result.Stats.FileCount))
	}

	return sb.String()
}
