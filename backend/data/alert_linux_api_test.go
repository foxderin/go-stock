//go:build linux
// +build linux

package data

import (
	"testing"
)

func TestAlertLinuxApi_SendNotification(t *testing.T) {
	// 测试基本通知功能
	alert := NewAlertWindowsApi("go-stock-test", "测试标题", "这是一条测试通知消息", "")
	success := alert.SendNotification()
	
	if !success {
		t.Logf("通知发送失败，这可能是因为系统没有安装通知支持或者在无GUI环境中运行")
	} else {
		t.Logf("通知发送成功")
	}
}

func TestCheckNotificationSupport(t *testing.T) {
	support := CheckNotificationSupport()
	methods := GetAvailableNotificationMethods()
	
	t.Logf("通知支持: %v", support)
	t.Logf("可用方法: %v", methods)
	
	if support && len(methods) == 0 {
		t.Error("检测到通知支持但没有可用方法")
	}
	
	if !support && len(methods) > 0 {
		t.Error("检测到可用方法但通知支持为false")
	}
}

func TestGetDesktopEnvironment(t *testing.T) {
	desktop := GetDesktopEnvironment()
	t.Logf("检测到的桌面环境: %s", desktop)
	
	if desktop == "" {
		t.Error("无法检测桌面环境")
	}
}

func TestGetAvailableNotificationMethods(t *testing.T) {
	methods := GetAvailableNotificationMethods()
	t.Logf("可用的通知方法: %v", methods)
	
	// 至少应该有一种方法可用（在大多数Linux系统中）
	if len(methods) == 0 {
		t.Logf("警告: 没有找到可用的通知方法。请确保安装了以下之一: notify-send, zenity, kdialog, dunstify")
	}
}

// TestNotificationWithIcon 测试带图标的通知
func TestNotificationWithIcon(t *testing.T) {
	// 使用系统图标测试
	iconPaths := []string{
		"dialog-information",
		"/usr/share/pixmaps/stock.png", // 假设的应用图标
		"",                             // 无图标
	}
	
	for _, iconPath := range iconPaths {
		alert := NewAlertWindowsApi("go-stock-test", "图标测试", "测试不同图标设置", iconPath)
		success := alert.SendNotification()
		
		if iconPath == "" {
			t.Logf("无图标通知: %v", success)
		} else {
			t.Logf("图标 '%s' 通知: %v", iconPath, success)
		}
	}
}

// BenchmarkSendNotification 性能测试
func BenchmarkSendNotification(b *testing.B) {
	alert := NewAlertWindowsApi("go-stock-bench", "性能测试", "这是性能测试通知", "")
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alert.SendNotification()
	}
}

// TestNotificationLongText 测试长文本通知
func TestNotificationLongText(t *testing.T) {
	longText := "这是一条非常长的通知消息，用来测试通知系统如何处理长文本内容。" +
		"包含多行文本和特殊字符：！@#￥%……&*（）——+{}|："《》？[]\\;',./" +
		"中文字符测试：你好世界，这是股票监控应用的测试通知。" +
		"English text test: Hello World, this is a test notification from stock monitoring app."
	
	alert := NewAlertWindowsApi("go-stock-test", "长文本测试", longText, "")
	success := alert.SendNotification()
	
	t.Logf("长文本通知发送: %v", success)
}
