package v1

import (
	"context"
	"github.com/gaochuang/cloudManagementSystem/cmd/app/config"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/database"
)

type RoleGetter interface {
	Role() RoleInterface
}

type role struct {
	config  config.Config
	factory database.ShareFactory
}

func newRole(p *platform) RoleInterface {
	return &role{
		config:  p.config,
		factory: p.factory,
	}
}

type RoleInterface interface {
	List(ctx context.Context, page *models.PageResult, keyword string) (result models.PageResult, err error)
}

func (r *role) List(ctx context.Context, page *models.PageResult, keyword string) (result models.PageResult, err error) {
	roles, err := r.factory.Role().List(ctx, page, keyword)
	result = models.PageResult{
		Page:     page.Page,
		PageSize: page.PageSize,
		Total:    page.Total,
		Items:    roles,
	}
	return result, err
}
