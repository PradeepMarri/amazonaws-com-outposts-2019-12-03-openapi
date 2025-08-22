package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/aws-outposts/mcp-server/config"
	"github.com/aws-outposts/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateoutpostHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/outposts", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CreateOutpostOutput
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreateoutpostTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_outposts",
		mcp.WithDescription("<p>Creates an Outpost.</p> <p>You can specify either an Availability one or an AZ ID.</p>"),
		mcp.WithString("AvailabilityZoneId", mcp.Description("Input parameter: The ID of the Availability Zone.")),
		mcp.WithString("Description", mcp.Description("Input parameter: The description of the Outpost.")),
		mcp.WithString("Name", mcp.Required(), mcp.Description("Input parameter: The name of the Outpost.")),
		mcp.WithString("SiteId", mcp.Required(), mcp.Description("Input parameter: The ID of the site.")),
		mcp.WithString("SupportedHardwareType", mcp.Description("Input parameter:  The type of hardware for this Outpost. ")),
		mcp.WithObject("Tags", mcp.Description("Input parameter: The tags to apply to the Outpost.")),
		mcp.WithString("AvailabilityZone", mcp.Description("Input parameter: The Availability Zone.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateoutpostHandler(cfg),
	}
}
