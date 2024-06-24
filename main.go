package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kaustubha-chaturvedi/expo.go/app"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
)

func getMimeType(ext string) string {
	switch ext {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".svg":
		return "image/svg+xml"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	default:
		return ""
	}
}

func outputMinifier() {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFunc("application/javascript", js.Minify)
	m.AddFunc("application/json", json.Minify)

	filepath.Walk("public", func(path string, info os.FileInfo, _ error) error {
		if !info.IsDir() {
			data, _ := os.ReadFile(path)
			ext := filepath.Ext(path)
			mimetype := getMimeType(ext)
			if mimetype != "" {
				minifiedData, _ := m.Bytes(mimetype, data)
				os.WriteFile(path, minifiedData, info.Mode())
			}
		}
		return nil
	})
	fmt.Println("Minified all files in public directory")
}

func main() {
	app.EnsureDir(ContentDir)
	app.EnsureDir(PublicDir)
	app.CopyStatic(StaticDir, PublicStaticDir)

	allPosts, _ := app.GetAllData(PostsDir, PublicPostsDir, "posts.html", "page", PageSize, ThemeDir)
	allProjects, _ := app.GetAllData(ProjectsDir, PublicProjectsDir, "projects.html", "page", ProjectsPerPage, ThemeDir)

	_, err := app.GetAboutContent(filepath.Join(ContentDir, AboutMeMd), ThemeDir, PublicDir, allProjects)
	if err != nil {
		log.Fatal(err)
	}
	for _, data := range allPosts {
		post, ok := data.(app.Post)
		if !ok {
			log.Fatal("Failed to assert type to Post")
		}
		outputPath := filepath.Join(PublicPostsDir, fmt.Sprintf("%s.html", strings.TrimSuffix(post.Slug, ".md")))
		app.RenderPage(filepath.Join(ThemeDir, "post.html"), outputPath, struct{ app.Post }{post})
	}
	outputMinifier()
	serve()
}

func serve() {
	dir := "public"
	http.Handle("/", http.FileServer(http.Dir(dir)))
	godotenv.Load()
	env := os.Getenv("ENV")
	if env == "DEV" {
		log.Printf("Server listening on http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		log.Println("Server can't be started in PROD.")
	}
}
