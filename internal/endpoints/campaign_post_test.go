package endpoints

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/walleksmr/golang-emailn/internal/contract"
	"github.com/walleksmr/golang-emailn/internal/domain/campaign"
	dto "github.com/walleksmr/golang-emailn/internal/domain/campaign/dto"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) Create(input contract.NewCampaign) (string, error) {
	args := s.Called(input)
	return args.String(0), args.Error(1)
}

func (s *serviceMock) ListAll() ([]campaign.Campaign, error) {
	args := s.Called()
	return nil, args.Error(1)
}

func (s *serviceMock) GetById(id string) (*dto.GetOneOutput, error) {
	args := s.Called(id)
	return args.Get(0).(*dto.GetOneOutput), args.Error(1)
}
func (s *serviceMock) Update(input contract.CampaingUpateInput) error {
	args := s.Called(input)
	return args.Error(0)
}

func Test_CampaignPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)

	body := contract.NewCampaign{
		Name:    "any name",
		Content: "any content",
		Emails:  []string{"any@email.com"},
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	service := new(serviceMock)
	service.On("Create", mock.Anything).Return("any id", nil)

	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()
	handler := Handler{CampaignService: service}

	_, status, err := handler.CampaignPost(res, req)

	assert.Equal(201, status)
	assert.Nil(err)
}

func Test_CampaignPost_should_retorn_erro_if_exists(t *testing.T) {
	assert := assert.New(t)

	body := contract.NewCampaign{
		Name:    "any name",
		Content: "any content",
		Emails:  []string{"any@email.com"},
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	service := new(serviceMock)
	service.On("Create", mock.Anything).Return("", errors.New("any error"))

	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()
	handler := Handler{CampaignService: service}

	id, _, err := handler.CampaignPost(res, req)

	assert.NotNil(err)
	assert.Nil(id)
}
