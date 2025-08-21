//go:build linux
// +build linux

package data

import (
	"os"
	"testing"
)

func TestCheckChrome_Linux(t *testing.T) {
	path, ok := CheckChrome()
	if ok {
		t.Logf("找到 Chrome 浏览器: %s", path)
		
		// 验证路径是否真实存在
		if path == "" {
			t.Error("Chrome 路径为空")
		}
	} else {
		t.Log("未找到 Chrome 浏览器")
	}
}

func TestCheckBrowser_Linux(t *testing.T) {
	path, ok := CheckBrowser()
	if ok {
		t.Logf("找到浏览器: %s", path)
		
		// 验证路径是否真实存在
		if path == "" {
			t.Error("浏览器路径为空")
		}
	} else {
		t.Log("未找到任何支持的浏览器")
		t.Log("建议安装以下浏览器之一: google-chrome, chromium, firefox")
	}
}

func TestCheckFirefox_Linux(t *testing.T) {
	path := checkFirefox()
	if path != "" {
		t.Logf("找到 Firefox 浏览器: %s", path)
	} else {
		t.Log("未找到 Firefox 浏览器")
	}
}

func TestFindExecutable_Linux(t *testing.T) {
	// 测试查找常见的Linux命令
	testCases := []string{
		"ls",        // 应该存在
		"which",     // 应该存在  
		"nonexistent", // 不应该存在
	}
	
	for _, cmd := range testCases {
		path := findExecutable(cmd)
		if path != "" {
			t.Logf("找到可执行文件 '%s': %s", cmd, path)
		} else {
			t.Logf("未找到可执行文件: %s", cmd)
		}
	}
}

// TestBrowsersDetection 全面测试浏览器检测
func TestBrowsersDetection_Linux(t *testing.T) {
	browsers := []string{
		"google-chrome",
		"google-chrome-stable", 
		"chromium",
		"chromium-browser",
		"firefox",
		"firefox-esr",
		"brave-browser",
		"opera",
		"vivaldi",
		"microsoft-edge",
	}
	
	t.Log("检测系统中安装的浏览器:")
	foundAny := false
	
	for _, browser := range browsers {
		path := findExecutable(browser)
		if path != "" {
			t.Logf("✓ %s: %s", browser, path)
			foundAny = true
		} else {
			t.Logf("✗ %s: 未安装", browser)
		}
	}
	
	if !foundAny {
		t.Log("警告: 没有找到任何浏览器，爬虫功能可能无法正常工作")
		t.Log("建议安装: sudo apt install chromium-browser 或 sudo apt install firefox")
	}
}

// BenchmarkCheckBrowser 性能测试
func BenchmarkCheckBrowser_Linux(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckBrowser()
	}
}

func BenchmarkCheckChrome_Linux(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckChrome()
	}
}

// TestLinuxDistroSpecific 测试特定Linux发行版的路径
func TestLinuxDistroSpecific_Linux(t *testing.T) {
	// Ubuntu/Debian 特有路径
	ubuntuPaths := []string{
		"/usr/bin/google-chrome-stable",
		"/usr/bin/chromium-browser",
		"/usr/bin/firefox",
	}
	
	// Fedora/RHEL 特有路径  
	fedoraPaths := []string{
		"/usr/bin/google-chrome",
		"/usr/bin/chromium",
		"/usr/bin/firefox",
	}
	
	// Arch Linux 特有路径
	archPaths := []string{
		"/usr/bin/google-chrome-stable",
		"/usr/bin/chromium", 
		"/usr/bin/firefox",
	}
	
	allPaths := append(append(ubuntuPaths, fedoraPaths...), archPaths...)
	
	t.Log("检查各发行版特定的浏览器路径:")
	for _, path := range allPaths {
		if _, err := os.Stat(path); err == nil {
			t.Logf("✓ 找到: %s", path)
		}
	}
}
