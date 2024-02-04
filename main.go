package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kaustubha-chaturvedi/portfolio-ssg-go/app"
)

func main() {
	app.EnsureDir(app.ContentDir)
	app.EnsureDir(app.PublicDir)
	app.CopyStatic(app.StaticDir, app.PublicStaticDir)

	_, err := app.GetAboutContent(filepath.Join(app.ContentDir, "about.md"))
	if err != nil {
		log.Fatal(err)
	}

	allPosts, _ := app.GetAllData(app.PostsDir, app.PublicPostsDir, "posts.html", "page", app.PageSize)
	app.GetAllData(app.ProjectsDir, app.PublicProjectsDir, "projects.html", "page", app.ProjectsPerPage)

	for _, data := range allPosts {
		post, ok := data.(app.Post)
		if !ok {
			log.Fatal("Failed to assert type to Post")
		}
		outputPath := filepath.Join(app.PublicPostsDir, fmt.Sprintf("%s.html", strings.TrimSuffix(post.Slug, ".md")))
		app.RenderPage(filepath.Join("templates", "post.html"), outputPath, struct{ app.Post }{post})
	}
	serve()
}

func serve() {
	devFlag := flag.Bool("dev", false, "Run in development mode")
	flag.Parse()
	dir := "public"
	http.Handle("/", http.FileServer(http.Dir(dir)))

	if *devFlag {
		log.Printf("Server listening on http://localhost:8000")
		log.Fatal(http.ListenAndServe(":8000", nil))
	} else {
		log.Println("Server can't be started in PROD.")
	}
}
