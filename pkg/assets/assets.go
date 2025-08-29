package assets

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type ManifestEntry struct {
	File    string   `json:"file"`
	Name    string   `json:"name"`
	Src     string   `json:"src"`
	CSS     []string `json:"css,omitempty"`
	IsEntry bool     `json:"isEntry,omitempty"`
}

type Manifest map[string]ManifestEntry

var manifest Manifest

func LoadManifest() error {
	manifestPath := filepath.Join("public", ".vite", "manifest.json")

	data, err := os.ReadFile(manifestPath)
	if err != nil {
		manifest = make(Manifest)
		return err
	}

	return json.Unmarshal(data, &manifest)
}

func GetAsset(src string) string {
	if isDevelopment() {
		return "http://localhost:5173/" + src
	}

	if manifest == nil {
		LoadManifest()
	}

	if entry, exists := manifest[src]; exists {
		return "/" + entry.File
	}

	return "/assets/" + src
}

func isDevelopment() bool {
	env := os.Getenv("APP_ENV")
	return env == "local" || env == "development" || env == "dev"
}

func GetAssetCSS(src string) []string {
	if isDevelopment() {
		return []string{}
	}

	if manifest == nil {
		LoadManifest()
	}

	if entry, exists := manifest[src]; exists {
		var cssFiles []string
		for _, css := range entry.CSS {
			cssFiles = append(cssFiles, "/"+css)
		}
		return cssFiles
	}

	return []string{"/assets/app.css"}
}
