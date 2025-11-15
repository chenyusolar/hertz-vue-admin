#!/bin/bash

# 修复剩余的gin.H引用
echo "正在修复gin.H引用..."

# 使用perl来替换，因为它对特殊字符处理更好
find . -name "*.go" -not -path "./vendor/*" -type f -exec perl -pi -e 's/gin\.H\{/map[string]interface{}\{/g' {} \;

echo "gin.H替换完成!"

# 移除未使用的导入
echo "清理未使用的导入..."
which goimports > /dev/null 2>&1
if [ $? -eq 0 ]; then
    find . -name "*.go" -not -path "./vendor/*" -type f -exec goimports -w {} \;
    echo "goimports完成!"
else
    echo "goimports未安装，跳过..."
fi

echo "修复完成！请运行 'go build' 检查编译错误"
