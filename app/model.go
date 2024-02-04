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
	Content     template.HTML
	ReadTime    string
}

type Project struct {
	Title       string
	Description string
	Author      string
	Date        string
	URL         string
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
}

func parseProject(header, body []string) Project {
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
		}
	}
	return project
}

func parsePost(header, body []string) Post {
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
		}
	}
	return post
}

func countWords(text string) int {
    words := strings.Fields(text)
    return len(words)
}

func parseAbout(file *os.File) AboutData {
	var aboutData AboutData
	var inHeader bool
	var header, body []string

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
		case strings.Contains(line, "social:"):
			aboutData.SocialMedia = parseSocialMedia(strings.TrimSpace(strings.TrimPrefix(line, "social:"))) 		
		}
	}
	content := strings.TrimSpace(strings.Join(body, "\n"))
	aboutData.Content = template.HTML(ParseMarkdown(content))

	return aboutData
}

func parseSocialMedia(social string) map[string]string {
	socialMedia := make(map[string]string)
	linkKey := strings.Split(social, ":")
	socialMedia[linkKey[0]] = linkKey[1]
	return socialMedia
}