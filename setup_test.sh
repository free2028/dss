#!/bin/bash

# 清理并创建测试根目录
TEST_DIR="test_data"
rm -rf "$TEST_DIR"
mkdir -p "$TEST_DIR"

# 在根目录创建文件和子目录
echo "Creating test structure in $TEST_DIR"
mkdir -p "$TEST_DIR/subdir1"
mkdir -p "$TEST_DIR/emptydir"

touch "$TEST_DIR/file1.txt"
touch "$TEST_DIR/file2.txt"
touch "$TEST_DIR/document.PDF"      # 测试大小写
touch "$TEST_DIR/archive.tar.gz"    # 测试多部分扩展名
touch "$TEST_DIR/no_extension_file" # 测试无扩展名
touch "$TEST_DIR/.dotfile"          # 测试隐藏文件（也无扩展名）

# 在子目录中创建文件
touch "$TEST_DIR/subdir1/script.py"

echo "Test setup complete."
