# Portfolio Static Site Generator Using Go

## Description

A simple static site generator written in Go. It uses Go's pongo'2 package to render the templates and markdown to write the content.
It is designed to be easily customizable. 
Motivation for this project was to create a simple static site generator that can be used to create a portfolio for me and my friends.
Also, I wanted to learn Go and this project was a good way to do that.

## Clarifications

Yup this repo is based on my [https://github.com/kaustubha-chaturvedi/portfolio-echo](this repo) from previous year that time I was just learning go syntax and it's framework and library ecosystem,
I halted that due to my eams and job search (ofcourse not because of laziness), I am back at learning Golang and this is my final version of static site portfolio,
I hope it to be useful (atleast to me), and to allow anyone to be able to edit this I have used pongo2 as more people are comfortable with jinja2 style templating.
If you are here reaing this please make use of your time at some better place.

## Installation

To install the project, you need to have Go and npm installed on your machine. Once you have these prerequisites, you can clone the repository and install the dependencies:

```sh
go mod download
```
## How to use

To use the project, you need to have Go and npm installed on your machine. Once you have these prerequisites, you can clone the repository and install the dependencies:

```sh
go run .
```

## Deployment

To deploy this project on Render,just import the repository and Render will automatically detect the project type and build it for you.

## How to edit Templates

To edit the templates, edit the files in the templates folder. The templates are written in Go's pongo2 template language. Which is similar to Jinja2.
