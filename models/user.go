package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email         string `gorm:"email;unique;not null"`
	Name          string
	ServicesOwned pq.StringArray `gorm:"type:varchar(64)[]"`
}

func (user *UserModel) TableName() string {
	return "users"
}

func (user *UserModel) Transform() map[string]interface{} {
	return map[string]interface{}{
		"email":          user.Email,
		"name":           user.Name,
		"services_owned": user.ServicesOwned,
		"created_at":     user.CreatedAt,
		"updated_at":     user.UpdatedAt,
	}
}
