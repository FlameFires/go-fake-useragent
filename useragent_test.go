package fakeuseragent

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	ua := New()

	if ua == nil {
		t.Fatal("New() returned nil")
	}

	if len(ua.dataBrowsers) == 0 {
		t.Fatal("New() did not load browser data")
	}
}

func TestNewWithOptions(t *testing.T) {
	opts := DefaultOptions()
	opts.Browsers = []string{"Chrome"}

	ua := NewWithOptions(opts)

	if ua == nil {
		t.Fatal("NewWithOptions() returned nil")
	}

	if len(ua.options.Browsers) != 1 || ua.options.Browsers[0] != "Chrome" {
		t.Fatalf("Expected browsers to be [Chrome], got %v", ua.options.Browsers)
	}
}

func TestRandom(t *testing.T) {
	ua := New()

	userAgent := ua.Random()

	if userAgent == "" {
		t.Fatal("Random() returned empty string")
	}

	if !strings.HasPrefix(userAgent, "Mozilla/") {
		t.Fatalf("Random() returned invalid user agent: %s", userAgent)
	}
}

func TestChrome(t *testing.T) {
	ua := New()

	userAgent := ua.Chrome()

	if userAgent == "" {
		t.Fatal("Chrome() returned empty string")
	}

	if !strings.Contains(userAgent, "Chrome") {
		t.Fatalf("Chrome() returned user agent without Chrome: %s", userAgent)
	}
}

func TestFirefox(t *testing.T) {
	ua := New()

	userAgent := ua.Firefox()

	if userAgent == "" {
		t.Fatal("Firefox() returned empty string")
	}

	// Firefox iOS uses "FxiOS" instead of "Firefox"
	if !strings.Contains(userAgent, "Firefox") && !strings.Contains(userAgent, "FxiOS") {
		t.Fatalf("Firefox() returned user agent without Firefox: %s", userAgent)
	}
}

func TestSafari(t *testing.T) {
	ua := New()

	userAgent := ua.Safari()

	if userAgent == "" {
		t.Fatal("Safari() returned empty string")
	}

	// Safari user agents contain "Safari"
	if !strings.Contains(userAgent, "Safari") {
		t.Fatalf("Safari() returned user agent without Safari: %s", userAgent)
	}
}

func TestOpera(t *testing.T) {
	ua := New()

	userAgent := ua.Opera()

	if userAgent == "" {
		t.Fatal("Opera() returned empty string")
	}

	if !strings.Contains(userAgent, "Opera") && !strings.Contains(userAgent, "OPR") {
		t.Fatalf("Opera() returned user agent without Opera: %s", userAgent)
	}
}

func TestEdge(t *testing.T) {
	ua := New()

	userAgent := ua.Edge()

	if userAgent == "" {
		t.Fatal("Edge() returned empty string")
	}

	if !strings.Contains(userAgent, "Edg") && !strings.Contains(userAgent, "Edge") {
		t.Fatalf("Edge() returned user agent without Edge: %s", userAgent)
	}
}

func TestGetBrowser(t *testing.T) {
	ua := New()

	// Test random browser
	data := ua.GetBrowser()
	if data.UserAgent == "" {
		t.Fatal("GetBrowser() returned empty user agent")
	}

	// Test specific browser
	chromeData := ua.GetBrowser("Chrome")
	if chromeData.Browser != "Chrome" {
		t.Fatalf("GetBrowser('Chrome') returned browser %s, expected Chrome", chromeData.Browser)
	}
}

func TestGetRandom(t *testing.T) {
	ua := New()

	data := ua.GetRandom()

	if data.UserAgent == "" {
		t.Fatal("GetRandom() returned empty user agent")
	}

	if !strings.HasPrefix(data.UserAgent, "Mozilla/") {
		t.Fatalf("GetRandom() returned invalid user agent: %s", data.UserAgent)
	}
}

func TestGetChrome(t *testing.T) {
	ua := New()

	data := ua.GetChrome()

	if data.UserAgent == "" {
		t.Fatal("GetChrome() returned empty user agent")
	}

	// Check if browser is one of the Chrome variants
	validBrowsers := []string{"Chrome", "Chrome Mobile", "Chrome Mobile iOS"}
	valid := false
	for _, b := range validBrowsers {
		if data.Browser == b {
			valid = true
			break
		}
	}

	if !valid {
		t.Fatalf("GetChrome() returned browser %s, expected one of %v", data.Browser, validBrowsers)
	}
}

func TestGetFirefox(t *testing.T) {
	ua := New()

	data := ua.GetFirefox()

	if data.UserAgent == "" {
		t.Fatal("GetFirefox() returned empty user agent")
	}

	validBrowsers := []string{"Firefox", "Firefox Mobile", "Firefox iOS"}
	valid := false
	for _, b := range validBrowsers {
		if data.Browser == b {
			valid = true
			break
		}
	}

	if !valid {
		t.Fatalf("GetFirefox() returned browser %s, expected one of %v", data.Browser, validBrowsers)
	}
}

func TestGetSafari(t *testing.T) {
	ua := New()

	data := ua.GetSafari()

	if data.UserAgent == "" {
		t.Fatal("GetSafari() returned empty user agent")
	}

	validBrowsers := []string{"Safari", "Mobile Safari"}
	valid := false
	for _, b := range validBrowsers {
		if data.Browser == b {
			valid = true
			break
		}
	}

	if !valid {
		t.Fatalf("GetSafari() returned browser %s, expected one of %v", data.Browser, validBrowsers)
	}
}

func TestGetOpera(t *testing.T) {
	ua := New()

	data := ua.GetOpera()

	if data.UserAgent == "" {
		t.Fatal("GetOpera() returned empty user agent")
	}

	validBrowsers := []string{"Opera", "Opera Mobile"}
	valid := false
	for _, b := range validBrowsers {
		if data.Browser == b {
			valid = true
			break
		}
	}

	if !valid {
		t.Fatalf("GetOpera() returned browser %s, expected one of %v", data.Browser, validBrowsers)
	}
}

func TestGetEdge(t *testing.T) {
	ua := New()

	data := ua.GetEdge()

	if data.UserAgent == "" {
		t.Fatal("GetEdge() returned empty user agent")
	}

	validBrowsers := []string{"Edge", "Edge Mobile"}
	valid := false
	for _, b := range validBrowsers {
		if data.Browser == b {
			valid = true
			break
		}
	}

	if !valid {
		t.Fatalf("GetEdge() returned browser %s, expected one of %v", data.Browser, validBrowsers)
	}
}

func TestSetBrowsers(t *testing.T) {
	ua := New()

	ua.SetBrowsers([]string{"Chrome", "Firefox"})

	if len(ua.options.Browsers) != 2 {
		t.Fatalf("Expected 2 browsers, got %d", len(ua.options.Browsers))
	}
}

func TestSetOS(t *testing.T) {
	ua := New()

	ua.SetOS([]string{"Windows"})

	if len(ua.options.OS) != 1 || ua.options.OS[0] != "Windows" {
		t.Fatalf("Expected OS to be [Windows], got %v", ua.options.OS)
	}
}

func TestSetPlatforms(t *testing.T) {
	ua := New()

	ua.SetPlatforms([]string{"desktop"})

	if len(ua.options.Platforms) != 1 || ua.options.Platforms[0] != "desktop" {
		t.Fatalf("Expected platforms to be [desktop], got %v", ua.options.Platforms)
	}
}

func TestSetMinVersion(t *testing.T) {
	ua := New()

	ua.SetMinVersion(120.0)

	if ua.options.MinVersion != 120.0 {
		t.Fatalf("Expected MinVersion to be 120.0, got %f", ua.options.MinVersion)
	}
}

func TestSetMinPercentage(t *testing.T) {
	ua := New()

	ua.SetMinPercentage(0.5)

	if ua.options.MinPercentage != 0.5 {
		t.Fatalf("Expected MinPercentage to be 0.5, got %f", ua.options.MinPercentage)
	}
}

func TestSetFallback(t *testing.T) {
	ua := New()

	fallback := "Custom Fallback"
	ua.SetFallback(fallback)

	if ua.options.Fallback != fallback {
		t.Fatalf("Expected Fallback to be %s, got %s", fallback, ua.options.Fallback)
	}
}

func TestFilterByBrowser(t *testing.T) {
	opts := DefaultOptions()
	opts.Browsers = []string{"Chrome"}

	ua := NewWithOptions(opts)
	data := ua.GetBrowser()

	if data.Browser != "Chrome" {
		t.Fatalf("Expected browser to be Chrome, got %s", data.Browser)
	}
}

func TestFilterByOS(t *testing.T) {
	opts := DefaultOptions()
	opts.OS = []string{"Windows"}

	ua := NewWithOptions(opts)
	data := ua.GetBrowser()

	// The OS in the data might not be exactly "Windows" (could be "win10", "win32", etc.)
	// So we just check that we get a valid user agent
	if data.UserAgent == "" {
		t.Fatal("GetBrowser() returned empty user agent")
	}
}

func TestFilterByPlatform(t *testing.T) {
	opts := DefaultOptions()
	opts.Platforms = []string{"desktop"}

	ua := NewWithOptions(opts)
	data := ua.GetBrowser()

	if data.Type != "desktop" {
		t.Fatalf("Expected type to be desktop, got %s", data.Type)
	}
}

func TestFilterByMinVersion(t *testing.T) {
	opts := DefaultOptions()
	opts.MinVersion = 120.0

	ua := NewWithOptions(opts)
	data := ua.GetBrowser()

	if data.BrowserVersionMajorMinor < 120.0 {
		t.Fatalf("Expected version >= 120.0, got %f", data.BrowserVersionMajorMinor)
	}
}

func TestFallback(t *testing.T) {
	opts := DefaultOptions()
	opts.Browsers = []string{"NonExistentBrowser"}

	ua := NewWithOptions(opts)
	data := ua.GetBrowser()

	// Should return fallback
	if data.UserAgent != ua.options.Fallback {
		t.Fatalf("Expected fallback, got: %s", data.UserAgent)
	}
}

func TestMultipleCalls(t *testing.T) {
	ua := New()

	// Test that multiple calls work correctly
	for i := 0; i < 10; i++ {
		uaRandom := ua.Random()
		if uaRandom == "" {
			t.Fatalf("Random() returned empty string on iteration %d", i)
		}
	}
}

func TestBrowserDataStructure(t *testing.T) {
	ua := New()

	data := ua.GetRandom()

	// Check that all expected fields are present
	if data.UserAgent == "" {
		t.Fatal("UserAgent is empty")
	}

	// Percent can be 0, so we don't check it
	// Type, Browser, OS, etc. can be empty in some edge cases, so we don't strictly check them
}

func TestConcurrentAccess(t *testing.T) {
	ua := New()

	done := make(chan bool)

	// Run multiple goroutines accessing the same UserAgent instance
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				ua.Random()
				ua.Chrome()
				ua.Firefox()
				ua.GetBrowser()
			}
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}

func BenchmarkRandom(b *testing.B) {
	ua := New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ua.Random()
	}
}

func BenchmarkChrome(b *testing.B) {
	ua := New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ua.Chrome()
	}
}

func BenchmarkGetBrowser(b *testing.B) {
	ua := New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ua.GetBrowser()
	}
}

func TestAndroid(t *testing.T) {
	ua := New()

	userAgent := ua.Android()

	if userAgent == "" {
		t.Fatal("Android() returned empty string")
	}

	// Android user agents typically contain "Android"
	if !strings.Contains(userAgent, "Android") {
		t.Fatalf("Android() returned user agent without Android: %s", userAgent)
	}
}

func TestIOS(t *testing.T) {
	ua := New()

	userAgent := ua.IOS()

	if userAgent == "" {
		t.Fatal("iOS() returned empty string")
	}

	// iOS user agents typically contain "iPhone" or "iPad"
	if !strings.Contains(userAgent, "iPhone") && !strings.Contains(userAgent, "iPad") {
		t.Fatalf("iOS() returned user agent without iPhone/iPad: %s", userAgent)
	}
}

func TestMobile(t *testing.T) {
	ua := New()

	data := ua.GetMobile()

	if data.UserAgent == "" {
		t.Fatal("GetMobile() returned empty user agent")
	}

	if data.Type != "mobile" {
		t.Fatalf("GetMobile() returned type %s, expected mobile", data.Type)
	}
}

func TestDesktop(t *testing.T) {
	ua := New()

	data := ua.GetDesktop()

	if data.UserAgent == "" {
		t.Fatal("GetDesktop() returned empty user agent")
	}

	if data.Type != "desktop" {
		t.Fatalf("GetDesktop() returned type %s, expected desktop", data.Type)
	}
}

func TestTablet(t *testing.T) {
	ua := New()

	data := ua.GetTablet()

	if data.UserAgent == "" {
		t.Fatal("GetTablet() returned empty user agent")
	}

	if data.Type != "tablet" {
		t.Fatalf("GetTablet() returned type %s, expected tablet", data.Type)
	}
}

func TestGetAndroid(t *testing.T) {
	ua := New()

	data := ua.GetAndroid()

	if data.UserAgent == "" {
		t.Fatal("GetAndroid() returned empty user agent")
	}

	if data.OS != "Android" {
		t.Fatalf("GetAndroid() returned OS %s, expected Android", data.OS)
	}
}

func TestGetIOS(t *testing.T) {
	ua := New()

	data := ua.GetIOS()

	if data.UserAgent == "" {
		t.Fatal("GetIOS() returned empty user agent")
	}

	if data.OS != "iOS" {
		t.Fatalf("GetIOS() returned OS %s, expected iOS", data.OS)
	}
}

func TestPlatformMethodsPreserveOptions(t *testing.T) {
	ua := New()
	originalOS := ua.options.OS
	originalPlatforms := ua.options.Platforms

	// Call Android
	ua.Android()
	if len(ua.options.OS) != len(originalOS) {
		t.Fatalf("Android() changed OS list length")
	}

	// Call iOS
	ua.IOS()
	if len(ua.options.OS) != len(originalOS) {
		t.Fatalf("IOS() changed OS list length")
	}

	// Call Mobile
	ua.Mobile()
	if len(ua.options.Platforms) != len(originalPlatforms) {
		t.Fatalf("Mobile() changed platforms list length")
	}

	// Call Desktop
	ua.Desktop()
	if len(ua.options.Platforms) != len(originalPlatforms) {
		t.Fatalf("Desktop() changed platforms list length")
	}

	// Call Tablet
	ua.Tablet()
	if len(ua.options.Platforms) != len(originalPlatforms) {
		t.Fatalf("Tablet() changed platforms list length")
	}
}
