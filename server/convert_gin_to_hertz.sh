#!/bin/bash

# 这个脚本会批量将Gin代码转换为Hertz代码

# 查找所有需要转换的go文件（排除vendor和adapter目录）
FILES=$(find . -name "*.go" -not -path "./vendor/*" -not -path "./adapter/*" -not -path "./hertz_compat/*" -type f)

for file in $FILES; do
    echo "Processing: $file"
    
    # 1. 替换import语句
    # 将 "github.com/gin-gonic/gin" 替换为 hertz imports
    if grep -q '"github.com/gin-gonic/gin"' "$file"; then
        # 检查是否已经有context import
        if ! grep -q '"context"' "$file"; then
            # 添加context import并替换gin
            sed -i 's|"github.com/gin-gonic/gin"|"context"\n\t"github.com/cloudwego/hertz/pkg/app"\n\t"github.com/cloudwego/hertz/pkg/route"|' "$file"
        else
            #  已有context，只替换gin
            sed -i 's|"github.com/gin-gonic/gin"|"github.com/cloudwego/hertz/pkg/app"\n\t"github.com/cloudwego/hertz/pkg/route"|' "$file"
        fi
    fi
    
    # 2. 替换函数签名中的 *gin.Context
    # func XXX(c *gin.Context) -> func XXX(ctx context.Context, c *app.RequestContext)
    sed -i 's/func \([^(]*\)(c \*gin\.Context)/func \1(ctx context.Context, c *app.RequestContext)/' "$file"
    
    # 3. 替换类型中的gin引用
    sed -i 's/\*gin\.RouterGroup/\*route.RouterGroup/g' "$file"
    sed -i 's/gin\.RouterGroup/route.RouterGroup/g' "$file"
    sed -i 's/\*gin\.Engine/\*server.Hertz/g' "$file"
    sed -i 's/gin\.Engine/server.Hertz/g' "$file"
    sed -i 's/gin\.HandlerFunc/app.HandlerFunc/g' "$file"
    sed -i 's/gin\.IRoutes/route.IRoutes/g' "$file"
    
    # 4. 替换中间件和处理函数的闭包
    # return func(c *gin.Context) -> return func(ctx context.Context, c *app.RequestContext)
    sed -i 's/return func(c \*gin\.Context)/return func(ctx context.Context, c *app.RequestContext)/' "$file"
    sed -i 's/func(c \*gin\.Context)/func(ctx context.Context, c *app.RequestContext)/' "$file"
    
    # 5. 在需要的地方添加server import
    if grep -q 'server\.Hertz' "$file"; then
        if ! grep -q '"github.com/cloudwego/hertz/pkg/app/server"' "$file"; then
            sed -i 's|"github.com/cloudwego/hertz/pkg/route"|"github.com/cloudwego/hertz/pkg/route"\n\t"github.com/cloudwego/hertz/pkg/app/server"|' "$file"
        fi
    fi
    
done

echo "转换完成!"
echo "请注意：可能还需要手动调整以下内容："
echo "1. c.Next() 需要改为 c.Next(ctx)"
echo "2. 部分API可能有细微差异需要手动调整"
echo "3. 请运行 'go mod tidy' 更新依赖"
