//go:build linux
// +build linux

package data

import (
	"os"
	"os/exec"
	"strings"
)

// CheckChrome 检查 Linux 是否安装了 Chrome 浏览器
func CheckChrome() (string, bool) {
	// 检查常见的 Chrome 安装位置
	locations := []string{
		// Chrome 常见安装路径
		"/usr/bin/google-chrome",
		"/usr/bin/google-chrome-stable",
		"/usr/bin/chromium",
		"/usr/bin/chromium-browser",
		"/opt/google/chrome/google-chrome",
		"/snap/bin/chromium",
		// Flatpak 安装路径
		"/var/lib/flatpak/exports/bin/com.google.Chrome",
		"/var/lib/flatpak/exports/bin/org.chromium.Chromium",
	}

	for _, location := range locations {
		if _, err := os.Stat(location); err == nil {
			return location, true
		}
	}

	// 尝试通过 which 命令查找
	browsers := []string{"google-chrome", "google-chrome-stable", "chromium", "chromium-browser"}
	for _, browser := range browsers {
		if path := findExecutable(browser); path != "" {
			return path, true
		}
	}

	return "", false
}

// CheckBrowser 检查 Linux 是否安装了浏览器，并返回安装路径
func CheckBrowser() (string, bool) {
	// 优先检查 Chrome
	if path, ok := CheckChrome(); ok {
		return path, ok
	}

	// 检查 Firefox
	if path := checkFirefox(); path != "" {
		return path, true
	}

	// 检查其他常见浏览器
	browsers := []string{
		"firefox",
		"firefox-esr", 
		"brave-browser",
		"opera",
		"vivaldi",
		"microsoft-edge",
		"microsoft-edge-stable",
	}

	for _, browser := range browsers {
		if path := findExecutable(browser); path != "" {
			return path, true
		}
	}

	return "", false
}

// checkFirefox 检查 Firefox 浏览器
func checkFirefox() string {
	locations := []string{
		"/usr/bin/firefox",
		"/usr/bin/firefox-esr",
		"/opt/firefox/firefox",
		"/snap/bin/firefox",
		"/var/lib/flatpak/exports/bin/org.mozilla.firefox",
	}

	for _, location := range locations {
		if _, err := os.Stat(location); err == nil {
			return location
		}
	}

	return findExecutable("firefox")
}

// findExecutable 使用 which 命令查找可执行文件
func findExecutable(name string) string {
	cmd := exec.Command("which", name)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	
	path := strings.TrimSpace(string(output))
	if path != "" {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	
	return ""
}
