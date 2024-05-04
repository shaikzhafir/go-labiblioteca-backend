#!/bin/bash
find ./.. \( -name "*.go" -o -name "*.html" -o -name "*tmpl" \) ! -name "*_test.go" | entr -r go run ./cmd/server/main.go
