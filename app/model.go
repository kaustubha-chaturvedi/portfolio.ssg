package app

import (
	"bufio"
	"html/template"
	"log"
	"os"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Author      string
	Date        string
	Slug        string
	HeaderImage string
	Tags		[]string
	Content     template.HTML
	ReadTime    string
}

type Project struct {
	Title       string
	Description string
	Author      string
	Date        string
	URL         string
	HeaderImage string
	Tags        []string
	Slug		string
}


type AboutData struct {
	Name        string
	Tagline     string
	Description string
	Image       string
	SocialMedia map[string]string
	Content     template.HTML
	Projects    []interface{}
}

func parseProject(header []string) Project {
	var project Project
	for _, line := range header {
		switch {
		case strings.Contains(line, "title:"):
			project.Title = strings.TrimSpace(strings.TrimPrefix(line, "title:"))
		case strings.Contains(line, "description:"):
			project.Description = strings.TrimSpace(strings.TrimPrefix(line, "description:"))
		case strings.Contains(line, "author:"):
			project.Author = strings.TrimSpace(strings.TrimPrefix(line, "author:"))
		case strings.Contains(line, "date:"):
			project.Date = strings.TrimSpace(strings.TrimPrefix(line, "date:"))
		case strings.Contains(line, "url:"):
			project.URL = strings.TrimSpace(strings.TrimPrefix(line, "url:"))
		case strings.Contains(line, "tags:"):
			project.Tags = strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "tags:")), ",")
		case strings.Contains(line, "header-image:"):
			project.HeaderImage = strings.TrimSpace(strings.TrimPrefix(line, "header-image:"))
		}
	}
	return project
}

func parsePost(header []string) Post {
	var post Post
	for _, line := range header {
		switch {
		case strings.Contains(line, "title:"):
			post.Title = strings.TrimSpace(strings.TrimPrefix(line, "title:"))
		case strings.Contains(line, "description:"):
			post.Description = strings.TrimSpace(strings.TrimPrefix(line, "description:"))
		case strings.Contains(line, "author:"):
			post.Author = strings.TrimSpace(strings.TrimPrefix(line, "author:"))
		case strings.Contains(line, "date:"):
			post.Date = strings.TrimSpace(strings.TrimPrefix(line, "date:"))
		case strings.Contains(line, "header-image:"):
			post.HeaderImage = strings.TrimSpace(strings.TrimPrefix(line, "header-image:"))
		case strings.Contains(line, "tags:"):
			post.Tags = strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "tags:")), ",")

		}
	}
	return post
}

func countWords(text string) int {
    words := strings.Fields(text)
    return len(words)
}
// optional parameter Projects
func parseAbout(file *os.File, Projects []interface{}) AboutData {
	var aboutData AboutData
	var inHeader bool
	var header, body []string
	socialMedia := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			inHeader = !inHeader
			continue
		}
		if inHeader {
			header = append(header, line)
		} else {
			body = append(body, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, line := range header {
		switch {
		case strings.Contains(line, "name:"):
			aboutData.Name = strings.TrimSpace(strings.TrimPrefix(line, "name:"))
		case strings.Contains(line, "tagline:"):
			aboutData.Tagline = strings.TrimSpace(strings.TrimPrefix(line, "tagline:"))
		case strings.Contains(line, "description:"):
			aboutData.Description = strings.TrimSpace(strings.TrimPrefix(line, "description:"))
		case strings.Contains(line, "image:"):
			aboutData.Image = strings.TrimSpace(strings.TrimPrefix(line, "image:"))
		case strings.Contains(line, "github:"):
			socialMedia["github"] = strings.TrimSpace(strings.TrimPrefix(line, "github:"))
		case strings.Contains(line, "linkedin:"):
			socialMedia["linkedin"] = strings.TrimSpace(strings.TrimPrefix(line, "linkedin:"))
		case strings.Contains(line, "twitter:"):
			socialMedia["twitter"] = strings.TrimSpace(strings.TrimPrefix(line, "twitter:"))
		case strings.Contains(line, "instagram:"):
			socialMedia["instagram"] = strings.TrimSpace(strings.TrimPrefix(line, "instagram:"))
		case strings.Contains(line, "facebook:"):
			socialMedia["facebook"] = strings.TrimSpace(strings.TrimPrefix(line, "facebook:"))
		case strings.Contains(line, "email:"):
			socialMedia["email"] = strings.TrimSpace(strings.TrimPrefix(line, "email:"))
		case strings.Contains(line, "dribbble:"):
			socialMedia["dribbble"] = strings.TrimSpace(strings.TrimPrefix(line, "dribbble:"))
		case strings.Contains(line, "behance:"):
			socialMedia["behance"] = strings.TrimSpace(strings.TrimPrefix(line, "behance:"))	
		}
	}
	content := strings.TrimSpace(strings.Join(body, "\n"))
	aboutData.Content = template.HTML(ParseMarkdown(content))
	aboutData.SocialMedia = socialMedia
	aboutData.Image = strings.ReplaceAll(aboutData.Image, " ", "%20")
	aboutData.Projects = Projects
	return aboutData
}