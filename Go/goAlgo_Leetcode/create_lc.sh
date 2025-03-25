#!/bin/bash

# 创建LeetCode题目目录结构的脚本

# 提示用户输入题号
read -p "请输入题号: " number

# 检查输入是否为空
while [[ -z "$number" ]]; do
    echo "题号不能为空！"
    read -p "请输入题号: " number
done

# 提示用户输入题名
read -p "请输入题名: " name

# 检查输入是否为空
while [[ -z "$name" ]]; do
    echo "题名不能为空！"
    read -p "请输入题名: " name
done

# 格式化目录名和文件名
dir_name="lc${number}_${name}"
file_name="lc${number}_${name}.go"

# 创建目录
mkdir -p "$dir_name"

# 创建Go文件并添加基本内容
cat > "$dir_name/$file_name" <<EOF
package main

/*
$number. $name

题目描述:
在这里添加题目描述

解题思路:
在这里添加解题思路
*/

func main() {
    // 测试代码
}
EOF

echo "已创建目录结构: $dir_name/$file_name"