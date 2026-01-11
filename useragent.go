package fakeuseragent

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"sync"
)

//go:embed data/browsers.jsonl
var browsersData string

// BrowserUserAgentData represents a browser user agent entry.
type BrowserUserAgentData struct {
	UserAgent                string  `json:"useragent"`
	Percent                  float64 `json:"percent"`
	Type                     string  `json:"type"`
	DeviceBrand              string  `json:"device_brand"`
	Browser                  string  `json:"browser"`
	BrowserVersion           string  `json:"browser_version"`
	BrowserVersionMajorMinor float64 `json:"browser_version_major_minor"`
	OS                       string  `json:"os"`
	OSVersion                string  `json:"os_version"`
	Platform                 string  `json:"platform"`
}

// Options represents configuration options for UserAgent.
type Options struct {
	Browsers      []string
	OS            []string
	Platforms     []string
	MinVersion    float64
	MinPercentage float64
	Fallback      string
}

// DefaultOptions returns the default options for UserAgent.
func DefaultOptions() Options {
	return Options{
		Browsers: []string{
			"Google", "Chrome", "Firefox", "Edge", "Opera", "Safari",
			"Android", "Yandex Browser", "Samsung Internet", "Opera Mobile",
			"Mobile Safari", "Firefox Mobile", "Firefox iOS", "Chrome Mobile",
			"Chrome Mobile iOS", "Mobile Safari UI/WKWebView", "Edge Mobile",
			"DuckDuckGo Mobile", "MiuiBrowser", "Whale", "Twitter", "Facebook",
			"Amazon Silk",
		},
		OS: []string{
			"Windows", "Linux", "Ubuntu", "Chrome OS", "Mac OS X", "Android", "iOS",
		},
		Platforms:     []string{"desktop", "mobile", "tablet"},
		MinVersion:    0.0,
		MinPercentage: 0.0,
		Fallback: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 " +
			"(KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0",
	}
}

// UserAgent is a fake user agent retriever.
type UserAgent struct {
	options       Options
	dataBrowsers  []BrowserUserAgentData
	fallbackData  BrowserUserAgentData
	mu            sync.RWMutex
	randSource    rand.Source
}

// New creates a new UserAgent instance with default options.
func New() *UserAgent {
	return NewWithOptions(DefaultOptions())
}

// NewWithOptions creates a new UserAgent instance with custom options.
func NewWithOptions(opts Options) *UserAgent {
	data, err := loadData()
	if err != nil {
		panic(fmt.Sprintf("failed to load browsers data: %v", err))
	}

	ua := &UserAgent{
		options:      opts,
		dataBrowsers: data,
		randSource:   rand.NewSource(rand.Int63()),
		fallbackData: BrowserUserAgentData{
			UserAgent:                opts.Fallback,
			Percent:                  100.0,
			Type:                     "desktop",
			DeviceBrand:              "",
			Browser:                  "Edge",
			BrowserVersion:           "122.0.0.0",
			BrowserVersionMajorMinor: 122.0,
			OS:                       "win32",
			OSVersion:                "10",
			Platform:                 "Win32",
		},
	}

	return ua
}

// loadData loads the browsers data from the embedded file.
func loadData() ([]BrowserUserAgentData, error) {
	var data []BrowserUserAgentData
	lines := strings.Split(browsersData, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var entry BrowserUserAgentData
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			return nil, fmt.Errorf("failed to parse line: %w", err)
		}

		// Normalize empty strings to actual empty values
		if entry.DeviceBrand == "null" {
			entry.DeviceBrand = ""
		}
		if entry.Browser == "null" {
			entry.Browser = ""
		}
		if entry.OS == "null" {
			entry.OS = ""
		}
		if entry.OSVersion == "null" {
			entry.OSVersion = ""
		}

		data = append(data, entry)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("data list is empty")
	}

	return data, nil
}

// filterUserAgents filters the user agents based on the configured options.
func (ua *UserAgent) filterUserAgents(browsersToFilter []string) []BrowserUserAgentData {
	ua.mu.RLock()
	defer ua.mu.RUnlock()

	var result []BrowserUserAgentData

	for _, entry := range ua.dataBrowsers {
		// Check if browser is in allowed browsers list
		if !slices.Contains(ua.options.Browsers, entry.Browser) {
			continue
		}

		// Check if OS is in allowed OS list
		if !slices.Contains(ua.options.OS, entry.OS) {
			continue
		}

		// Check if type is in allowed platforms list
		if !slices.Contains(ua.options.Platforms, entry.Type) {
			continue
		}

		// Check minimum version
		if entry.BrowserVersionMajorMinor < ua.options.MinVersion {
			continue
		}

		// Check minimum percentage
		if entry.Percent < ua.options.MinPercentage {
			continue
		}

		// Filter by specific browser if provided
		if len(browsersToFilter) > 0 {
			found := false
			for _, b := range browsersToFilter {
				if strings.EqualFold(entry.Browser, b) {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		result = append(result, entry)
	}

	return result
}

// GetBrowser returns a user agent data for the specified browser(s).
// If browsers is empty, returns a random user agent from all allowed browsers.
func (ua *UserAgent) GetBrowser(browsers ...string) BrowserUserAgentData {
	var filtered []BrowserUserAgentData

	if len(browsers) == 0 || (len(browsers) == 1 && strings.ToLower(browsers[0]) == "random") {
		filtered = ua.filterUserAgents(nil)
	} else {
		filtered = ua.filterUserAgents(browsers)
	}

	if len(filtered) == 0 {
		return ua.fallbackData
	}

	ua.mu.Lock()
	r := rand.New(ua.randSource)
	ua.mu.Unlock()

	idx := r.Intn(len(filtered))
	return filtered[idx]
}

// Random returns a random user agent string.
func (ua *UserAgent) Random() string {
	return ua.GetBrowser("random").UserAgent
}

// Chrome returns a random Chrome user agent string.
func (ua *UserAgent) Chrome() string {
	return ua.GetBrowser("Chrome", "Chrome Mobile", "Chrome Mobile iOS").UserAgent
}

// Firefox returns a random Firefox user agent string.
func (ua *UserAgent) Firefox() string {
	return ua.GetBrowser("Firefox", "Firefox Mobile", "Firefox iOS").UserAgent
}

// Safari returns a random Safari user agent string.
func (ua *UserAgent) Safari() string {
	return ua.GetBrowser("Safari", "Mobile Safari").UserAgent
}

// Opera returns a random Opera user agent string.
func (ua *UserAgent) Opera() string {
	return ua.GetBrowser("Opera", "Opera Mobile").UserAgent
}

// Edge returns a random Edge user agent string.
func (ua *UserAgent) Edge() string {
	return ua.GetBrowser("Edge", "Edge Mobile").UserAgent
}

// InternetExplorer returns a random Edge user agent string (alias for Edge).
func (ua *UserAgent) InternetExplorer() string {
	return ua.Edge()
}

// GetRandom returns a random user agent data.
func (ua *UserAgent) GetRandom() BrowserUserAgentData {
	return ua.GetBrowser("random")
}

// GetChrome returns a random Chrome user agent data.
func (ua *UserAgent) GetChrome() BrowserUserAgentData {
	return ua.GetBrowser("Chrome", "Chrome Mobile", "Chrome Mobile iOS")
}

// GetFirefox returns a random Firefox user agent data.
func (ua *UserAgent) GetFirefox() BrowserUserAgentData {
	return ua.GetBrowser("Firefox", "Firefox Mobile", "Firefox iOS")
}

// GetSafari returns a random Safari user agent data.
func (ua *UserAgent) GetSafari() BrowserUserAgentData {
	return ua.GetBrowser("Safari", "Mobile Safari")
}

// GetOpera returns a random Opera user agent data.
func (ua *UserAgent) GetOpera() BrowserUserAgentData {
	return ua.GetBrowser("Opera", "Opera Mobile")
}

// GetEdge returns a random Edge user agent data.
func (ua *UserAgent) GetEdge() BrowserUserAgentData {
	return ua.GetBrowser("Edge", "Edge Mobile")
}

// Android returns a random Android user agent string.
func (ua *UserAgent) Android() string {
	ua.mu.RLock()
	originalOS := ua.options.OS
	ua.mu.RUnlock()

	ua.SetOS([]string{"Android"})
	result := ua.Random()
	ua.SetOS(originalOS)

	return result
}

// IOS returns a random iOS user agent string.
func (ua *UserAgent) IOS() string {
	ua.mu.RLock()
	originalOS := ua.options.OS
	ua.mu.RUnlock()

	ua.SetOS([]string{"iOS"})
	result := ua.Random()
	ua.SetOS(originalOS)

	return result
}

// Mobile returns a random mobile user agent string.
func (ua *UserAgent) Mobile() string {
	ua.mu.RLock()
	originalPlatforms := ua.options.Platforms
	ua.mu.RUnlock()

	ua.SetPlatforms([]string{"mobile"})
	result := ua.Random()
	ua.SetPlatforms(originalPlatforms)

	return result
}

// Desktop returns a random desktop user agent string.
func (ua *UserAgent) Desktop() string {
	ua.mu.RLock()
	originalPlatforms := ua.options.Platforms
	ua.mu.RUnlock()

	ua.SetPlatforms([]string{"desktop"})
	result := ua.Random()
	ua.SetPlatforms(originalPlatforms)

	return result
}

// Tablet returns a random tablet user agent string.
func (ua *UserAgent) Tablet() string {
	ua.mu.RLock()
	originalPlatforms := ua.options.Platforms
	ua.mu.RUnlock()

	ua.SetPlatforms([]string{"tablet"})
	result := ua.Random()
	ua.SetPlatforms(originalPlatforms)

	return result
}

// GetAndroid returns a random Android user agent data.
func (ua *UserAgent) GetAndroid() BrowserUserAgentData {
	ua.mu.RLock()
	originalOS := ua.options.OS
	ua.mu.RUnlock()

	ua.SetOS([]string{"Android"})
	result := ua.GetRandom()
	ua.SetOS(originalOS)

	return result
}

// GetIOS returns a random iOS user agent data.
func (ua *UserAgent) GetIOS() BrowserUserAgentData {
	ua.mu.RLock()
	originalOS := ua.options.OS
	ua.mu.RUnlock()

	ua.SetOS([]string{"iOS"})
	result := ua.GetRandom()
	ua.SetOS(originalOS)

	return result
}

// GetMobile returns a random mobile user agent data.
func (ua *UserAgent) GetMobile() BrowserUserAgentData {
	ua.mu.RLock()
	originalPlatforms := ua.options.Platforms
	ua.mu.RUnlock()

	ua.SetPlatforms([]string{"mobile"})
	result := ua.GetRandom()
	ua.SetPlatforms(originalPlatforms)

	return result
}

// GetDesktop returns a random desktop user agent data.
func (ua *UserAgent) GetDesktop() BrowserUserAgentData {
	ua.mu.RLock()
	originalPlatforms := ua.options.Platforms
	ua.mu.RUnlock()

	ua.SetPlatforms([]string{"desktop"})
	result := ua.GetRandom()
	ua.SetPlatforms(originalPlatforms)

	return result
}

// GetTablet returns a random tablet user agent data.
func (ua *UserAgent) GetTablet() BrowserUserAgentData {
	ua.mu.RLock()
	originalPlatforms := ua.options.Platforms
	ua.mu.RUnlock()

	ua.SetPlatforms([]string{"tablet"})
	result := ua.GetRandom()
	ua.SetPlatforms(originalPlatforms)

	return result
}

// SetFallback sets the fallback user agent string.
func (ua *UserAgent) SetFallback(fallback string) {
	ua.mu.Lock()
	defer ua.mu.Unlock()

	ua.options.Fallback = fallback
	ua.fallbackData.UserAgent = fallback
}

// SetBrowsers sets the allowed browsers list.
func (ua *UserAgent) SetBrowsers(browsers []string) {
	ua.mu.Lock()
	defer ua.mu.Unlock()

	ua.options.Browsers = browsers
}

// SetOS sets the allowed operating systems list.
func (ua *UserAgent) SetOS(os []string) {
	ua.mu.Lock()
	defer ua.mu.Unlock()

	ua.options.OS = os
}

// SetPlatforms sets the allowed platforms list.
func (ua *UserAgent) SetPlatforms(platforms []string) {
	ua.mu.Lock()
	defer ua.mu.Unlock()

	ua.options.Platforms = platforms
}

// SetMinVersion sets the minimum browser version.
func (ua *UserAgent) SetMinVersion(version float64) {
	ua.mu.Lock()
	defer ua.mu.Unlock()

	ua.options.MinVersion = version
}

// SetMinPercentage sets the minimum usage percentage.
func (ua *UserAgent) SetMinPercentage(percentage float64) {
	ua.mu.Lock()
	defer ua.mu.Unlock()

	ua.options.MinPercentage = percentage
}
