#!/bin/bash

# Find and delete files with the specified pattern
echo "🧹 cleaning up ..."
find . -type f -name "*_templ.go" -exec rm -f {} \;

echo "🧪 generating ..."
templ generate

# Build the Go program
echo "🚀 deploying ..."
go run *.go

echo "✨ done"