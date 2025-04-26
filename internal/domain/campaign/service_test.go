package campaign

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/walleksmr/golang-emailn/internal/contract"
	"github.com/walleksmr/golang-emailn/internal/excptions"
)

type campaignRepositoryMock struct {
	mock.Mock
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "any name",
		Content: "any content",
		Emails:  []string{"any@email.com"},
	}
	service = Service{}
)

func (r *campaignRepositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *campaignRepositoryMock) ListAll() ([]Campaign, error) {
	return nil, nil
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(campaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotEmpty(id)
	assert.Nil(err)
	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(campaignRepositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("Save", mock.Anything).Return(errors.New(""))

	id, err := service.Create(newCampaign)

	assert.Empty(id)
	assert.Equal(excptions.ErrInternal.Error(), err.Error())
	assert.True(errors.Is(excptions.ErrInternal, err))
	repositoryMock.AssertExpectations(t)
}
