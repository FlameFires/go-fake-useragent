# go-fake-useragent

一个 Go 语言实现的 User-Agent 伪装库，参考 [fake-useragent/fake-useragent](https://github.com/fake-useragent/fake-useragent) Python 库。

## 特性

- 内置真实浏览器 User-Agent 数据库
- 支持多种浏览器（Chrome, Firefox, Safari, Opera, Edge 等）
- 支持桌面、平板和移动设备平台
- 支持按操作系统、浏览器、平台、版本等条件过滤
- 线程安全
- 零外部依赖

## 安装

```bash
go get github.com/flamefires/go-fake-useragent
```

## 使用方法

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/flamefires/go-fake-useragent"
)

func main() {
    // 创建 UserAgent 实例
    ua := fakeuseragent.New()

    // 获取随机 User-Agent
    fmt.Println(ua.Random())

    // 获取特定浏览器的 User-Agent
    fmt.Println(ua.Chrome())
    fmt.Println(ua.Firefox())
    fmt.Println(ua.Safari())
    fmt.Println(ua.Opera())
    fmt.Println(ua.Edge())
}
```

### 获取完整数据

除了 User-Agent 字符串，还可以获取包含更多信息的完整数据：

```go
data := ua.GetRandom()
fmt.Printf("User-Agent: %s\n", data.UserAgent)
fmt.Printf("Browser: %s\n", data.Browser)
fmt.Printf("Browser Version: %s\n", data.BrowserVersion)
fmt.Printf("OS: %s\n", data.OS)
fmt.Printf("OS Version: %s\n", data.OSVersion)
fmt.Printf("Platform: %s\n", data.Platform)
fmt.Printf("Type: %s\n", data.Type)
```

### 自定义选项

```go
opts := fakeuseragent.DefaultOptions()
opts.Browsers = []string{"Chrome", "Firefox"}
opts.Platforms = []string{"desktop"}
opts.MinVersion = 120.0

ua := fakeuseragent.NewWithOptions(opts)
fmt.Println(ua.Random())
```

### 获取特定操作系统的 User-Agent

```go
ua := fakeuseragent.New()

// 获取 Android User-Agent
fmt.Println(ua.Android())

// 获取 iOS User-Agent
fmt.Println(ua.IOS())
```

### 获取特定平台类型的 User-Agent

```go
ua := fakeuseragent.New()

// 获取移动端 User-Agent
fmt.Println(ua.Mobile())

// 获取桌面端 User-Agent
fmt.Println(ua.Desktop())

// 获取平板端 User-Agent
fmt.Println(ua.Tablet())
```

### 动态修改选项

```go
ua := fakeuseragent.New()

// 设置允许的浏览器列表
ua.SetBrowsers([]string{"Chrome", "Safari"})

// 设置允许的操作系统
ua.SetOS([]string{"Windows", "Mac OS X"})

// 设置允许的平台类型
ua.SetPlatforms([]string{"desktop"})

// 设置最低浏览器版本
ua.SetMinVersion(120.0)

// 设置最低使用百分比
ua.SetMinPercentage(0.5)

// 设置备用 User-Agent
ua.SetFallback("Custom User Agent")
```

### 支持的浏览器

- Google
- Chrome
- Chrome Mobile
- Chrome Mobile iOS
- Firefox
- Firefox Mobile
- Firefox iOS
- Safari
- Mobile Safari
- Edge
- Edge Mobile
- Opera
- Opera Mobile
- Android
- Yandex Browser
- Samsung Internet
- Mobile Safari UI/WKWebView
- DuckDuckGo Mobile
- MiuiBrowser
- Whale
- Twitter
- Facebook
- Amazon Silk

### 支持的操作系统

- Windows
- Linux
- Ubuntu
- Chrome OS
- Mac OS X
- Android
- iOS

### 支持的平台类型

- desktop
- mobile
- tablet

## API 参考

### 类型

```go
// BrowserUserAgentData 表示浏览器 User-Agent 数据
type BrowserUserAgentData struct {
    UserAgent                string  // User-Agent 字符串
    Percent                  float64 // 使用百分比
    Type                     string  // 设备类型 (desktop/mobile/tablet)
    DeviceBrand              string  // 设备品牌
    Browser                  string  // 浏览器名称
    BrowserVersion           string  // 浏览器版本
    BrowserVersionMajorMinor float64 // 浏览器主版本号
    OS                       string  // 操作系统
    OSVersion                string  // 操作系统版本
    Platform                 string  // 平台信息
}

// Options 表示 UserAgent 配置选项
type Options struct {
    Browsers      []string  // 允许的浏览器列表
    OS            []string  // 允许的操作系统列表
    Platforms     []string  // 允许的平台列表
    MinVersion    float64   // 最低浏览器版本
    MinPercentage float64   // 最低使用百分比
    Fallback      string    // 备用 User-Agent
}
```

### 函数

```go
// DefaultOptions 返回默认配置选项
func DefaultOptions() Options

// New 创建带有默认选项的 UserAgent 实例
func New() *UserAgent

// NewWithOptions 创建带有自定义选项的 UserAgent 实例
func NewWithOptions(opts Options) *UserAgent
```

### 方法

#### 获取 User-Agent 字符串

```go
// Random 返回随机 User-Agent 字符串
func (ua *UserAgent) Random() string

// Chrome 返回随机 Chrome User-Agent 字符串
func (ua *UserAgent) Chrome() string

// Firefox 返回随机 Firefox User-Agent 字符串
func (ua *UserAgent) Firefox() string

// Safari 返回随机 Safari User-Agent 字符串
func (ua *UserAgent) Safari() string

// Opera 返回随机 Opera User-Agent 字符串
func (ua *UserAgent) Opera() string

// Edge 返回随机 Edge User-Agent 字符串
func (ua *UserAgent) Edge() string

// Android 返回随机 Android User-Agent 字符串
func (ua *UserAgent) Android() string

// IOS 返回随机 iOS User-Agent 字符串
func (ua *UserAgent) IOS() string

// Mobile 返回随机移动端 User-Agent 字符串
func (ua *UserAgent) Mobile() string

// Desktop 返回随机桌面端 User-Agent 字符串
func (ua *UserAgent) Desktop() string

// Tablet 返回随机平板端 User-Agent 字符串
func (ua *UserAgent) Tablet() string
```

#### 获取完整数据

```go
// GetBrowser 返回指定浏览器的完整 User-Agent 数据
func (ua *UserAgent) GetBrowser(browsers ...string) BrowserUserAgentData

// GetRandom 返回随机 User-Agent 的完整数据
func (ua *UserAgent) GetRandom() BrowserUserAgentData

// GetChrome 返回随机 Chrome 的完整数据
func (ua *UserAgent) GetChrome() BrowserUserAgentData

// GetFirefox 返回随机 Firefox 的完整数据
func (ua *UserAgent) GetFirefox() BrowserUserAgentData

// GetSafari 返回随机 Safari 的完整数据
func (ua *UserAgent) GetSafari() BrowserUserAgentData

// GetOpera 返回随机 Opera 的完整数据
func (ua *UserAgent) GetOpera() BrowserUserAgentData

// GetEdge 返回随机 Edge 的完整数据
func (ua *UserAgent) GetEdge() BrowserUserAgentData

// GetAndroid 返回随机 Android 的完整数据
func (ua *UserAgent) GetAndroid() BrowserUserAgentData

// GetIOS 返回随机 iOS 的完整数据
func (ua *UserAgent) GetIOS() BrowserUserAgentData

// GetMobile 返回随机移动端的完整数据
func (ua *UserAgent) GetMobile() BrowserUserAgentData

// GetDesktop 返回随机桌面端的完整数据
func (ua *UserAgent) GetDesktop() BrowserUserAgentData

// GetTablet 返回随机平板端的完整数据
func (ua *UserAgent) GetTablet() BrowserUserAgentData
```

#### 配置方法

```go
// SetFallback 设置备用 User-Agent
func (ua *UserAgent) SetFallback(fallback string)

// SetBrowsers 设置允许的浏览器列表
func (ua *UserAgent) SetBrowsers(browsers []string)

// SetOS 设置允许的操作系统列表
func (ua *UserAgent) SetOS(os []string)

// SetPlatforms 设置允许的平台列表
func (ua *UserAgent) SetPlatforms(platforms []string)

// SetMinVersion 设置最低浏览器版本
func (ua *UserAgent) SetMinVersion(version float64)

// SetMinPercentage 设置最低使用百分比
func (ua *UserAgent) SetMinPercentage(percentage float64)
```

## 运行示例

```bash
cd examples
go run main.go
```

## 运行测试

```bash
# 运行所有测试
go test -v

# 运行基准测试
go test -bench=.

# 运行测试并查看覆盖率
go test -cover
```

## 许可证

本项目参考 [fake-useragent/fake-useragent](https://github.com/fake-useragent/fake-useragent) 实现，数据来源于 [Intoli LLC](https://github.com/intoli/user-agents)。
