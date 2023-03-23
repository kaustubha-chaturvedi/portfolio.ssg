package main

import (
	"bufio"
	"html"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

type AboutData struct {
	Name        string
	Tagline     string
	Description string
	Image       string
	SocialMedia map[string]string
	Content     template.HTML
}

func getAboutContent(filePath string) (AboutData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return AboutData{}, err
	}
	defer file.Close()

	var aboutData AboutData
	aboutData.SocialMedia = make(map[string]string)

	inMarkdown := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "---") {
			inMarkdown = !inMarkdown
			continue
		}
		if inMarkdown {
			// Parse properties
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				switch key {
				case "name":
					aboutData.Name = value
				case "tagline":
					aboutData.Tagline = value
				case "description":
					aboutData.Description = value
				case "image":
					aboutData.Image = value
				case "github":
					aboutData.SocialMedia["GitHub"] = value
				case "linkedin":
					aboutData.SocialMedia["LinkedIn"] = value
				case "twitter":
				    aboutData.SocialMedia["Twitter"] = value
				case "instagram":
				    aboutData.SocialMedia["Instagram"] = value
				case "facebook":
				    aboutData.SocialMedia["Facebook"] = value
				case "medium":
					aboutData.SocialMedia["Medium"] = value
				case "dev":
					aboutData.SocialMedia["Dev"] = value
				case "email":
					aboutData.SocialMedia["Email"] = value
				default:
					// Ignore unknown properties
				}
			}
		} else {
			aboutData.Content += template.HTML(blackfriday.Run([]byte(html.EscapeString(line))))
		}
	}

	if err := scanner.Err(); err != nil {
		return AboutData{}, err
	}

	renderPage(filepath.Join("templates", "index.html"), aboutTemplate, struct{ AboutData }{aboutData})

	return aboutData, nil
}
