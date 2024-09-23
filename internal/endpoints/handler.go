package endpoints

import "EmailGO/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
