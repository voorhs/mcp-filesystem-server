package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mark3labs/mcp-filesystem-server/filesystemserver"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Parse command line arguments
	if len(os.Args) < 2 {
		fmt.Fprintf(
			os.Stderr,
			"Usage: %s <allowed-directory> [additional-directories...]\n",
			os.Args[0],
		)
		os.Exit(1)
	}

	// Create and start the server
	fss, err := filesystemserver.NewFilesystemServerWithOptions(
		os.Args[1:],
		filesystemserver.ServerOptions{
			JunkToolsEnabled: parseBoolEnv("MCP_FS_JUNK_TOOLS_ENABLED"),
			JunkToolsCount:   parseIntEnv("MCP_FS_JUNK_TOOLS_COUNT"),
		},
	)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Serve requests
	if err := server.ServeStdio(fss); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func parseBoolEnv(key string) bool {
	value := strings.TrimSpace(strings.ToLower(os.Getenv(key)))
	switch value {
	case "1", "true", "yes", "y", "on":
		return true
	default:
		return false
	}
}

func parseIntEnv(key string) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return 0
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}
