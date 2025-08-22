package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws-outposts/mcp-server/config"
	"github.com/aws-outposts/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetsiteaddressHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		SiteIdVal, ok := args["SiteId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: SiteId"), nil
		}
		SiteId, ok := SiteIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: SiteId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["AddressType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("AddressType=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/sites/%s/address#AddressType%s", cfg.BaseURL, SiteId, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.GetSiteAddressOutput
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

func CreateGetsiteaddressTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_sites_SiteId_address#AddressType",
		mcp.WithDescription(" Gets the site address of the specified site. "),
		mcp.WithString("SiteId", mcp.Required(), mcp.Description(" The ID or the Amazon Resource Name (ARN) of the site. ")),
		mcp.WithString("AddressType", mcp.Required(), mcp.Description("The type of the address you request. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsiteaddressHandler(cfg),
	}
}
