package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Offers struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);unique_index"`
	Desc      string
	StartDate time.Time
	EndDate   time.Time

	UserID int
	User   User
}
