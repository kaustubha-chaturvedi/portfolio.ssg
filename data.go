package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

// Post represents a blog post
type Post struct {
	Title       string
	Description string
	Author      string
	Date        string
	Slug        string
	Content     template.HTML
}

// Project represents a project
type Project struct {
	Title       string
	Description string
	Author      string
	Date        string
	URL         string
	Tags        []string
}

func getAllData(srcDir, destDir, templateFile, pagePrefix string, itemsPerPage int) ([]interface{}, int) {
	var allData []interface{}

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || !strings.Contains(path, ".md") {
			return err
		}

		item := getData(path)
		allData = append(allData, item)
		return nil
	})

	totalPages := (len(allData) + itemsPerPage - 1) / itemsPerPage
	if err != nil {
		log.Fatal(err)
	}

	renderPages(srcDir, destDir, templateFile, pagePrefix, allData, totalPages, itemsPerPage)

	return allData, totalPages
}

func getData(path string) interface{} {
	isProject := strings.Contains(path, "projects")
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var content string
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

	if isProject {
		project := parseProject(header, body)
		return project
	}

	post := parsePost(header, body)
	post.Slug = strings.ReplaceAll(strings.ToLower(post.Title), " ", "-")
	content = strings.TrimSpace(strings.Join(body, "\n"))
	post.Content = template.HTML(blackfriday.Run([]byte(content)))
	return post
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

func renderPages(srcDir, destDir, templateFile, pagePrefix string, data []interface{}, totalPages, itemsPerPage int) {
	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * itemsPerPage
		end := start + itemsPerPage
		if end > len(data) {
			end = len(data)
		}
		currentData := data[start:end]
		renderPage(filepath.Join("templates", templateFile), filepath.Join(destDir, fmt.Sprintf("%s%d.html", pagePrefix, page)), struct {
			Data        []interface{}
			TotalPages  int
			CurrentPage int
		}{
			Data:        currentData,
			TotalPages:  totalPages,
			CurrentPage: page,
		})
	}
}
