package app

import "fmt"

var (
	PageSize          = 5
	ProjectsPerPage   = 6
	ContentDir        = "./content"
	PostsDir          = fmt.Sprintf("%s/posts", ContentDir)
	ProjectsDir       = fmt.Sprintf("%s/projects", ContentDir)
	StaticDir         = "static"
	PublicDir         = "public"
	PublicStaticDir   = fmt.Sprintf("%s/static", PublicDir)
	PublicPostsDir    = fmt.Sprintf("%s/posts", PublicDir)
	PublicProjectsDir = fmt.Sprintf("%s/projects", PublicDir)
)
