package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	inertia "github.com/kohkimakimoto/inertia-echo/v2"
	"github.com/labstack/echo/v4"
)

// 1. Embed the directories.
// We embed 'public' for assets and 'views' for the base HTML.
//
//go:embed public/build/* views/*.html
var embeddedFiles embed.FS

func main() {
	e := echo.New()
	isDev := os.Getenv("APP_ENV") == "development"

	// 1. Initialize the built-in HTML Renderer
	r := inertia.NewHTMLRenderer()

	// 2. Configure built-in Vite support
	r.ViteBasePath = "/build" // Matches vite.config.js outDir: 'public/build'

	if isDev {
		log.Println("running in dev mode")
		r.Debug = true
		r.ViteDevServerURL = "http://localhost:5173"
		r.MustParseGlob("views/*.html")

		// Serve static files from disk in dev
		e.Static("/build", "public/build")
	} else {
		// --- PRODUCTION MODE (Embedded) ---

		log.Println("running in prod mode")
		// A. Parse templates from the embedded filesystem
		r.MustParseFS(embeddedFiles, "views/*.html")

		// B. Parse the manifest from the embedded filesystem
		r.MustParseViteManifestFS(embeddedFiles, "public/build/.vite/manifest.json")

		// C. Serve static assets from the embedded 'public/build' folder
		// We use fs.Sub to "zoom in" on the subfolder we want to serve
		publicBuild, _ := fs.Sub(embeddedFiles, "public/build")
		assetHandler := http.FileServer(http.FS(publicBuild))
		e.GET("/build/*", echo.WrapHandler(http.StripPrefix("/build/", assetHandler)))
	}

	// 4. Set up Middleware
	e.Use(inertia.MiddlewareWithConfig(inertia.MiddlewareConfig{
		Renderer: r,
	}))

	e.GET("/", func(c echo.Context) error {
		return inertia.Render(c, "Home", map[string]any{
			"message": "Welcome to the Home Page!",
		})
	})

	e.GET("/about", func(c echo.Context) error {
		return inertia.Render(c, "About", map[string]any{
			"content": "We are building a highly performant Go web app using Svelte 5.",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
