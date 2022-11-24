package postgres

import (
	"vinbigdata/internal/delivery/http/model"
	"vinbigdata/internal/repository/entity"

	"gorm.io/gorm"
)

type MobileRepository struct {
	db *gorm.DB
}

func NewMobileRepository(db *gorm.DB) *MobileRepository {
	return &MobileRepository{
		db: db,
	}
}

// SaveUserCall save user call
func (r *MobileRepository) SaveUserCall(userCall *entity.UserCalls) error {
	return r.db.Create(&userCall).Error
}

// GetUserCallSummary get call summary by username
func (r *MobileRepository) GetUserCallSummary(userName string) (*model.BillingData, error) {
	var result model.BillingData
	err := r.db.Model(&entity.UserCalls{}).Where("user_name", userName).
		Select("count(*) as call_count, sum(block) as block_count").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
