package main

import "fmt"

var (
	PageSize          = 6
	ProjectsPerPage   = 6 // unused with glass theme
	StaticDir         = "static"
	AboutMeMd		  = "index.md"
	PublicDir         = "public"
	ContentDir        = "./content"
	ThemeDir          = "./theme/glass"
	PostsDir          = fmt.Sprintf("%s/posts", ContentDir)
	ProjectsDir       = fmt.Sprintf("%s/projects", ContentDir)
	PublicStaticDir   = fmt.Sprintf("%s/static", PublicDir)
	PublicPostsDir    = fmt.Sprintf("%s/posts", PublicDir)
	PublicProjectsDir = fmt.Sprintf("%s/projects", PublicDir)
)
