package main

import (
	"fmt"
	"os"
)

var (
	version = "0.1.0-dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "new":
		if err := handleNewProject(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "version":
		showVersion()
	case "templates":
		if err := listTemplates(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "help", "--help", "-h":
		showHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", cmd)
		showHelp()
		os.Exit(1)
	}
}

func handleNewProject(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("project name is required")
	}

	projectName := args[0]
	fmt.Printf("Creating new GoPlate project: %s\n", projectName)
	fmt.Println("This will create a project using goplate-core...")

	// TODO: Implement project generation logic
	return fmt.Errorf("not implemented yet - coming in next phase")
}

func showVersion() {
	fmt.Printf("GoPlate CLI v%s\n", version)
	fmt.Printf("Built: %s\n", date)
	fmt.Printf("Commit: %s\n", commit)
	fmt.Println("goplate-core: not yet released")
}

func listTemplates() error {
	fmt.Println("Available Templates:")
	fmt.Println("  api     - REST API only (default)")
	fmt.Println("  full    - Full-stack application")
	fmt.Println("  micro   - Microservice template")

	// TODO: Load from actual template registry
	return nil
}

func showHelp() {
	fmt.Println(`GoPlate CLI - Go REST API Boilerplate Downloader

Usage:
  goplate <command> [arguments]

Commands:
  new <project-name>     Download and setup a new GoPlate project
  version               Show CLI version information
  templates             List available project templates
  help                  Show this help message

Flags for 'new' command:
  --template=<name>     Template to use (api, full, micro) [default: api]
  --db=<type>          Database type (postgres, mysql) [default: postgres]
  --module=<name>      Custom Go module name [default: project-name]
  --no-git            Skip git initialization
  --force             Overwrite existing directory

Examples:
  goplate new my-api
  goplate new my-app --template=full --db=mysql
  goplate new microservice --template=micro --no-git

After project creation, use the built-in generators:
  cd my-api
  go run main.go console make:model User
  go run main.go console make:controller UserController
  go run main.go console list                    # See all available commands

Note: This CLI downloads project templates. All generators (make:model, make:controller, 
etc.) are built into the project itself via the console command system.

For more information, visit: https://github.com/sheenazien8/goplate`)
}
