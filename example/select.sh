#!/bin/bash

# 读取所选文本所在的文件
file=$1

# 获取所选文本的起始和结束行号和列号
startLine=$2
startColumn=$3
endLine=$4
endColumn=$5

# 提取所选文本内容
selectedText=$(awk -v sline="$startLine" -v scol="$startColumn" -v eline="$endLine" -v ecol="$endColumn" 'NR>=sline && NR<=eline {if(NR==sline && NR==eline) {print substr($0,scol,ecol-scol+1)} else if(NR==sline) {print substr($0,scol)} else if(NR==eline) {print substr($0,1,ecol)} else {print}}' "$file")

# 打印所选文本内容
echo "$selectedText"

