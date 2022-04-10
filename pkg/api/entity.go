package api

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID                uuid.UUID           `gorm:"size:36;primaryKey" json:"id`
	Username          string              `gorm:"size:20;not null;unique" json:"username"`
	Password          string              `gorm:"size:100;not null" json:"password"`
	Name              string              `gorm:"size:40;not null" json:"name"`
	Email             string              `gorm:"size:40;not null;unique" json:"email"`
	Enabled           bool                `json:"enabled"`
	RoleName          string              `gorm:"size:40;not null" json:"role_name"`
	LoginFailCount    int                 `gorm:"size:1" json:"login_fail_count"`
	LastLoginAt       time.Time           `json:"last_login_at"`
	PasswordChangedAt time.Time           `json:"password_changed_at"`
	CreatedAt         time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("ID", uuid)
   }