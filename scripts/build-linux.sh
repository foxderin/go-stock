#!/bin/bash

echo "开始构建 Linux 版本的 go-stock..."
cd "$(dirname "$0")/.."

# 确保依赖已安装
echo "检查构建依赖..."

# 检查 Go 版本
if ! command -v go &> /dev/null; then
    echo "错误: 未找到 Go 编译器"
    exit 1
fi

echo "Go 版本: $(go version)"

# 检查 Wails CLI
if ! command -v wails &> /dev/null; then
    echo "错误: 未找到 Wails CLI"
    echo "请安装 Wails CLI: go install github.com/wailsapp/wails/v2/cmd/wails@latest"
    exit 1
fi

echo "Wails 版本: $(wails version)"

# 检查 Linux 依赖
echo "检查 Linux 系统依赖..."

# 检查必要的系统包
REQUIRED_PACKAGES=("pkg-config" "libgtk-3-dev" "libwebkit2gtk-4.0-dev")
MISSING_PACKAGES=()

for package in "${REQUIRED_PACKAGES[@]}"; do
    if ! dpkg -l | grep -q "^ii  $package "; then
        MISSING_PACKAGES+=("$package")
    fi
done

if [ ${#MISSING_PACKAGES[@]} -ne 0 ]; then
    echo "警告: 以下开发包可能缺失: ${MISSING_PACKAGES[*]}"
    echo "在 Ubuntu/Debian 系统上安装: sudo apt-get install ${MISSING_PACKAGES[*]}"
    echo "在 Fedora 系统上安装: sudo dnf install gtk3-devel webkit2gtk3-devel pkg-config"
    echo "在 Arch Linux 上安装: sudo pacman -S gtk3 webkit2gtk pkg-config"
    echo ""
    echo "继续构建..."
fi

# 检查通知支持
echo "检查通知系统支持..."
NOTIFICATION_TOOLS=("notify-send" "zenity" "kdialog")
FOUND_TOOLS=()

for tool in "${NOTIFICATION_TOOLS[@]}"; do
    if command -v "$tool" &> /dev/null; then
        FOUND_TOOLS+=("$tool")
    fi
done

if [ ${#FOUND_TOOLS[@]} -eq 0 ]; then
    echo "警告: 未找到任何通知工具"
    echo "建议安装: sudo apt-get install libnotify-bin zenity"
else
    echo "找到通知工具: ${FOUND_TOOLS[*]}"
fi

# 检查浏览器支持
echo "检查浏览器支持..."
BROWSERS=("google-chrome" "google-chrome-stable" "chromium" "chromium-browser" "firefox")
FOUND_BROWSERS=()

for browser in "${BROWSERS[@]}"; do
    if command -v "$browser" &> /dev/null; then
        FOUND_BROWSERS+=("$browser")
    fi
done

if [ ${#FOUND_BROWSERS[@]} -eq 0 ]; then
    echo "警告: 未找到支持的浏览器"
    echo "建议安装: sudo apt-get install chromium-browser 或 firefox"
else
    echo "找到浏览器: ${FOUND_BROWSERS[*]}"
fi

echo ""
echo "开始构建应用..."

# 设置环境变量
export CGO_ENABLED=1
export GOOS=linux

# 构建应用
echo "执行 wails build..."
wails build --platform linux/amd64 --clean

BUILD_STATUS=$?

if [ $BUILD_STATUS -eq 0 ]; then
    echo ""
    echo "✅ 构建成功!"
    echo "可执行文件位置: ./build/bin/go-stock"
    
    # 检查生成的文件
    if [ -f "./build/bin/go-stock" ]; then
        echo "文件大小: $(du -h ./build/bin/go-stock | cut -f1)"
        echo "文件信息: $(file ./build/bin/go-stock)"
        
        # 创建桌面文件
        echo ""
        echo "创建桌面应用程序文件..."
        
        DESKTOP_FILE="$HOME/.local/share/applications/go-stock.desktop"
        APP_PATH="$(pwd)/build/bin/go-stock"
        ICON_PATH="$(pwd)/build/appicon.png"
        
        mkdir -p "$HOME/.local/share/applications"
        
        cat > "$DESKTOP_FILE" << EOF
[Desktop Entry]
Version=1.0
Type=Application
Name=go-stock
Comment=AI赋能股票分析✨
Exec=$APP_PATH
Icon=$ICON_PATH
Terminal=false
Categories=Office;Finance;
Keywords=stock;finance;analysis;AI;
StartupWMClass=go-stock
EOF
        
        if [ -f "$DESKTOP_FILE" ]; then
            chmod +x "$DESKTOP_FILE"
            echo "桌面文件已创建: $DESKTOP_FILE"
        fi
        
        # 使桌面文件生效
        if command -v update-desktop-database &> /dev/null; then
            update-desktop-database "$HOME/.local/share/applications"
            echo "桌面数据库已更新"
        fi
        
        echo ""
        echo "🚀 安装说明:"
        echo "1. 直接运行: ./build/bin/go-stock"
        echo "2. 或从应用程序菜单启动"
        echo "3. 首次运行前请确保安装了浏览器和通知支持"
        
    else
        echo "❌ 警告: 找不到构建的可执行文件"
    fi
else
    echo ""
    echo "❌ 构建失败 (退出码: $BUILD_STATUS)"
    echo "请检查错误信息并确保所有依赖都已正确安装"
    exit $BUILD_STATUS
fi

echo ""
echo "Linux 构建完成!"
