package model

import (
	"time"
)

type User struct {
	Base              Base                `gorm:"embedded"`
	Username          string              `gorm:"size:20;not null;unique" json:"username"`
	Password          string              `gorm:"size:100;not null" json:"password"`
	Name              string              `gorm:"size:40;not null" json:"name"`
	Email             string              `gorm:"size:40;not null;unique" json:"email"`
	Enabled           int                 `gorm:"default:1" json:"enabled"`
	RoleName          string              `gorm:"size:40;not null" json:"role_name"`
	LoginFailCount    int                 `gorm:"default:0;size:1" json:"login_fail_count"`
	LastLoginAt       time.Time           `json:"last_login_at"`
	PasswordChangedAt time.Time           `json:"password_changed_at"`
}