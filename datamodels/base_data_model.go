package datamodels

import (
	"time"
)

type BaseDataModel struct {
	Id           string    `gorm:"primary_key;default:uuid_generate_v4()"`
	DateCreated  time.Time `gorm:"autoCreateTime"`
	DateModified time.Time `gorm:"autoUpdateTime"`
}
