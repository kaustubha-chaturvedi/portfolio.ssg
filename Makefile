dev:
	@npx tailwindcss -i ./static/!tw.css -o ./static/style.css --minify
	@go run .
