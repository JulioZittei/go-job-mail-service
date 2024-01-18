package mockTests

import (
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func(r *CampaignRepositoryMock) Save(campaign *model.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func(r *CampaignRepositoryMock) Get() ([]model.Campaign, error) {
	args := r.Called()
	return args.Get(0).([]model.Campaign), args.Error(1)
}

func(r *CampaignRepositoryMock) GetById(id string) (*model.Campaign, error) {
	args := r.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(*model.Campaign), nil
}