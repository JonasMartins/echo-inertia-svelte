// Package web
package web

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"echo-inertia.com/src/internal/configs"
	"echo-inertia.com/src/pkg/utils"
	inertia "github.com/kohkimakimoto/inertia-echo/v2"
	"github.com/labstack/echo/v4"
)

// 1. Embed the directories.
// We embed 'public' for assets and 'views' for the base HTML.
//
//go:embed public/build/* views/*.html
var embeddedFiles embed.FS

func SetRender(e *echo.Echo, r *inertia.HTMLRenderer, cfg *configs.Config) {

	// 2. Configure built-in Vite support
	r.ViteBasePath = "/build" // Matches vite.config.js outDir: 'public/build'

	if cfg.Env == "development" {
		log.Println("running in dev mode")
		r.Debug = true
		r.ViteDevServerURL = "http://localhost:5173"
		path, err := utils.GetFilePath([]string{"src", "services", "web"})
		if err != nil {
			utils.FatalResult("error getting abs template path", err)
		}

		r.MustParseGlob(path + "/views/*.html")
		// Serve static files from disk in dev
		e.Static("/build", path+"/public/build")
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
}
