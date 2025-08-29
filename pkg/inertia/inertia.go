package inertia

import (
	"os"

	"github.com/sheenazien8/goplate/pkg/assets"
	"github.com/sheenazien8/inertia-go"
)

var Manager *inertia.Inertia
var FiberManager *FiberAdapter

func Init() {
	// Get the base URL from environment or use default
	url := os.Getenv("APP_URL")
	if url == "" {
		url = "http://localhost:3000"
	}

	// Initialize Inertia manager with the root template
	Manager = inertia.New(url, "./templates/app.html", "1.0")

	// Add custom template functions
	Manager.SharedFuncMap["asset"] = func(src string) string {
		return assets.GetAsset(src)
	}
	Manager.SharedFuncMap["assetCSS"] = func(src string) []string {
		return assets.GetAssetCSS(src)
	}
	Manager.SharedFuncMap["getenv"] = func(key string) string {
		return os.Getenv(key)
	}

	// Create Fiber adapter
	FiberManager = NewFiberAdapter(Manager)

	// Share some global props
	Manager.Share("appName", os.Getenv("APP_NAME"))

	// You can enable SSR if needed
	// Manager.EnableSsrWithDefault()
}
