package models

type RolePermission struct {
	Mode
	RoleId       int `gorm:"column:role_id;comment:'role id'" json:"role_id"`
	PermissionId int `gorm:"column:permission_id;comment:'permission id'" json:"permission_id"`
}

func (m RolePermission) TableName() string {
	return m.Mode.TableName("role_permission")
}
