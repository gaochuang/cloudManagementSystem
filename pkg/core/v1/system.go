package v1

import (
	"context"
	"github.com/gaochuang/cloudManagementSystem/cmd/app/config"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/database"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"go.uber.org/zap"
)

type SystemSettingGetter interface {
	SystemSetting() SystemsInterface
}

type SystemsInterface interface {
	SystemSafeSettings(ctx context.Context, systemSafeRequest *models.SystemSafeSettingsRequest) error
	GetSystemSafeSettings(ctx context.Context) (*models.SystemSettings, error)
}

type systemSetting struct {
	config  config.Config
	factory database.ShareFactory
}

func newSystemSetting(p *platform) SystemsInterface {
	return &systemSetting{
		config:  p.config,
		factory: p.factory,
	}
}

func (s *systemSetting) GetSystemSafeSettings(ctx context.Context) (*models.SystemSettings, error) {
	return s.factory.System().GetSystemSafeSettings(ctx)
}

func (s *systemSetting) SystemSafeSettings(ctx context.Context, systemSafeRequest *models.SystemSafeSettingsRequest) error {
	if err := s.factory.System().SystemSafeSettings(ctx, systemSafeRequest); err != nil {
		log.Logger.LogErrorWithCtx(ctx, "set system safe failed", zap.Error(err))
	}
	return nil
}
