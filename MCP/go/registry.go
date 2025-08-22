package main

import (
	"github.com/aws-outposts/mcp-server/config"
	"github.com/aws-outposts/mcp-server/models"
	tools_connections "github.com/aws-outposts/mcp-server/tools/connections"
	tools_sites "github.com/aws-outposts/mcp-server/tools/sites"
	tools_outposts "github.com/aws-outposts/mcp-server/tools/outposts"
	tools_catalog "github.com/aws-outposts/mcp-server/tools/catalog"
	tools_tags "github.com/aws-outposts/mcp-server/tools/tags"
	tools_orders "github.com/aws-outposts/mcp-server/tools/orders"
	tools_list_orders "github.com/aws-outposts/mcp-server/tools/list_orders"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_connections.CreateStartconnectionTool(cfg),
		tools_sites.CreateDeletesiteTool(cfg),
		tools_sites.CreateGetsiteTool(cfg),
		tools_sites.CreateUpdatesiteTool(cfg),
		tools_outposts.CreateCreateoutpostTool(cfg),
		tools_outposts.CreateListoutpostsTool(cfg),
		tools_sites.CreateCreatesiteTool(cfg),
		tools_sites.CreateListsitesTool(cfg),
		tools_catalog.CreateGetcatalogitemTool(cfg),
		tools_outposts.CreateDeleteoutpostTool(cfg),
		tools_outposts.CreateGetoutpostTool(cfg),
		tools_outposts.CreateUpdateoutpostTool(cfg),
		tools_sites.CreateUpdatesiterackphysicalpropertiesTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_catalog.CreateListcatalogitemsTool(cfg),
		tools_outposts.CreateListassetsTool(cfg),
		tools_connections.CreateGetconnectionTool(cfg),
		tools_orders.CreateCreateorderTool(cfg),
		tools_list_orders.CreateListordersTool(cfg),
		tools_orders.CreateGetorderTool(cfg),
		tools_orders.CreateCancelorderTool(cfg),
		tools_sites.CreateGetsiteaddressTool(cfg),
		tools_outposts.CreateGetoutpostinstancetypesTool(cfg),
		tools_sites.CreateUpdatesiteaddressTool(cfg),
	}
}
