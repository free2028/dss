package dirstats

import (
	"os"
	"reflect"
	"testing"
)

// 在运行测试前，请确保已经执行了 ./setup_test.sh
func TestCollect(t *testing.T) {
	testDir := "../test_data" // 相对测试文件位置

	// 检查测试目录是否存在
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Fatalf("Test directory '%s' not found. Please run setup_test.sh first.", testDir)
	}

	t.Run("collects stats from test directory", func(t *testing.T) {
		stats, err := Collect(testDir)
		if err != nil {
			t.Fatalf("Collect failed: %v", err)
		}

		// 验证目录数
		expectedDirCount := 2 // subdir1, emptydir
		if stats.DirCount != expectedDirCount {
			t.Errorf("Expected DirCount %d, got %d", expectedDirCount, stats.DirCount)
		}

		// 验证文件数
		expectedFileCount := 6 // file1.txt, file2.txt, document.PDF, archive.tar.gz, no_extension_file, .dotfile
		if stats.FileCount != expectedFileCount {
			t.Errorf("Expected FileCount %d, got %d", expectedFileCount, stats.FileCount)
		}

		// 验证扩展名统计
		expectedExtCount := map[string]int{
			"txt":          2,
			"pdf":          1,
			"gz":           1, // filepath.Ext("archive.tar.gz") is ".gz"
			"no_extension": 2, // no_extension_file and .dotfile
		}
		if !reflect.DeepEqual(stats.ExtCount, expectedExtCount) {
			t.Errorf("Expected ExtCount %v, got %v", expectedExtCount, stats.ExtCount)
		}
	})

	t.Run("handles non-existent directory", func(t *testing.T) {
		_, err := Collect("non_existent_dir_12345")
		if err == nil {
			t.Error("Expected an error for non-existent directory, but got nil")
		}
	})
}
