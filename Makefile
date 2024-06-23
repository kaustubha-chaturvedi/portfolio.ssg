.PHONY: all dev

run:
	@go run .
dev:
	@npx tailwindcss -i ./static/!tw.css -o ./static/level5gyat.css --minify
	@go run .

air:
	@npx tailwindcss -i ./static/!tw.css -o ./static/level5gyat.css --minify
	@air