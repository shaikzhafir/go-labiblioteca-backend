generate:
	@echo "Generating..."
	cd sqlc; sqlc generate
run:
	@echo "Running..."
	source .env && ./run_dev.sh

seed:
	@echo "Seeding..."
	go run cmd/sample-data/main.go

tw:
	@echo "Running Tailwind..."
	./tailwindcss -i ./static/input.css -o ./static/output.css --watch

build:
	@echo "Building..."
	go build -o ./bin/server cmd/server/main.go
	./tailwindcss -i ./static/input.css -o ./static/output.css --minify