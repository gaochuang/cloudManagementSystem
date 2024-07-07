package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/models"
	"gorm.io/gorm"
)

type SystemsSettingInterface interface {
	SystemSafeSettings(ctx context.Context, systemSafeRequest *models.SystemSafeSettingsRequest) error
	GetSystemSafeSettings(ctx context.Context) (*models.SystemSettings, error)
}

type systemSetting struct {
	db *gorm.DB
}

func NewSystem(db *gorm.DB) SystemsSettingInterface {
	return &systemSetting{
		db: db,
	}
}
func (s *systemSetting) GetSystemSafeSettings(ctx context.Context) (*models.SystemSettings, error) {
	systemSetting := &models.SystemSettings{}
	err := s.db.Model(systemSetting).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get system safe setting")
	}
	return systemSetting, nil
}

func (s *systemSetting) SystemSafeSettings(ctx context.Context, systemSafeRequest *models.SystemSafeSettingsRequest) error {
	var system *models.SystemSettings

	if err := s.db.Model(&models.SystemSettings{}).First(&system).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return s.db.Model(&models.SystemSettings{}).Create(&models.SystemSettings{
				LoginFail:      systemSafeRequest.LoginFail,
				LockTime:       systemSafeRequest.LockTime,
				PasswordExpire: systemSafeRequest.PasswordExpire,
			}).Error
		}
		return fmt.Errorf("fialed to query system safe: %v", err)
	}

	return s.db.Model(&models.SystemSettings{Mode: models.Mode{ID: system.ID}}).Updates(map[string]interface{}{
		"login_fail":      systemSafeRequest.LoginFail,
		"lock_time":       systemSafeRequest.LockTime,
		"password_expire": systemSafeRequest.PasswordExpire,
	}).Error
}
