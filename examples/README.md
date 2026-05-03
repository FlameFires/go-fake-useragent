# go-fake-useragent 示例

本目录包含两个示例程序，演示 `go-fake-useragent` 库的使用方法。

## 运行环境

- Go 1.26+
- 依赖自动从网络拉取，无需额外配置

## 快速开始

```bash
# 进入示例目录
cd examples

# 运行示例 1：基础用法（首次运行会自动下载依赖）
go run example_1.go

# 运行示例 2：功能测试（含实际 HTTP 请求）
go run example_2.go
```

## 示例说明

### example_1.go — 基础用法

演示库的基本 API：

| 功能 | 方法 |
|------|------|
| 随机 User-Agent | `ua.Random()` |
| 按浏览器获取 | `ua.Chrome()` `ua.Firefox()` `ua.Safari()` `ua.Opera()` `ua.Edge()` |
| 按操作系统获取 | `ua.Android()` `ua.IOS()` |
| 按平台类型获取 | `ua.Mobile()` `ua.Desktop()` `ua.Tablet()` |
| 获取完整数据 | `ua.GetRandom()` 返回 `BrowserUserAgentData` 结构体 |
| 自定义选项 | `NewWithOptions(opts)` 创建带过滤条件的实例 |
| 动态修改选项 | `ua.SetBrowsers()` `ua.SetOS()` `ua.SetPlatforms()` 等方法 |

### example_2.go — 功能测试

覆盖更全面的测试场景：

1. 所有浏览器类型输出
2. 操作系统过滤
3. 平台类型过滤
4. 完整数据获取
5. 自定义选项过滤
6. **实际 HTTP 请求测试**（向 httpbin.org 发送请求验证 User-Agent 可用性）

## 预期输出示例

```
随机 User-Agent:
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.6778.140 Safari/537.36

Chrome User-Agent:
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.6723.92 Safari/537.36

Android User-Agent:
Mozilla/5.0 (Linux; Android 14; Pixel 8 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.6723.102 Mobile Safari/537.36
```

## 自定义选项

```go
opts := fakeuseragent.DefaultOptions()
opts.Browsers = []string{"Chrome", "Firefox"}  // 仅限 Chrome 和 Firefox
opts.OS = []string{"Windows"}                    // 仅限 Windows 系统
opts.Platforms = []string{"desktop"}             // 仅限桌面端
opts.MinVersion = 120.0                          // 最低版本 120

ua := fakeuseragent.NewWithOptions(opts)
fmt.Println(ua.Random())
```

## 在你的项目中使用

```go
import "github.com/flamefires/go-fake-useragent"

func main() {
    ua := fakeuseragent.New()
    req, _ := http.NewRequest("GET", "https://example.com", nil)
    req.Header.Set("User-Agent", ua.Random())
    // ...
}
```
