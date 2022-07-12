package model

import (
	"time"
)

// Report defines the db table structure for report
type Report struct {
	//ObjectMeta `json:"metadata:omitempty"`
	//ObjectMeta `json:"metadata"`
	ObjectMeta

	CreateUser    string         `gorm:"column:create_user" json:"create_user"`
	NetworkID     int64          `gorm:"column:network_id" json:"-"`
	Status        string         `gorm:"column:status" json:"-"`
	LastRunAt     *time.Time     `gorm:"column:last_run_at" json:"last_run_at"`
	DateRange     string         `gorm:"column:date_range" json:"date_range"`
}

// TableName maps to mysql table name.
func (r *Report) TableName() string {
	return "report"
}
