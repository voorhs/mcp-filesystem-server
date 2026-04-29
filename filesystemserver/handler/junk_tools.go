package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

func (fs *FilesystemHandler) HandleJunkTool(
	ctx context.Context,
	request mcp.CallToolRequest,
	toolName string,
	toolKind string,
	argTemplate string,
) (*mcp.CallToolResult, error) {
	_ = ctx

	params, err := fs.extractJunkParams(request, argTemplate)
	if err != nil {
		return nil, err
	}

	report := map[string]any{
		"tool":   toolName,
		"status": "ok",
		"params": params,
	}

	if toolKind == "report_write" {
		artifactPath, writeErr := fs.writeJunkWorkspaceArtifact(params, toolName, report)
		if writeErr != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error writing report artifact: %v", writeErr),
					},
				},
				IsError: true,
			}, nil
		}
		report["artifact_path"] = artifactPath
		report["artifact_written"] = true
	}

	payload, marshalErr := json.MarshalIndent(report, "", "  ")
	if marshalErr != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error generating report payload: %v", marshalErr),
				},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: string(payload),
			},
		},
	}, nil
}

func (fs *FilesystemHandler) extractJunkParams(request mcp.CallToolRequest, argTemplate string) (map[string]any, error) {
	params := map[string]any{}

	switch argTemplate {
	case "A1":
		path, err := request.RequireString("path")
		if err != nil {
			return nil, err
		}
		validPath, err := fs.validatePath(path)
		if err != nil {
			return nil, err
		}
		params["path"] = validPath
	case "A2":
		path, err := request.RequireString("path")
		if err != nil {
			return nil, err
		}
		validPath, err := fs.validatePath(path)
		if err != nil {
			return nil, err
		}
		depth := 3
		if depthVal, depthErr := request.RequireFloat("depth"); depthErr == nil {
			depth = int(depthVal)
		}
		params["path"] = validPath
		params["depth"] = depth
	case "A3":
		path, err := request.RequireString("path")
		if err != nil {
			return nil, err
		}
		validPath, err := fs.validatePath(path)
		if err != nil {
			return nil, err
		}
		pattern, err := request.RequireString("pattern")
		if err != nil {
			return nil, err
		}
		params["path"] = validPath
		params["pattern"] = pattern
	case "A4":
		path, err := request.RequireString("path")
		if err != nil {
			return nil, err
		}
		validPath, err := fs.validatePath(path)
		if err != nil {
			return nil, err
		}
		substring, err := request.RequireString("substring")
		if err != nil {
			return nil, err
		}
		maxResults := 100
		if maxVal, maxErr := request.RequireFloat("max_results"); maxErr == nil {
			maxResults = int(maxVal)
		}
		params["path"] = validPath
		params["substring"] = substring
		params["max_results"] = maxResults
	case "A5":
		path, err := request.RequireString("path")
		if err != nil {
			return nil, err
		}
		validPath, err := fs.validatePath(path)
		if err != nil {
			return nil, err
		}
		limit := 20
		if limitVal, limitErr := request.RequireFloat("limit"); limitErr == nil {
			limit = int(limitVal)
		}
		params["path"] = validPath
		params["limit"] = limit
	case "A6":
		source, err := request.RequireString("source")
		if err != nil {
			return nil, err
		}
		destination, err := request.RequireString("destination")
		if err != nil {
			return nil, err
		}
		validSource, err := fs.validatePath(source)
		if err != nil {
			return nil, err
		}
		validDestination, err := fs.validatePath(destination)
		if err != nil {
			return nil, err
		}
		params["source"] = validSource
		params["destination"] = validDestination
	case "A7":
		path, err := request.RequireString("path")
		if err != nil {
			return nil, err
		}
		validPath, err := fs.validatePath(path)
		if err != nil {
			return nil, err
		}
		mode := "default"
		if modeVal, modeErr := request.RequireString("mode"); modeErr == nil {
			mode = modeVal
		}
		params["path"] = validPath
		params["mode"] = mode
	case "A8":
		path, err := request.RequireString("path")
		if err != nil {
			return nil, err
		}
		validPath, err := fs.validatePath(path)
		if err != nil {
			return nil, err
		}
		outputName := ""
		if outputNameVal, outputNameErr := request.RequireString("output_name"); outputNameErr == nil {
			outputName = outputNameVal
		}
		params["path"] = validPath
		params["output_name"] = outputName
	default:
		path, err := request.RequireString("path")
		if err != nil {
			return nil, err
		}
		validPath, err := fs.validatePath(path)
		if err != nil {
			return nil, err
		}
		params["path"] = validPath
	}

	return params, nil
}

func (fs *FilesystemHandler) writeJunkWorkspaceArtifact(
	params map[string]any,
	toolName string,
	report map[string]any,
) (string, error) {
	pathVal, ok := params["path"].(string)
	if !ok || pathVal == "" {
		return "", fmt.Errorf("missing path for report write")
	}

	baseDir := filepath.Dir(pathVal)
	if info, statErr := os.Stat(pathVal); statErr == nil && info.IsDir() {
		baseDir = pathVal
	}

	junkDir := filepath.Join(baseDir, ".junk_workspace")
	validJunkDir, err := fs.validatePath(junkDir)
	if err != nil {
		return "", err
	}
	if mkdirErr := os.MkdirAll(validJunkDir, 0o755); mkdirErr != nil {
		return "", mkdirErr
	}

	outputName := toolName + ".json"
	if outputNameRaw, ok := params["output_name"].(string); ok && strings.TrimSpace(outputNameRaw) != "" {
		outputName = outputNameRaw
	}
	if filepath.Ext(outputName) == "" {
		outputName += ".json"
	}
	outputName = filepath.Base(outputName)

	artifactPath := filepath.Join(validJunkDir, outputName)
	payload, marshalErr := json.MarshalIndent(report, "", "  ")
	if marshalErr != nil {
		return "", marshalErr
	}
	if writeErr := os.WriteFile(artifactPath, payload, 0o644); writeErr != nil {
		return "", writeErr
	}

	return artifactPath, nil
}
