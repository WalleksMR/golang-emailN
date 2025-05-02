package database

import (
	"errors"

	"github.com/walleksmr/golang-emailn/internal/domain/campaign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	result := c.Db.Create(&campaign)
	return result.Error
}

func (c *CampaignRepository) ListAll() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	result := c.Db.Find(&campaigns)

	return campaigns, result.Error
}

func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	var result *campaign.Campaign

	query := c.Db.Model(campaign.Campaign{}).Preload("Contacts").First(&result, "id = ?", id)

	if result.ID == "" {
		return nil, errors.New("campaign not found")
	}

	return result, query.Error
}
