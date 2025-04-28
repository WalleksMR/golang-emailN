package campaign

import (
	"github.com/walleksmr/golang-emailn/internal/contract"
	"github.com/walleksmr/golang-emailn/internal/domain/campaign/dto"
	"github.com/walleksmr/golang-emailn/internal/excptions"
)

type IService interface {
	Create(input contract.NewCampaign) (string, error)
	ListAll() ([]Campaign, error)
	GetById(id string) (*dto.GetOneOutput, error)
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

func (s *Service) GetById(id string) (*dto.GetOneOutput, error) {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return nil, err
	}

	if campaign == nil {
		return nil, nil
	}

	return &dto.GetOneOutput{
		Id:      campaign.ID,
		Name:    campaign.Name,
		Status:  string(campaign.Status),
		Content: campaign.Content,
	}, nil
}
