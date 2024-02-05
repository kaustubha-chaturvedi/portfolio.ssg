.PHONY: all dev

run:
	@go run main.go
dev:
	@npx tailwindcss -i ./static/!tw.css -o ./static/style.css --minify
	@go run . -dev