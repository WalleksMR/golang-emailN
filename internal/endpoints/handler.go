package endpoints

import "github.com/walleksmr/golang-emailn/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.IService
}
