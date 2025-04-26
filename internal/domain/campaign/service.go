package campaign

import (
	"github.com/walleksmr/golang-emailn/internal/contract"
	"github.com/walleksmr/golang-emailn/internal/excptions"
)

type IService interface {
	Create(input contract.NewCampaign) (string, error)
	ListAll() ([]Campaign, error)
}

type Service struct {
	Repository Repository
}

func (s *Service) Create(input contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(input.Name, input.Content, input.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)

	if err != nil {
		return "", excptions.ErrInternal
	}

	return campaign.ID, nil
}

func (s *Service) ListAll() ([]Campaign, error) {
	campaigns, err := s.Repository.ListAll()

	if err != nil {
		return nil, err
	}

	return campaigns, nil
}
