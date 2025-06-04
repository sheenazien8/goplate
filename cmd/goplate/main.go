package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	githubRepo = "https://github.com/sheenazien8/goplate"
	zipURL     = "https://github.com/sheenazien8/goplate/archive/refs/heads/master.zip"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: goplate <project-name>")
		os.Exit(1)
	}

	projectName := os.Args[1]

	fmt.Printf("Creating new GoPlate project: %s\n", projectName)

	if err := createProject(projectName); err != nil {
		fmt.Printf("Error creating project: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Project '%s' created successfully!\n", projectName)
	fmt.Printf("📁 cd %s\n", projectName)
	fmt.Printf("⚙️  go mod tidy\n")
	fmt.Printf("🔧 cp .env.example .env\n")
	fmt.Printf("🚀 go run main.go\n")
}

func createProject(projectName string) error {
	if _, err := os.Stat(projectName); !os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' already exists", projectName)
	}

	tempDir, err := os.MkdirTemp("", "goplate-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	fmt.Println("📥 Downloading boilerplate...")
	zipPath := filepath.Join(tempDir, "goplate.zip")
	if err := downloadFile(zipURL, zipPath); err != nil {
		return fmt.Errorf("failed to download boilerplate: %w", err)
	}

	fmt.Println("📦 Extracting files...")
	extractPath := filepath.Join(tempDir, "extracted")
	if err := unzip(zipPath, extractPath); err != nil {
		return fmt.Errorf("failed to extract archive: %w", err)
	}

	sourcePath := filepath.Join(extractPath, "goplate-master")
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("unexpected archive structure")
	}

	fmt.Println("🔧 Setting up project...")
	if err := os.Rename(sourcePath, projectName); err != nil {
		return fmt.Errorf("failed to move project: %w", err)
	}

	if err := setupProject(projectName); err != nil {
		return fmt.Errorf("failed to setup project: %w", err)
	}

	return nil
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	os.MkdirAll(dest, 0755)

	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid file path: %s", f.Name)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.FileInfo().Mode())
			return nil
		}

		os.MkdirAll(filepath.Dir(path), 0755)

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.FileInfo().Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, rc)
		return err
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func setupProject(projectPath string) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your module name (e.g., github.com/username/project): ")
	moduleName, _ := reader.ReadString('\n')
	moduleName = strings.TrimSpace(moduleName)

	if moduleName == "" {
		moduleName = fmt.Sprintf("github.com/username/%s", filepath.Base(projectPath))
		fmt.Printf("Using default module name: %s\n", moduleName)
	}

	if err := updateGoMod(projectPath, moduleName); err != nil {
		return fmt.Errorf("failed to update go.mod: %w", err)
	}

	if err := updateImports(projectPath, moduleName); err != nil {
		return fmt.Errorf("failed to update imports: %w", err)
	}

	if err := createEnvExample(projectPath); err != nil {
		return fmt.Errorf("failed to create .env.example: %w", err)
	}

	if err := removeInstallerFiles(projectPath); err != nil {
		return fmt.Errorf("failed to cleanup installer files: %w", err)
	}

	return nil
}

func updateGoMod(projectPath, moduleName string) error {
	goModPath := filepath.Join(projectPath, "go.mod")

	content, err := os.ReadFile(goModPath)
	if err != nil {
		return err
	}

	newContent := strings.Replace(string(content), "github.com/sheenazien8/goplate", moduleName, -1)

	return os.WriteFile(goModPath, []byte(newContent), 0644)
}

func updateImports(projectPath, moduleName string) error {
	return filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		newContent := strings.Replace(string(content), "github.com/sheenazien8/goplate", moduleName, -1)

		if string(content) != newContent {
			return os.WriteFile(path, []byte(newContent), info.Mode())
		}

		return nil
	})
}

func createEnvExample(projectPath string) error {
	envExample := `APP_NAME=GoPlate
APP_ENV=local
APP_DEBUG=true
APP_URL=http://localhost
APP_PORT=8080
APP_SCREET=your-secret-key-here

DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE=goplate
DB_USERNAME=root
DB_PASSWORD=
`

	envPath := filepath.Join(projectPath, ".env.example")
	return os.WriteFile(envPath, []byte(envExample), 0644)
}

func removeInstallerFiles(projectPath string) error {
	filesToRemove := []string{
		filepath.Join(projectPath, "cmd"),
		filepath.Join(projectPath, "install.sh"),
	}

	for _, file := range filesToRemove {
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			os.RemoveAll(file)
		}
	}

	return nil
}
