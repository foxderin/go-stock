//go:build linux
// +build linux

package data

import (
	"fmt"
	"go-stock/backend/logger"
	"os/exec"
	"strings"
)

// AlertWindowsApi Linux平台的系统通知API
// @Author spark
// @Date 2025/8/20
// @Desc Linux平台系统通知实现
type AlertWindowsApi struct {
	AppID string
	// 窗口标题
	Title string
	// 窗口内容
	Content string
	// 窗口图标
	Icon string
}

func NewAlertWindowsApi(AppID string, Title string, Content string, Icon string) *AlertWindowsApi {
	return &AlertWindowsApi{
		AppID:   AppID,
		Title:   Title,
		Content: Content,
		Icon:    Icon,
	}
}

func (a AlertWindowsApi) SendNotification() bool {
	if GetSettingConfig().LocalPushEnable == false {
		logger.SugaredLogger.Error("本地推送未开启")
		return false
	}

	// 优先尝试使用 notify-send (libnotify)
	if err := a.sendWithNotifySend(); err == nil {
		return true
	}

	// 如果 notify-send 不可用，尝试使用 zenity
	if err := a.sendWithZenity(); err == nil {
		return true
	}

	// 如果都不可用，尝试使用 kdialog (KDE环境)
	if err := a.sendWithKDialog(); err == nil {
		return true
	}

	// 最后尝试使用 dunst 直接发送通知
	if err := a.sendWithDunst(); err == nil {
		return true
	}

	logger.SugaredLogger.Error("无法发送系统通知：未找到支持的通知系统")
	return false
}

// sendWithNotifySend 使用 notify-send 发送通知 (GNOME, Unity, 大多数Linux发行版)
func (a AlertWindowsApi) sendWithNotifySend() error {
	// 检查是否安装了 notify-send
	if _, err := exec.LookPath("notify-send"); err != nil {
		return fmt.Errorf("notify-send not found")
	}

	args := []string{
		"-t", "5000", // 显示5秒
		"-u", "normal", // 正常优先级
	}

	// 如果有图标路径，添加图标参数
	if a.Icon != "" {
		args = append(args, "-i", a.Icon)
	}

	// 如果有应用ID，添加应用名称
	if a.AppID != "" {
		args = append(args, "-a", a.AppID)
	}

	// 添加标题和内容
	args = append(args, a.Title, a.Content)

	cmd := exec.Command("notify-send", args...)
	err := cmd.Run()
	if err != nil {
		logger.SugaredLogger.Errorf("notify-send 执行失败: %v", err)
		return err
	}

	logger.SugaredLogger.Infof("已通过 notify-send 发送通知: %s", a.Title)
	return nil
}

// sendWithZenity 使用 zenity 发送通知
func (a AlertWindowsApi) sendWithZenity() error {
	if _, err := exec.LookPath("zenity"); err != nil {
		return fmt.Errorf("zenity not found")
	}

	text := fmt.Sprintf("%s\n\n%s", a.Title, a.Content)
	
	args := []string{
		"--info",
		"--text=" + text,
		"--timeout=5",
	}

	if a.Title != "" {
		args = append(args, "--title="+a.Title)
	}

	cmd := exec.Command("zenity", args...)
	err := cmd.Run()
	if err != nil {
		logger.SugaredLogger.Errorf("zenity 执行失败: %v", err)
		return err
	}

	logger.SugaredLogger.Infof("已通过 zenity 发送通知: %s", a.Title)
	return nil
}

// sendWithKDialog 使用 kdialog 发送通知 (KDE环境)
func (a AlertWindowsApi) sendWithKDialog() error {
	if _, err := exec.LookPath("kdialog"); err != nil {
		return fmt.Errorf("kdialog not found")
	}

	text := fmt.Sprintf("%s\n%s", a.Title, a.Content)
	
	args := []string{
		"--passivepopup",
		text,
		"5", // 显示5秒
	}

	if a.Title != "" {
		args = append(args, "--title", a.Title)
	}

	cmd := exec.Command("kdialog", args...)
	err := cmd.Run()
	if err != nil {
		logger.SugaredLogger.Errorf("kdialog 执行失败: %v", err)
		return err
	}

	logger.SugaredLogger.Infof("已通过 kdialog 发送通知: %s", a.Title)
	return nil
}

// sendWithDunst 尝试直接与 dunst 通知守护进程通信
func (a AlertWindowsApi) sendWithDunst() error {
	// 检查 dunst 是否在运行
	cmd := exec.Command("pgrep", "dunst")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("dunst not running")
	}

	// 使用 dunstify 如果可用 (dunst 的增强版本)
	if _, err := exec.LookPath("dunstify"); err == nil {
		return a.sendWithDunstify()
	}

	return fmt.Errorf("dunstify not available")
}

// sendWithDunstify 使用 dunstify 发送通知
func (a AlertWindowsApi) sendWithDunstify() error {
	args := []string{
		"-t", "5000", // 显示5秒
		"-u", "normal", // 正常优先级
	}

	if a.Icon != "" {
		args = append(args, "-i", a.Icon)
	}

	if a.AppID != "" {
		args = append(args, "-a", a.AppID)
	}

	args = append(args, a.Title, a.Content)

	cmd := exec.Command("dunstify", args...)
	err := cmd.Run()
	if err != nil {
		logger.SugaredLogger.Errorf("dunstify 执行失败: %v", err)
		return err
	}

	logger.SugaredLogger.Infof("已通过 dunstify 发送通知: %s", a.Title)
	return nil
}

// GetAvailableNotificationMethods 获取当前系统可用的通知方法
func GetAvailableNotificationMethods() []string {
	var methods []string
	
	if _, err := exec.LookPath("notify-send"); err == nil {
		methods = append(methods, "notify-send")
	}
	
	if _, err := exec.LookPath("zenity"); err == nil {
		methods = append(methods, "zenity")
	}
	
	if _, err := exec.LookPath("kdialog"); err == nil {
		methods = append(methods, "kdialog")
	}
	
	if _, err := exec.LookPath("dunstify"); err == nil {
		methods = append(methods, "dunstify")
	}
	
	return methods
}

// CheckNotificationSupport 检查系统是否支持通知
func CheckNotificationSupport() bool {
	methods := GetAvailableNotificationMethods()
	return len(methods) > 0
}

// GetDesktopEnvironment 获取当前桌面环境
func GetDesktopEnvironment() string {
	// 检查常见的桌面环境变量
	envVars := []string{
		"XDG_CURRENT_DESKTOP",
		"DESKTOP_SESSION", 
		"GDMSESSION",
	}
	
	for _, envVar := range envVars {
		if value := strings.ToLower(strings.TrimSpace(exec.Command("printenv", envVar).String())); value != "" {
			return value
		}
	}
	
	// 如果环境变量不可用，尝试检测正在运行的进程
	processes := []string{"gnome-session", "kde-session", "xfce4-session", "lxsession"}
	for _, process := range processes {
		if exec.Command("pgrep", process).Run() == nil {
			return strings.Replace(process, "-session", "", 1)
		}
	}
	
	return "unknown"
}
