#!/bin/bash

# 检查当前目录是否是Git仓库
if [ ! -d .git ]; then
    echo "错误：当前目录不是一个Git仓库！"
    exit 1
fi

# 添加所有更改到暂存区
git add .

# 获取用户输入提交信息
read -p "请输入提交信息: " commit_message

# 如果用户没有输入提交信息，则使用默认信息
if [ -z "$commit_message" ]; then
    commit_message="自动提交于 $(date '+%Y-%m-%d %H:%M:%S')"
fi

# 提交更改
git commit -m "$commit_message"

# 获取当前分支名称
current_branch=$(git symbolic-ref --short HEAD)

# 推送到远程仓库
git push origin "$current_branch"

# 检查推送是否成功
if [ $? -eq 0 ]; then
    echo "✅ 代码成功推送到GitHub分支: $current_branch"
else
    echo "❌ 推送失败，请检查错误信息"
    exit 1
fi