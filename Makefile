generate:
	@echo "Generating..."
	cd sqlc; sqlc generate
run:
	@echo "Running..."
	go run 	cmd/main-server/main.go
seed:
	@echo "Seeding..."
	go run cmd/sample-data/main.go