#!/bin/sh

BUILD_DIR="./build"

# Build for linux systems
echo "YttriumDB | Build for linux"
GOOS="linux" go build -o $BUILD_DIR/yttriumDB-linux
echo "YttriumDB | Build for linux finished."

# Build for windows systems
echo "YttriumDB | Build for windows"
GOOS="windows" go build -o $BUILD_DIR/yttriumDB-windows.exe
echo "YttriumDB | Build for windows finished."

# Build for macOS systems
echo "YttriumDB | Build for macOS"
GOOS="darwin" go build -o $BUILD_DIR/yttriumDB-macOS
echo "YttriumDB | Build for macOS finished."