package role

import (
	"context"
	"github.com/gaochuang/cloudManagementSystem/models"
	"gorm.io/gorm"
)

type RoleInterface interface {
	List(ctx context.Context, page *models.PageResult, keyword string) (roles *[]models.Role, err error)
}

type role struct {
	db *gorm.DB
}

func NewRole(db *gorm.DB) RoleInterface {
	return &role{
		db: db,
	}
}

func (r *role) List(ctx context.Context, page *models.PageResult, keyword string) (roles *[]models.Role, err error) {
	page, offset := models.SetPageDefaults(page)
	db := r.db.Model(&models.Role{})
	if keyword != "" {
		db = db.Where("name like ?", "%"+keyword+"%")
	}

	if err = db.Limit(page.PageSize).Offset(offset).Find(&roles).Count(&page.Total).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
