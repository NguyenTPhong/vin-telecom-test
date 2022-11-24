package service

import (
	"net/http"
	"vinbigdata/internal/delivery/http/model"
	"vinbigdata/internal/repository/entity"
	error2 "vinbigdata/package/error"
)

const BlockDuration = 30 // 30s
type MobileService struct {
	mobileRepo MobileRepository
}

type MobileRepository interface {
	GetUserCallSummary(userName string) (*model.BillingData, error)
	SaveUserCall(userCall *entity.UserCalls) error
}

func NewMobileService(mobileRepo MobileRepository) *MobileService {
	return &MobileService{
		mobileRepo: mobileRepo,
	}
}

func (s *MobileService) SaveUserCall(userName string, duration int64) error {
	roundDuration := duration / 1000
	block := roundDuration / BlockDuration
	if roundDuration%BlockDuration > 0 {
		block++
	}

	userCall := &entity.UserCalls{
		UserName: userName,
		Duration: duration,
		Block:    block,
	}

	if err := s.mobileRepo.SaveUserCall(userCall); err != nil {
		return error2.NewCError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (s *MobileService) GetUserBilling(userName string) (*model.BillingData, error) {
	data, err := s.mobileRepo.GetUserCallSummary(userName)
	if err != nil {
		return nil, error2.NewCError(http.StatusInternalServerError, err.Error())
	}
	return data, nil
}
