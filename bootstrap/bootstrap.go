package bootstrap

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/sheenazien8/goplate/db"
	"github.com/sheenazien8/goplate/env"
	"github.com/sheenazien8/goplate/logs"
	"github.com/sheenazien8/goplate/pkg/assets"
	"github.com/sheenazien8/goplate/pkg/inertia"
	"github.com/sheenazien8/goplate/pkg/queue"
	"github.com/sheenazien8/goplate/pkg/scheduler"
	"github.com/sheenazien8/goplate/pkg/utils"
	"github.com/sheenazien8/goplate/router"
)

func App() *fiber.App {
	screet := env.Get("APP_SCREET")
	if screet == "" {
		logs.Fatal("You must generate the screet key first")
	}

	engine := html.New("./templates", ".html")

	// Add Inertia template functions
	engine.AddFunc("marshal", func(v any) template.JS {
		js, err := json.Marshal(v)
		if err != nil {
			return template.JS("")
		}
		return template.JS(js)
	})

	engine.AddFunc("raw", func(v any) template.HTML {
		if str, ok := v.(string); ok {
			return template.HTML(str)
		}
		return template.HTML("")
	})

	// Add asset helper functions
	engine.AddFunc("asset", func(src string) string {
		logs.Info(fmt.Sprintf("Requesting asset for source: %s", src)) // Debug log
		return assets.GetAsset(src)
	})
	engine.AddFunc("assetCSS", func(src string) []string {
		return assets.GetAssetCSS(src)
	})
	engine.AddFunc("getenv", func(key string) string {
		return os.Getenv(key)
	})

	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var errResponse utils.GlobalErrorHandlerResp
			if json.Unmarshal([]byte(err.Error()), &errResponse) != nil {
				var e *fiber.Error
				code := fiber.StatusInternalServerError
				message := "Internal Server Error"
				if errors.As(err, &e) {
					code = e.Code
					message = e.Message
				}
				logs.Error(err)

				return c.Status(code).JSON(fiber.Map{
					"success": false,
					"message": message,
					"error":   message,
				})
			}
			return c.Status(errResponse.Status).JSON(errResponse)
		},
	})

	// Connect DB (can be swapped with test DB)
	db.ConnectDB()

	// Initialize Inertia
	inertia.Init()

	// Load Vite manifest for asset helpers (ignore errors during development)
	_ = assets.LoadManifest()

	// Setup routes
	router.SetupRouter(app)

	// Start background jobs (can be skipped in tests)
	q := queue.New(100)
	q.Start(5)

	sch := scheduler.New()
	sch.RunTasks()
	sch.Start()

	return app
}

func Init() {
	screet := env.Get("APP_SCREET")
	if screet == "" {
		logs.Fatal("You must generate the screet key first")
	}

	// Connect DB for console commands
	db.ConnectDB()

	// Initialize Inertia
	inertia.Init()

	// Load Vite manifest for asset helpers (ignore errors during development)
	_ = assets.LoadManifest()
}
