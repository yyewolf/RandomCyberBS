dev:
		air

prettier:
	npx prettier --write .

css:
	npx tailwindcss -i ./tailwind.css -o ./static/tailwind.min.css --minify

css-watch:
	npx tailwindcss -i ./tailwind.css -o ./static/tailwind.min.css --minify --watch