package database

import (
	"errors"

	"github.com/walleksmr/golang-emailn/internal/contract"
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

	query := c.Db.Model(campaign.Campaign{}).First(&result, "id = ?", id)

	if result.ID == "" {
		return nil, errors.New("campaign not found")
	}

	return result, query.Error
}

func (c *CampaignRepository) Update(input contract.CampaingUpateInput) error {

	var campaign *campaign.Campaign
	c.Db.First(&campaign, "id = ?", input.ID)

	if campaign.ID == "" {
		return errors.New("campaign not found")
	}

	if input.Name != nil {
		campaign.Name = *input.Name
	}

	if input.Content != nil {
		campaign.Content = *input.Content
	}

	if input.Status != nil {
		campaign.SetStatus(*input.Status)
	}

	result := c.Db.Save(&campaign)

	return result.Error
}
