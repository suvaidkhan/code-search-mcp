package main

import (
	"log"
	"os"
	"strings"

	_ "embed"

	"github.com/suvaidkhan/code-explore-mcp/internal/mcp"
)

var Version string

func main() {
	Version = strings.TrimSpace(Version)

	workspaceRoot := os.Getenv("CODE_SEARCH_WORKSPACE_ROOT")
	if workspaceRoot == "" {
		workspaceRoot = "."
	}

	server, err := mcp.NewServer(workspaceRoot, Version)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()
	log.Println("Starting server")
	if err := server.Serve(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
