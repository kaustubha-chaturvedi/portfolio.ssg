package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	pageSize          = 5
	projectsPerPage   = 6
	postsDir          = "./posts"
	projectsDir       = "./projects"
	staticDir         = "static"
	publicStaticDir   = "public/static"
	publicPostsDir    = "public/posts"
	publicProjectsDir = "public/projects"
	aboutTemplate     = "public/index.html"
)

func main() {
	copyStatic(staticDir, publicStaticDir)

	e := echo.New()
	e.Static("/static", publicStaticDir)

	_, err := getAboutContent("about.md")
	if err != nil {
		log.Fatal(err)
	}
	allPosts, _ := getAllData(postsDir, publicPostsDir, "posts.html", "page", pageSize)
	getAllData(projectsDir, publicProjectsDir, "projects.html", "page", projectsPerPage)

	for _, data := range allPosts {
		post, ok := data.(Post)
		if !ok {
			log.Fatal("Failed to assert type to Post")
		}
		outputPath := filepath.Join(publicPostsDir, fmt.Sprintf("%s.html", strings.TrimSuffix(post.Slug, ".md")))
		renderPage(filepath.Join("templates", "post.html"), outputPath, struct{ Post }{post})
	}

	e.GET("/", func(c echo.Context) error {
		return c.File(aboutTemplate)
	})

	e.GET("/posts", func(c echo.Context) error {
		return c.File(fmt.Sprintf("%s/page1.html", publicPostsDir))
	})

	e.GET("/post/:slug", func(c echo.Context) error {
		slug := c.Param("slug")
		return c.File(fmt.Sprintf("%s/%s.html", publicPostsDir, slug))
	})

	e.GET("/posts/page/:page", func(c echo.Context) error {
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid page number")
		}
		return c.File(fmt.Sprintf("%s/page%d.html", publicPostsDir, page))
	})

	e.GET("/projects", func(c echo.Context) error {
		return c.File(fmt.Sprintf("%s/page1.html", publicProjectsDir))
	})

	e.GET("/projects/page/:page", func(c echo.Context) error {
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid page number")
		}
		return c.File(fmt.Sprintf("%s/page%d.html", publicProjectsDir, page))
	})

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}