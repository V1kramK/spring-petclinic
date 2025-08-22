package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_owners "github.com/input-api/mcp-server/tools/owners"
	tools_pets "github.com/input-api/mcp-server/tools/pets"
	tools_vets_html "github.com/input-api/mcp-server/tools/vets_html"
	tools_api "github.com/input-api/mcp-server/tools/api"
	tools_oups "github.com/input-api/mcp-server/tools/oups"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_owners.CreateGet_owners_owneridTool(cfg),
		tools_pets.CreateGet_pets_petid_editTool(cfg),
		tools_pets.CreatePost_pets_petid_editTool(cfg),
		tools_owners.CreateGet_ownersTool(cfg),
		tools_owners.CreatePost_owners_ownerid_pets_petid_visits_newTool(cfg),
		tools_owners.CreateGet_owners_ownerid_pets_petid_visits_newTool(cfg),
		tools_pets.CreateGet_pets_newTool(cfg),
		tools_pets.CreatePost_pets_newTool(cfg),
		tools_vets_html.CreateGet_vets_htmlTool(cfg),
		tools_api.CreateGetTool(cfg),
		tools_owners.CreateGet_owners_newTool(cfg),
		tools_owners.CreatePost_owners_newTool(cfg),
		tools_owners.CreatePost_owners_ownerid_editTool(cfg),
		tools_owners.CreateGet_owners_ownerid_editTool(cfg),
		tools_oups.CreateGet_oupsTool(cfg),
		tools_owners.CreateGet_owners_findTool(cfg),
	}
}
