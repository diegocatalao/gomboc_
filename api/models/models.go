package models

import (
	"time"

	gorm "gorm.io/gorm"
)

type (
	NodeModel struct {
		ID         int            `gorm:"primarykey"`
		PublicId   string         `json:"public_id" gorm:"unique;not null;type:varchar(100);default:null"`
		IsActive   bool           `json:"is_active" gorm:"default false"`
		IsPending  bool           `json:"is_pending" gorm:"default true"`
		LastIOTime time.Time      `json:"last_io_time"`
		CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
		UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
		DeletedAt  gorm.DeletedAt `gorm:"index"`
	}

	NodeRequestRegistrationModel struct {
		PublicId string `json:"public_id" gorm:"unique;not null;type:varchar(100);default:null"`
	}

	UserModel struct {
		ID        int
		Username  string         `json:"username" gorm:"not null"`
		Password  string         `json:"password" gorm:"not null"`
		CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
		UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
)
