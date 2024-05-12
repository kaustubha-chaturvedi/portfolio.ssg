package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kaustubha-chaturvedi/expo.go/app"
)

func main() {
	app.EnsureDir(ContentDir)
	app.EnsureDir(PublicDir)
	app.CopyStatic(StaticDir, PublicStaticDir)

	
	allPosts, _ := app.GetAllData(PostsDir, PublicPostsDir, "posts.html", "page", PageSize, ThemeDir)
	allProjects, _ := app.GetAllData(ProjectsDir, PublicProjectsDir, "projects.html", "page", ProjectsPerPage, ThemeDir)

	_, err := app.GetAboutContent(filepath.Join(ContentDir, AboutMeMd),ThemeDir, PublicDir, allProjects)
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
