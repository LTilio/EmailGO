package endpoints

import "EmailGO/internal/campaign"

type Handler struct {
	CampaignService campaign.Service
}
