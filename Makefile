.PHONY: all dev

run:
	@go run .
dev:
	@npx tailwindcss -i ./static/!tw.css -o ./static/glass.css --minify
	@go run . -dev
