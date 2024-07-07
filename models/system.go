package models

type SystemSettings struct {
	Mode
	LoginFail      int `gorm:"column:login_fail;default:3;comment:Limit the number of login failures" json:"loginFail"`
	LockTime       int `gorm:"column:lock_time;default:30;comment:Login lock duration (zw)" json:"lockTime"`
	PasswordExpire int `gorm:"column:password_expire;default:-1; comment:Password expiration time (days). -1 The password will never expire" json:"passwordExpire"`
}

type SystemSafeSettingsRequest struct {
	LoginFail      int `json:"loginFail"`
	LockTime       int `json:"lockTime"`
	PasswordExpire int `jsopasswordExpiren:""`
}

func (system SystemSettings) TableName() string {
	return system.Mode.TableName("system")
}
