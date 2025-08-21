//go:build linux
// +build linux

package main

import (
	"context"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"go-stock/backend/data"
	"go-stock/backend/db"
	"go-stock/backend/logger"
	"time"
)

// startup is called at application startup (Linux-specific implementation)
func (a *App) startup(ctx context.Context) {
	defer PanicHandler()
	runtime.EventsOn(ctx, "frontendError", func(optionalData ...interface{}) {
		logger.SugaredLogger.Errorf("Frontend error: %v\n", optionalData)
	})
	logger.SugaredLogger.Infof("Version:%s", Version)
	// Perform your setup here
	a.ctx = ctx

	// 监听设置更新事件
	runtime.EventsOn(ctx, "updateSettings", func(optionalData ...interface{}) {
		config := data.GetSettingConfig()
		logger.SugaredLogger.Infof("updateSettings config:%+v", config)
		if config.DarkTheme {
			runtime.WindowSetBackgroundColour(ctx, 27, 38, 54, 1)
			runtime.WindowSetDarkTheme(ctx)
		} else {
			runtime.WindowSetBackgroundColour(ctx, 255, 255, 255, 1)
			runtime.WindowSetLightTheme(ctx)
		}
		runtime.WindowReloadApp(ctx)
	})

	// Linux 系统通知初始化
	go func() {
		// 检查通知支持
		if data.CheckNotificationSupport() {
			methods := data.GetAvailableNotificationMethods()
			logger.SugaredLogger.Infof("Linux 通知支持已启用，可用方法: %v", methods)
			
			// 发送启动通知
			alert := data.NewAlertWindowsApi("go-stock", "go-stock", "应用程序已启动", "")
			alert.SendNotification()
		} else {
			logger.SugaredLogger.Warn("Linux 系统不支持通知功能，请安装 notify-send、zenity 或 kdialog")
		}
		
		// 检测桌面环境
		desktop := data.GetDesktopEnvironment() 
		logger.SugaredLogger.Infof("检测到桌面环境: %s", desktop)
	}()
	
	go setUpScreen(a)
	logger.SugaredLogger.Infof("Linux application startup Version:%s", Version)
}

func setUpScreen(a *App) {
	screens, _ := runtime.ScreenGetAll(a.ctx)
	if len(screens) == 0 {
		return
	}
	screen := screens[0]
	sw, sh := screen.Width, screen.Height

	// Linux 环境可能有顶部面板或底部面板，为它们留出空间
	topPanelHeight := 30   // 顶部面板高度（GNOME顶栏、KDE面板等）
	bottomPanelHeight := 40 // 底部面板高度（任务栏、Dock等）
	verticalMargin := topPanelHeight + bottomPanelHeight

	// 设置窗口为屏幕 80% 宽 × 可用高度 90%
	w := int(float64(sw) * 0.8)
	h := int(float64(sh-verticalMargin) * 0.9)

	runtime.WindowSetSize(a.ctx, w, h)
	runtime.WindowCenter(a.ctx)
}

// OnSecondInstanceLaunch 处理第二实例启动时的通知 (Linux-specific implementation)
func OnSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	alert := data.NewAlertWindowsApi("go-stock", "go-stock", "程序已经在运行了", "")
	success := alert.SendNotification()
	if !success {
		logger.SugaredLogger.Error("发送重复启动通知失败")
	}
	time.Sleep(time.Second * 3)
}

// MonitorStockPrices Linux版本的股票价格监控
func MonitorStockPrices(a *App) {
	dest := &[]data.FollowedStock{}
	db.Dao.Model(&data.FollowedStock{}).Find(dest)
	total := float64(0)

	// 股票信息处理逻辑
	stockInfos := GetStockInfos(*dest...)
	for _, stockInfo := range *stockInfos {
		if strutil.HasPrefixAny(stockInfo.Code, []string{"SZ", "SH", "sh", "sz"}) && (!isTradingTime(time.Now())) {
			continue
		}
		if strutil.HasPrefixAny(stockInfo.Code, []string{"hk", "HK"}) && (!IsHKTradingTime(time.Now())) {
			continue
		}
		if strutil.HasPrefixAny(stockInfo.Code, []string{"us", "US", "gb_"}) && (!IsUSTradingTime(time.Now())) {
			continue
		}

		total += stockInfo.ProfitAmountToday
		price, _ := convertor.ToFloat(stockInfo.Price)

		if stockInfo.PrePrice != price {
			go runtime.EventsEmit(a.ctx, "stock_price", stockInfo)
		}
	}

	// 计算总收益并更新状态
	if total != 0 {
		// Linux下使用窗口标题显示总收益（代替Windows的托盘）
		title := "go-stock " + time.Now().Format(time.DateTime) + fmt.Sprintf("  %.2f¥", total)
		runtime.WindowSetTitle(a.ctx, title)
		
		// 发送桌面通知（如果收益有显著变化）
		if total > 100 || total < -100 { // 收益超过100元或亏损超过100元时通知
			message := fmt.Sprintf("当前总收益: %.2f¥", total)
			alert := data.NewAlertWindowsApi("go-stock", "股票收益提醒", message, "")
			go alert.SendNotification()
		}
	}

	go runtime.EventsEmit(a.ctx, "realtime_profit", fmt.Sprintf("  %.2f", total))
}

// beforeClose Linux版本的关闭前确认对话框
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	defer PanicHandler()

	// 在 Linux 上使用 MessageDialog 显示确认窗口
	dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:         runtime.QuestionDialog,
		Title:        "go-stock",
		Message:      "确定关闭吗？",
		Buttons:      []string{"确定", "取消"},
		Icon:         icon,
		CancelButton: "取消",
	})

	if err != nil {
		logger.SugaredLogger.Errorf("dialog error:%s", err.Error())
		return false
	}

	logger.SugaredLogger.Debugf("dialog:%s", dialog)
	if dialog == "取消" {
		return true // 如果选择了取消，不关闭应用
	} else {
		// Linux 应用退出时执行清理工作
		if a.cron != nil {
			a.cron.Stop() // 停止定时任务
		}
		return false // 如果选择了确定，继续关闭应用
	}
}

// getScreenResolution Linux版本获取屏幕分辨率
func getScreenResolution() (int, int, int, int, error) {
	// Linux下的默认值，因为在启动时会动态获取
	// 这些值会在setUpScreen函数中被实际的屏幕尺寸覆盖
	width := 1920
	height := 1080
	minWidth := 1000
	minHeight := 600
	
	return width, height, minWidth, minHeight, nil
}

// getFrameless Linux版本是否使用无边框窗口
func getFrameless() bool {
	return false
}
