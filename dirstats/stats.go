package dirstats

import (
	"os"
	"path/filepath"
	"strings"
)

// DirectoryStats 存储单个目录的统计信息
type DirectoryStats struct {
	DirCount  int
	FileCount int
	ExtCount  map[string]int
}

// Collect 读取指定目录路径，并收集其统计信息
func Collect(dirPath string) (DirectoryStats, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return DirectoryStats{}, err
	}

	stats := DirectoryStats{
		ExtCount: make(map[string]int),
	}

	for _, entry := range entries {
		if entry.IsDir() {
			stats.DirCount++
			continue
		}

		stats.FileCount++

		// --- FIX START: Improved extension detection ---
		name := entry.Name()
		ext := filepath.Ext(name)

		// Handle dotfiles like .bashrc, .dotfile
		// If the name starts with a dot and it's the only dot, treat as no extension.
		if strings.HasPrefix(name, ".") && strings.Count(name, ".") == 1 {
			ext = ""
		}
		// --- FIX END ---

		if ext == "" {
			ext = "no_extension"
		} else {
			ext = strings.ToLower(ext[1:]) // Remove dot and convert to lower case
		}
		stats.ExtCount[ext]++
	}

	return stats, nil
}
