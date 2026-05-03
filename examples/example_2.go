package main

import (
	"fmt"
	"net/http"
	"time"

	fakeuseragent "github.com/flamefires/go-fake-useragent"
)

func main() {
	fmt.Println("=== 测试 go-fake-useragent 库 ===")
	fmt.Println()

	// 创建 UserAgent 实例
	ua := fakeuseragent.New()

	// 测试各种浏览器
	fmt.Println("1. 测试浏览器 User-Agent:")
	fmt.Printf("   Chrome:  %s\n", ua.Chrome())
	fmt.Printf("   Firefox: %s\n", ua.Firefox())
	fmt.Printf("   Safari:  %s\n", ua.Safari())
	fmt.Printf("   Edge:    %s\n", ua.Edge())
	fmt.Printf("   Opera:   %s\n", ua.Opera())
	fmt.Println()

	// 测试操作系统
	fmt.Println("2. 测试操作系统 User-Agent:")
	fmt.Printf("   Android: %s\n", ua.Android())
	fmt.Printf("   iOS:     %s\n", ua.IOS())
	fmt.Println()

	// 测试平台类型
	fmt.Println("3. 测试平台类型 User-Agent:")
	fmt.Printf("   Mobile:  %s\n", ua.Mobile())
	fmt.Printf("   Desktop: %s\n", ua.Desktop())
	fmt.Printf("   Tablet:  %s\n", ua.Tablet())
	fmt.Println()

	// 测试获取完整数据
	fmt.Println("4. 测试获取完整数据:")
	data := ua.GetRandom()
	fmt.Printf("   User-Agent: %s\n", data.UserAgent)
	fmt.Printf("   Browser:    %s\n", data.Browser)
	fmt.Printf("   Version:    %s\n", data.BrowserVersion)
	fmt.Printf("   OS:         %s\n", data.OS)
	fmt.Printf("   Platform:   %s\n", data.Platform)
	fmt.Printf("   Type:       %s\n", data.Type)
	fmt.Println()

	// 测试自定义选项
	fmt.Println("5. 测试自定义选项（仅 Chrome 桌面版）:")
	opts := fakeuseragent.DefaultOptions()
	opts.Browsers = []string{"Chrome"}
	opts.Platforms = []string{"desktop"}
	uaCustom := fakeuseragent.NewWithOptions(opts)
	fmt.Printf("   %s\n", uaCustom.Random())
	fmt.Println()

	// 测试实际 HTTP 请求
	fmt.Println("6. 测试实际 HTTP 请求:")
	testHTTPRequest(ua.Random())
	fmt.Println()

	fmt.Println("=== 测试完成 ===")
}

func testHTTPRequest(userAgent string) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://httpbin.org/user-agent", nil)
	if err != nil {
		fmt.Printf("   创建请求失败: %v\n", err)
		return
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("   请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("   请求成功! 状态码: %d\n", resp.StatusCode)
		fmt.Printf("   使用的 User-Agent: %s\n", userAgent)
	} else {
		fmt.Printf("   请求返回状态码: %d\n", resp.StatusCode)
	}
}
