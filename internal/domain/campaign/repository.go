package campaign

import "github.com/walleksmr/golang-emailn/internal/contract"

type Repository interface {
	Save(campaign *Campaign) error
	ListAll() ([]Campaign, error)
	GetById(id string) (*Campaign, error)
	Update(input contract.CampaingUpateInput) error
}
