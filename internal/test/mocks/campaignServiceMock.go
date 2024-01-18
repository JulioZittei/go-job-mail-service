package mockTests

import (
	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (s *CampaignServiceMock) Create(campaignInput *contract.NewCampaignInput) (id string, err error) {
	args := s.Called(campaignInput)
	return args.String(0), args.Error(1)
}

func (s *CampaignServiceMock) GetById(id string) (*contract.CampaignOutput, error) {
	args := s.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract.CampaignOutput), nil
}
