package campaign

import (
	"errors"

	"github.com/walleksmr/golang-emailn/internal/contract"
	"github.com/walleksmr/golang-emailn/internal/domain/campaign/dto"
	"github.com/walleksmr/golang-emailn/internal/excptions"
	"gorm.io/gorm"
)

type IService interface {
	Create(input contract.NewCampaign) (string, error)
	ListAll() ([]Campaign, error)
	GetById(id string) (*dto.GetOneOutput, error)
	Cancel(input contract.CampaingUpateInput) error
}

type Service struct {
	Repository Repository
	Db         *gorm.DB
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
		Id:                   campaign.ID,
		Name:                 campaign.Name,
		Status:               string(campaign.Status),
		Content:              campaign.Content,
		AmountOfEmailsToSend: len(campaign.Contacts),
	}, nil
}

func (s *Service) Cancel(input contract.CampaingUpateInput) error {
	var campaing *Campaign
	s.Db.First(&campaing, "id = ?", input.ID)
	if campaing.ID == "" {
		return errors.New("campaign not fount")
	}

	if campaing.Status == StatusCanceled {
		return errors.New("campaign already is canceled")
	}

	err := campaing.Cancel()
	if err != nil {
		return err
	}
	s.Db.Save(&campaing)
	return nil
}
