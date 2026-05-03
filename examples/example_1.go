package main

import (
	"fmt"

	"github.com/flamefires/go-fake-useragent"
)

func main() {
	// 使用默认选项创建 UserAgent 实例
	ua := fakeuseragent.New()

	// 获取随机 User-Agent
	fmt.Println("随机 User-Agent:")
	fmt.Println(ua.Random())
	fmt.Println()

	// 获取特定浏览器的 User-Agent
	fmt.Println("Chrome User-Agent:")
	fmt.Println(ua.Chrome())
	fmt.Println()

	fmt.Println("Firefox User-Agent:")
	fmt.Println(ua.Firefox())
	fmt.Println()

	fmt.Println("Safari User-Agent:")
	fmt.Println(ua.Safari())
	fmt.Println()

	fmt.Println("Opera User-Agent:")
	fmt.Println(ua.Opera())
	fmt.Println()

	fmt.Println("Edge User-Agent:")
	fmt.Println(ua.Edge())
	fmt.Println()

	// 获取特定操作系统的 User-Agent
	fmt.Println("Android User-Agent:")
	fmt.Println(ua.Android())
	fmt.Println()

	fmt.Println("iOS User-Agent:")
	fmt.Println(ua.IOS())
	fmt.Println()

	// 获取特定平台类型的 User-Agent
	fmt.Println("Mobile User-Agent:")
	fmt.Println(ua.Mobile())
	fmt.Println()

	fmt.Println("Desktop User-Agent:")
	fmt.Println(ua.Desktop())
	fmt.Println()

	fmt.Println("Tablet User-Agent:")
	fmt.Println(ua.Tablet())
	fmt.Println()

	// 获取完整的 User-Agent 数据（包含更多信息）
	fmt.Println("随机 User-Agent 完整数据:")
	data := ua.GetRandom()
	fmt.Printf("User-Agent: %s\n", data.UserAgent)
	fmt.Printf("Browser: %s\n", data.Browser)
	fmt.Printf("Browser Version: %s\n", data.BrowserVersion)
	fmt.Printf("OS: %s\n", data.OS)
	fmt.Printf("OS Version: %s\n", data.OSVersion)
	fmt.Printf("Platform: %s\n", data.Platform)
	fmt.Printf("Type: %s\n", data.Type)
	fmt.Println()

	// 使用自定义选项创建 UserAgent 实例
	opts := fakeuseragent.DefaultOptions()
	opts.Browsers = []string{"Chrome", "Firefox"}
	opts.Platforms = []string{"desktop"}
	opts.MinVersion = 120.0

	uaCustom := fakeuseragent.NewWithOptions(opts)
	fmt.Println("自定义选项（仅 Chrome/Firefox 桌面版，版本 >= 120.0）:")
	fmt.Println(uaCustom.Random())
	fmt.Println()

	// 使用 Set 方法动态修改选项
	ua.SetBrowsers([]string{"Safari"})
	ua.SetOS([]string{"Mac OS X"})
	ua.SetPlatforms([]string{"mobile"})

	fmt.Println("动态设置选项（仅 Safari，Mac OS X，移动端）:")
	fmt.Println(ua.Random())
}
