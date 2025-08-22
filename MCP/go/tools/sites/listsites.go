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

func ListsitesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		if val, ok := args["OperatingAddressCountryCodeFilter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("OperatingAddressCountryCodeFilter=%v", val))
		}
		if val, ok := args["OperatingAddressStateOrRegionFilter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("OperatingAddressStateOrRegionFilter=%v", val))
		}
		if val, ok := args["OperatingAddressCityFilter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("OperatingAddressCityFilter=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/sites%s", cfg.BaseURL, queryString)
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
		var result models.ListSitesOutput
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

func CreateListsitesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_sites",
		mcp.WithDescription("<p>Lists the Outpost sites for your Amazon Web Services account. Use filters to return specific results.</p> <p>Use filters to return specific results. If you specify multiple filters, the results include only the resources that match all of the specified filters. For a filter where you can specify multiple values, the results include items that match any of the values that you specify for the filter.</p>"),
		mcp.WithString("NextToken", mcp.Description("")),
		mcp.WithNumber("MaxResults", mcp.Description("")),
		mcp.WithArray("OperatingAddressCountryCodeFilter", mcp.Description("Filters the results by country code.")),
		mcp.WithArray("OperatingAddressStateOrRegionFilter", mcp.Description("Filters the results by state or region.")),
		mcp.WithArray("OperatingAddressCityFilter", mcp.Description("Filters the results by city.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListsitesHandler(cfg),
	}
}
