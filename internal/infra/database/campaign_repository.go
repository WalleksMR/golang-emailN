package database

import (
	"github.com/walleksmr/golang-emailn/internal/domain/campaign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
	Db        *gorm.DB
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	result := c.Db.Create(&campaign)
	return result.Error
}

func (c *CampaignRepository) ListAll() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	result := c.Db.Find(&campaigns)

	return c.campaigns, result.Error
}

func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	var campaign *campaign.Campaign

	result := c.Db.First(&campaign, id)

	return campaign, result.Error
}
